package controller

import (
	"log"
	"net/http"
	"path"
	"time"

	"github.com/ashmintech/azurewithgo/data"
)

func DeviceData(w http.ResponseWriter, r *http.Request) {

	devID := path.Base(r.URL.Path)

	if _, found := existDevice(devID); !found {
		http.Redirect(w, r, "/customer", http.StatusSeeOther)
		return
	}

	if err := tpl.ExecuteTemplate(w, "devicedata.gohtml", data.GetDeviceData(devID, time.Now().UTC().Day())); err != nil {
		log.Fatalln("Not able to call the template", err)
	}
}
