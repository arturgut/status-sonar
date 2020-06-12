package main

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

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

func main() {

	initLogging()                    // initialse
	printVersion()                   // Print version
	loadConfiguration("config.yaml") // load configuration
	readEnv(&config)                 // Read environment attributes
	startServer()                    // Start HTTP server
}
