package data

import (
	"log"

	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mongodb"
	"github.com/globalsign/mgo"
)

var mongoClient *mgo.Session
var err error

const (
	MongoDBName        = "goDB"
	CosmosDBConnString = "mongodb://gocosmos:hzD974FS34V4oUVaA47Pttbc8kYyo8AzNLSWQivZ2G8B9456hbSHOsAiB77rr1doXKbMoVdppQGE9ba81CJ2Wg==@gocosmos.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&maxIdleTimeMS=120000&appName=@gocosmos@"
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
