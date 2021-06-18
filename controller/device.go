package controller

import (
	"log"
	"net/http"
	"path"

	"github.com/ashmintech/azurewithgo/data"
)

func Devices(w http.ResponseWriter, r *http.Request) {

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
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	d = findDevices4Customer(c.CustID)

	type sendData struct {
		Dev  data.Devices
		Cust *data.Customer
	}

	if err := tpl.ExecuteTemplate(w, "devices.gohtml", sendData{d, c}); err != nil {
		log.Fatalln("Not able to call the template", err)
	}
}

func existDevice(d string) (*data.Device, bool) {
	return data.GetDevice(d)
}

func findCustomer4Device(d string) (*data.Customer, bool) {
	return data.GetCustomer4Device(d)
}

func DeviceToggleStatus(w http.ResponseWriter, r *http.Request) {

	devID := path.Base(r.URL.Path)

	d, ok := data.ToggleDeviceStatus(devID)

	if !ok {
		http.Redirect(w, r, "/customer/devices", http.StatusSeeOther)
		return
	} else {
		http.Redirect(w, r, "/customer/devices/"+d.DeviceID, http.StatusSeeOther)
		return

	}

}

func DeviceDetails(w http.ResponseWriter, r *http.Request) {

	devID := path.Base(r.URL.Path)
	var d *data.Device
	var c *data.Customer

	d, found := existDevice(devID)

	if !found {
		http.Redirect(w, r, "/customer", http.StatusSeeOther)
		return
	}

	c, _ = findCustomer4Device(d.DeviceID)

	type sendData struct {
		Dev  *data.Device
		Cust *data.Customer
	}

	if err := tpl.ExecuteTemplate(w, "devicedetails.gohtml", sendData{d, c}); err != nil {
		log.Fatalln("Not able to call the template", err)
	}
}
