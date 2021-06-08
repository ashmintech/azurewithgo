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

var dummyCust = Customer{
	CustID:       "DummyID",
	FName:        "Dummy",
	LName:        "Name",
	Address:      "Pluto",
	Phone:        "(123)456 7890",
	Email:        "Dummy@Email",
	SubType:      "Not sure",
	Active:       false,
	CreationDate: "January 1, 1900",
}

// Customers is a collection of customer
type Customers []*Customer

var customerList = Customers{}

func GetCustomers() []*Customer {
	return customerList
}

func GetCustomer(custID string) *Customer {
	for _, b := range customerList {
		if b.CustID == custID {
			return b
		}
	}
	return &dummyCust
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
