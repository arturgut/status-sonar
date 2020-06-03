package main

import "time"

// ScanResult ...
type ScanResult struct {
	URL          string    `json:"url"`
	ResponseCode int       `json:"responseCode"` // also known as HTTP status code
	DurationInMs int       `json:"durationInMs"` // measured duration in ms
	Timestamp    time.Time `json:"timestamp"`    // measured duration in ms
}

// var initTimeNow = time.Now()

var scanResultsMap = make(map[string]ScanResult)

// { // initialise map with a single element
// 	"google.com": {URL: "http://google.com", ResponseCode: 200, DurationInMs: 232, },
//}
var c = make(chan string) // Initiate go routine channel. This needs to be globally accessible

func main() {

	initLogging()
	printVersion()                   // Print version
	loadConfiguration("config.yaml") // load configuration
	readEnv(&config)                 // Read environment attributes
	go scannerLoop()                 // Start URL scanner routines
	startServer()                    // Start HTTP server
}
