package data

import (
	"log"

	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mongodb"
	"github.com/globalsign/mgo"
)

var mongoClient *mgo.Session
var err error

const (
	MongoDBName        = "godb"
	CosmosDBConnString = "mongodb://gowithcosmos:8Va49WtynpHjpVyvwhiwcMTtuRVDt5Vbv9STwpiU7Hx28LMAeLIVl7CwBz8MEumiTA87r2ylMUnHjat0cI5dRg==@gowithcosmos.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&maxIdleTimeMS=120000&appName=@gowithcosmos@"
)

func init() {
	// Setting up MongoDB Client
	log.Println("Setting up Mongo Client")
	mongoClient, err = mongodb.NewMongoDBClientWithConnectionString(CosmosDBConnString)

	if err != nil {
		log.Fatalln("Error with the connection string:\n", err)
	}

}

func GetCollection(collname string) *mgo.Collection {
	return mongoClient.DB(MongoDBName).C(collname)
}
