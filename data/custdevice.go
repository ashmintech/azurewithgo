package data

type C2D struct {
	CustomerID string `json:"customerid"`
	DeviceID   string `json:"deviceid"`
}

var C2DList = []C2D{
	{
		CustomerID: "32891c71-4b55-401f-a819-31950f331b5b",
		DeviceID:   "smartdevice1",
	},
	{
		CustomerID: "32891c71-4b55-401f-a819-31950f331b5b",
		DeviceID:   "smartdevice2",
	},
}

var cust2Device C2D
var devices2cust []C2D

func GetCustomer4Device(d string) (*Customer, bool) {
	for pos, dev := range C2DList {
		if dev.DeviceID == d {
			cust2Device = C2DList[pos]
		}
	}
	return GetCustomer(cust2Device.CustomerID)
}

func GetDevices4Customer(c string) Devices {
	var deviceNames []string

	for _, cust := range C2DList {
		if cust.CustomerID == c {
			devices2cust = append(devices2cust, cust)
			deviceNames = append(deviceNames, cust.DeviceID)
		}
	}
	return GetDevices(deviceNames)
}
