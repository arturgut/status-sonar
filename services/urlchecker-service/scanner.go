// This is URL scanner file
package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func scannerLoop() {

	log.Info("Starting URL scanner")

	for _, url := range config.Urls { // Main loop. For each item initiate go routine
		go urlScan(url, c)
	}

	for l := range c { // Infinite for loop to continusly scan URL listed in the channel 'c'
		go func(url string) {
			time.Sleep(config.Client.Period * time.Second)
			urlScan(url, c)
		}(l)
	}
}

func urlScan(url string, c chan string) { // Main urlScan function
	log.Info("Starting URL scan for: ", url)
	start := time.Now() // Measure time taken to establish http connection

	client := &http.Client{
		Timeout: config.Client.Timeout * time.Second,
	}

	resp, err := client.Get(url)
	t := time.Now()                            // Measure time taken to establish http connection
	elapsed := int(t.Sub(start)) / 1000 / 1000 // Convert to from date.tinme to miliseconds
	if err != nil {
		log.Warn("This URL is unreachable from here! ", err)
		scanResultsMap[url] = ScanResult{url, 404, 0, t} // Set HTTP STATUS CODE TO 404 and elapsed to '0'
		c <- url                                         // return result of scan the channel
		return
	}

	scanResultsMap[url] = ScanResult{url, resp.StatusCode, elapsed, t} // Add URL to global 'u' map
	log.Debug("Map updated for the following element: ", scanResultsMap[url], "\tCurrent map size: ", len(scanResultsMap))

	c <- url // return result of scan the channel
}
