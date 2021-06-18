package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	cont "github.com/ashmintech/azurewithgo/controller"
)

const hostAddress = ":8080"

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/logout", cont.Logout)
	r.HandleFunc("/login", cont.Login)

	cu := r.Methods(http.MethodGet).Subrouter()
	cu.HandleFunc("/customer", cont.CustomerDetails)
	cu.HandleFunc("/customer/devices", cont.Devices)
	cu.HandleFunc("/customer/devices/{id}", cont.DeviceDetails)
	cu.HandleFunc("/customer/devicedata/{id}", cont.DeviceData)
	cu.HandleFunc("/customer/devicestatus/{id}", cont.DeviceToggleStatus)
	cu.Use(cont.Middleware)

	// Static Files
	r.PathPrefix("/home").Handler(http.FileServer(http.Dir(".")))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/home", http.StatusMovedPermanently)
	})

	srv := &http.Server{
		Handler:      r,
		Addr:         hostAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Println(("Starting Server on port 8080"))

		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("Server shutting down")
	os.Exit(0)
}
