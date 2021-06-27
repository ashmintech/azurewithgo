package data

import (
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mongodb"
	"github.com/globalsign/mgo"
)

var mongoClient *mgo.Session
var err error

const (
	MongoDBName = "godb"
)

func init() {
	// Setting up MongoDB Client
	log.Println("Setting up Mongo Client")

	CosmosDBConnString := os.Getenv("COSMOSDB_CONN_STRING")
	if CosmosDBConnString == "" {
		log.Fatalln("Database connection string empty")
	}

	mongoClient, err = mongodb.NewMongoDBClientWithConnectionString(CosmosDBConnString)

	if err != nil {
		log.Fatalln("Error with the connection string:\n", err)
	}

}

func GetCollection(collname string) *mgo.Collection {
	return mongoClient.DB(MongoDBName).C(collname)
}
