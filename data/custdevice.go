package data

import (
	"log"

	"github.com/globalsign/mgo/bson"
)

type C2D struct {
	CustomerID string `json:"customerid"`
	DeviceID   string `json:"deviceid"`
}

const (
	C2DCollName = "cust2device"
)

func GetCustomer4Device(d string) (*Customer, bool) {
	var cust2Device C2D

	mcoll := GetCollection(C2DCollName)
	q := mcoll.Find(bson.M{"deviceid": d})

	err := q.One(&cust2Device)
	if err != nil {
		log.Println("Cannot find customer for this device:\n", err)
		return nil, false
	}

	return GetCustomer(cust2Device.CustomerID)
}

func GetDevices4Customer(c string) Devices {
	var deviceNames []string
	var devices2cust []C2D

	mcoll := GetCollection(C2DCollName)

	err = mcoll.Find(bson.M{"customerid": c}).Iter().All(&devices2cust)
	if err != nil {
		log.Println("Error while querying the Collection:\n", err)
		return nil
	}

	for _, d := range devices2cust {
		deviceNames = append(deviceNames, d.DeviceID)
	}

	return GetDevices(deviceNames)
}
