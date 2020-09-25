package main

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func loadAccountsData() {

	log.Debug("len of accountData: ", len(accountData))
	accountData = nil // flush account data array
	log.Debug("len of accountData: ", len(accountData))

	findOptions := options.Find()
	findOptions.SetLimit(1000)

	c := ConnectDatabase()

	filter := bson.D{{}}

	collection = c.Database("status-sonar").Collection("accounts")
	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Account
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		log.Debug("Loading account: ", elem.AccountName)
		accountData = append(accountData, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	log.Debug("Found multiple documents (array of pointers): ", accountData)

}

func findAccount() {}

// ConnectDatabase - connect to mongodb
func ConnectDatabase() *mongo.Client {

	// Set client options, use configuration yaml or env vars
	connectionString := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d",
		config.DbClient.Username,
		config.DbClient.Password,
		config.DbClient.Host,
		config.DbClient.Port,
	)
	log.Debug("ConnectDatabase(): Connection String to MongoDB:  ", connectionString)

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions) // Connect to MongoDB
	if err != nil {
		log.Fatal(err)
	}

	return client
}

// Helper function
func isAccountSet(accountName string) bool {

	var account Account

	log.Debug("isAccountSet(), accountName: ", accountName)
	c := ConnectDatabase()

	filter := bson.D{{"accountname", accountName}}
	log.Debug("isAccountSet() => filter: ", filter)
	collection = c.Database("status-sonar").Collection("accounts")
	response := collection.FindOne(context.TODO(), filter).Decode(&account)
	log.Debug("isAccountSet(): response from MongoDB: ", response)

	if response != nil {

		log.Debug("isAccountSet(): Found 0 documents for account: ", account)
		return false

	} else {
		log.Debug("isAccountSet(): Found a single document: ", account)
		return true
	}

}

func dbUpdateAccount(a Account) interface{} {

	log.Debug("dbAddAccount(), accountName: ", a)
	c := ConnectDatabase()

	filter := bson.D{{"accountname", a.AccountName}}
	collection = c.Database("status-sonar").Collection("accounts")

	updateBson := bson.D{{Key: "$set", Value: a}}

	updatedResult, err := collection.UpdateOne(context.TODO(), filter, updateBson)
	if err != nil {
		log.Fatal("Error on updating one Hero", err)
	}
	return updatedResult.ModifiedCount

}

func dbGetAccount(accountName string) Account {

	log.Debug("dbAddAccount(), accountName: ", accountName)
	c := ConnectDatabase()

	var account Account

	filter := bson.D{{"accountname", accountName}}
	log.Debug("dbGetAccount() => filter: ", filter)

	collection = c.Database("status-sonar").Collection("accounts")

	response := collection.FindOne(context.TODO(), filter).Decode(&account)
	log.Debug("dbGetAccount(): response from MongoDB: ", response)
	if response != nil {
		log.Debug("dbGetAccount(): Found 0 results: ", account)
	}

	log.Debug("dbGetAccount(): Found a single document: ", account)

	return account

}

func dbAddAccount(a Account) interface{} {

	log.Debug("dbAddAccount(): Recived following Account variable", a)
	log.Debug("dbAddAccount(): Account name:", a.AccountName)

	if isAccountSet(a.AccountName) {

		log.Debug("dbAddAccount(): Account already set: ", a.AccountName)
		return nil

	} else {

		log.Debug("dbAddAccount(): Account not defined: ", a.AccountName)

		c := ConnectDatabase()

		err := c.Ping(context.Background(), readpref.Primary())
		if err != nil {
			log.Fatal(err)
		} else {
			log.Info("Connected!")
		}

		// Insert a document
		collection = c.Database("status-sonar").Collection("accounts")

		insertResult, err := collection.InsertOne(context.TODO(), a)
		if err != nil {
			log.Fatal(err)
		}

		log.Debug("dbAddAccount(): Inserted a single document: ", insertResult.InsertedID)

		return insertResult.InsertedID
	}
}
