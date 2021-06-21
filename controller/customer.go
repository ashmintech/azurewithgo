package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"

	data "github.com/ashmintech/azurewithgo/data"
)

var tpl *template.Template
var c *http.Cookie
var oauthvar *oauth2.Config

var cID, cSecret, loginUrl, logoutUrl string

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	cID = os.Getenv("AZURE_CLIENT_ID")
	cSecret = os.Getenv("AZURE_CLIENT_SECRET")
	loginUrl = os.Getenv("AZUREB2C_LOGIN_REDIRECT_URL")
	logoutUrl = os.Getenv("AZUREB2C_LOGOUT_REDIRECT_URL")

	// If any of the value is empty
	if cID == "" || cSecret == "" || loginUrl == "" || logoutUrl == "" {
		log.Fatalln("Not able to set environmental variables")
	}

	oauthvar = &oauth2.Config{
		ClientID:     cID,
		RedirectURL:  loginUrl,
		ClientSecret: cSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://ashmintech.b2clogin.com/ashmintech.onmicrosoft.com/B2C_1_signinsignup/oauth2/v2.0/authorize",
			TokenURL: "https://ashmintech.b2clogin.com/ashmintech.onmicrosoft.com/B2C_1_signinsignup/oauth2/v2.0/token",
		},
		Scopes: []string{"openid", cID, "offline_access"},
	}

	// OpenID Configuration: https://ashmintech.b2clogin.com/ashmintech.onmicrosoft.com/B2C_1_signinsignup/v2.0/.well-known/openid-configuration
	// Keys: https://ashmintech.b2clogin.com/ashmintech.onmicrosoft.com/B2C_1_signinsignup/discovery/v2.0/keys

	go data.RunEventHubListener()
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cook, err := r.Cookie("session")
		if err == http.ErrNoCookie {
			urlstring := oauthvar.AuthCodeURL("thisstate")
			http.Redirect(w, r, urlstring, http.StatusSeeOther)
			return
		}

		v := r.URL.Query()
		v.Add("cust", cook.Value)
		r.URL.RawQuery = v.Encode()

		next.ServeHTTP(w, r)
	})
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("state") != "thisstate" {
		http.Redirect(w, r, "/home?status=nosuccess", http.StatusSeeOther)
		return
	}

	if r.FormValue("error") == "access_denied" {
		http.Redirect(w, r, "/home?status=nosuccess", http.StatusSeeOther)
		return
	}

	oauthtoken, err := oauthvar.Exchange(r.Context(), r.FormValue("code"))
	if err != nil {
		log.Fatalln("Token is not valid\n", err)

	}

	a, err := jwt.Parse(oauthtoken.AccessToken, nil)

	var custID string
	var newCust string

	custClaims := a.Claims.(jwt.MapClaims)

	custID = fmt.Sprintf("%v", custClaims["sub"])
	email := fmt.Sprintf("%v", custClaims["emails"])
	email = strings.TrimLeft(email, "[")
	email = strings.TrimRight(email, "]")

	newCust = fmt.Sprintf("%v", custClaims["newUser"])

	if newCust == "true" {
		log.Println("It's a new customer: ")
	}

	cust := &data.Customer{
		CustomerID:   custID,
		FName:        fmt.Sprintf("%v", custClaims["given_name"]),
		LName:        fmt.Sprintf("%v", custClaims["family_name"]),
		Address:      fmt.Sprintf("%v", custClaims["country"]),
		Phone:        "(---)--- ----",
		Email:        email,
		SubType:      fmt.Sprintf("%v", custClaims["extension_Subscription"]),
		Active:       true,
		CreationDate: time.Now().Format("January 2, 2006"),
	}

	if ok, _ := data.AddCustomer(cust); !ok {
		http.Redirect(w, r, "/home?status=nosuccess", http.StatusSeeOther)
		return
	}

	c = &http.Cookie{
		Name:  "session",
		Value: custID,
	}
	http.SetCookie(w, c)
	http.Redirect(w, r, "/customer", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {

	c = &http.Cookie{
		Name:   "session",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	u, _ := url.Parse("https://ashmintech.b2clogin.com/ashmintech.onmicrosoft.com/B2C_1_signinsignup/oauth2/v2.0/logout")
	q := u.Query()
	q.Add("post_logout_redirect_uri", logoutUrl)
	u.RawQuery = q.Encode()
	http.Redirect(w, r, u.String(), http.StatusSeeOther)
}

func findDevices4Customer(c string) data.Devices {
	return data.GetDevices4Customer(c)
}

func existCustomer(c string) (*data.Customer, bool) {
	return data.GetCustomer(c)
}

func CustomerDetails(w http.ResponseWriter, r *http.Request) {

	var d data.Devices
	var c *data.Customer

	custID := r.URL.Query().Get("cust")

	_, err := r.Cookie("session")

	if err == http.ErrNoCookie {
		log.Println("No cookie Found. Redirecting to home page")
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	c, found := existCustomer(custID)

	if !found {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	d = findDevices4Customer(c.CustomerID)

	type sendData struct {
		Dev  data.Devices
		Cust *data.Customer
	}

	if err := tpl.ExecuteTemplate(w, "customerdetails.gohtml", sendData{d, c}); err != nil {
		log.Fatalln("Not able to call the template", err)
	}
}
