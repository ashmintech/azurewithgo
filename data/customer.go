package data

import (
	"errors"
)

type Customer struct {
	CustID       string `json:"customerid"`
	FName        string `json:"fname"`
	LName        string `json:"lname"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	SubType      string `json:"subtype"`
	Active       bool   `json:"active"`
	CreationDate string `json:"creation"`
}

// Customers is a collection of customer
type Customers []*Customer

var customerList = []*Customer{
	{
		CustID:       "32891c71-4b55-401f-a819-31950f331b5b",
		FName:        "Ashish",
		LName:        "Minocha",
		Address:      "Canada",
		Phone:        "(123) 456-7890",
		Email:        "minocha_ashish@hotmail.com",
		SubType:      "Premium",
		Active:       true,
		CreationDate: "Apr 10, 2021",
	},
	{
		CustID:       "custid2",
		FName:        "Ashish",
		LName:        "Minocha",
		Address:      "USA",
		Phone:        "(987) 654-3210",
		Email:        "ashmintech@outlook.com",
		SubType:      "Standard",
		Active:       false,
		CreationDate: "Mar 2 2021",
	},
}

func GetCustomer(custID string) (*Customer, bool) {

	for _, b := range customerList {
		if b.CustID == custID {
			return b, true
		}
	}
	return nil, false
}

func findCustomer(custID string) bool {
	for _, c := range customerList {
		if c.CustID == custID {
			return true
		}
	}
	return false
}

func AddCustomer(p *Customer) (bool, error) {
	if !findCustomer(p.CustID) {
		customerList = append(customerList, p)
	} else {
		return false, errors.New("Customer already exists")
	}
	return true, nil
}
