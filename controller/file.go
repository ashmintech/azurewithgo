package controller

import (
	"log"
	"net/http"

	"github.com/ashmintech/azurewithgo/data"
)

func Files(w http.ResponseWriter, r *http.Request) {

	custID := r.URL.Query().Get("cust")

	_, err := r.Cookie("session")

	if err == http.ErrNoCookie {
		log.Println("No cookie Found. Redirecting to home page")
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	c, found := existCustomer(custID)

	if !found {
		http.Redirect(w, r, "/customer", http.StatusSeeOther)
		return
	}
	type sendData struct {
		Cust  *data.Customer
		Files data.Files
	}

	sd := sendData{
		c,
		data.GetFiles(),
	}

	if err := tpl.ExecuteTemplate(w, "files.gohtml", sd); err != nil {
		log.Fatalln("Not able to call the template", err)
	}

}
