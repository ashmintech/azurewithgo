package data

import (
	"log"

	"github.com/globalsign/mgo/bson"
)

const (
	CustomerCollName = "customers"
)

type Customer struct {
	CustomerID   string `json:"customerid"`
	FName        string `json:"fname"`
	LName        string `json:"lname"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	SubType      string `json:"subtype"`
	Active       bool   `json:"active"`
	CreationDate string `json:"creation"`
}

func GetCustomer(custID string) (*Customer, bool) {
	var cust Customer

	mcoll := GetCollection(CustomerCollName)
	err := mcoll.Find(bson.M{"customerid": custID}).One(&cust)
	if err != nil {
		log.Println("Error while querying the Collection: ", err)
		return nil, false
	}

	return &cust, true
}

func findCustomer(custID string) bool {

	mcoll := GetCollection(CustomerCollName)

	n, err := mcoll.Find(bson.M{"customerid": custID}).Count()
	if err != nil {
		log.Println("Error while querying the Collection: ", err)
		return false
	}
	if n == 1 {
		return true
	} else {
		log.Println("Issue with Customer Find: ", n)
		return false
	}

}

func AddCustomer(p *Customer) (bool, error) {
	if !findCustomer(p.CustomerID) {

		mcoll := GetCollection(CustomerCollName)
		err = mcoll.Insert(p)
		if err != nil {
			log.Println("Error while inserting the customer:\n", err)
			return false, err
		}
	}
	return true, nil
}
