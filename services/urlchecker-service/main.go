package main

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
) // ScanResult ...

// ScanResult structure
type ScanResult struct {
	URL              string    `json:"url"`
	ResponseCode     int       `json:"responseCode"`     // Latest response
	DurationInMs     int       `json:"durationInMs"`     // Latest duration
	Timestamp        time.Time `json:"timestamp"`        // Timestamp of sample taken
	RefreshRateInSec int       `json:"refreshRateInSec"` // How often to rescan a site
}

// Account structure
type Account struct {
	AccountName string                `json:"accounName"`
	URLList     map[string]ScanResult `json:"URLList"`
}

var account = new(Account)
var collection *mongo.Collection

var accountData []*Account // Global account data arrray

var scanResultsMap = make(map[string]ScanResult)

var c = make(chan string) // Initiate go routine channel. This needs to be globally accessible

func main() {

	initLogging()
	printVersion()                   // Print version
	loadConfiguration("config.yaml") // load configuration
	readEnv(&config)                 // Read environment attributes
	loadAccountsData()               // Load acccounts data
	go scannerRoutine()              // Start URL scanner routines
	startServer()                    // Start HTTP server
}
