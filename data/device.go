package data

import (
	"log"

	"github.com/globalsign/mgo/bson"
)

const (
	DeviceCollName = "devices"
)

// Devices is a collection of device
type Devices []*Device

type Device struct {
	DeviceID           string `json:"deviceid"`
	DeviceName         string `json:"devicename"`
	DeviceModel        string `json:"devicemodel"`
	DeviceType         string `json:"devicetype"`
	DeviceStatus       string `json:"devicestatus"`
	DeviceCreationDate string `json:"devicecreationdate"`
}

var DeviceList = []*Device{}

func GetDevice(d string) (*Device, bool) {
	var dev Device

	mcoll := GetCollection(DeviceCollName)

	err := mcoll.Find(bson.M{"deviceid": d}).One(&dev)
	if err != nil {
		log.Println("Error while querying the Collection:\n", err)
		return nil, false
	}
	return &dev, true
}

func GetDevices(deviceList []string) Devices {

	var dList Devices
	for _, devID := range deviceList {
		if d, found := GetDevice(devID); found {
			dList = append(dList, d)
		}
	}
	return dList
}
