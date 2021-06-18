package data

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

var DeviceList = []*Device{
	{
		DeviceID:           "smartdevice1",
		DeviceName:         "Device Name 1",
		DeviceModel:        "Globomantics Fridge",
		DeviceType:         "Fridge",
		DeviceStatus:       "Active",
		DeviceCreationDate: "May 4, 2021",
	},

	{
		DeviceID:           "smartdevice2",
		DeviceName:         "Device Name 2",
		DeviceModel:        "Globomantics Fridge",
		DeviceType:         "Fridge",
		DeviceStatus:       "Inactive",
		DeviceCreationDate: "May 14, 2021",
	},
}

func GetAllDevices() Devices {
	return DeviceList
}

func GetDeviceCount() int {
	return len(DeviceList)
}

func GetDevice(d string) (*Device, bool) {
	for _, b := range DeviceList {
		if b.DeviceID == d {
			return b, true
		}
	}
	return nil, false
}

func GetDevices(d []string) Devices {

	var dList Devices

	for _, devID := range d {
		for _, d := range DeviceList {
			if devID == d.DeviceID {
				dList = append(dList, d)
			}
		}
	}

	return dList

}

func ToggleDeviceStatus(d string) (*Device, bool) {

	if dev, found := GetDevice(d); found {
		if dev.DeviceStatus == "Active" {
			dev.DeviceStatus = "Inactive"
		} else {
			dev.DeviceStatus = "Active"
		}
		return dev, true

	} else {
		return nil, false
	}

}
