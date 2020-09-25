package main

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
) // Site ...

// Site structure
type Site struct {
	URL              string    `json:"url"`
	ResponseCode     int       `json:"responseCode"`     // Latest response
	DurationInMs     int       `json:"durationInMs"`     // Latest duration
	Timestamp        time.Time `json:"timestamp"`        // Timestamp of sample taken
	RefreshRateInSec int       `json:"refreshRateInSec"` // How often to rescan a site
}

// Account structure
type Account struct {
	AccountName string          `json:"accounName"`
	URLList     map[string]Site `json:"URLList"`
}

// var account = new(Account)
var collection *mongo.Collection

// Global account data arrray. To be used with admin functionality
var accountData []*Account

// SitesMap Global site scan result map
var SitesMap = make(map[string]Site)

var c = make(chan string) // Initiate go routine channel. This needs to be globally accessible

func main() {

	initLogging()
	printVersion()                   // Print version
	loadConfiguration("config.yaml") // load configuration
	readEnv(&config)                 // Read environment attributes
	go scanScheduler()
	startServer() // Start HTTP server
}
