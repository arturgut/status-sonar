// This is URL scanner file
package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func scanScheduler() {
	log.Info("Starting Scan Scheduler routine")

	for {
		log.Info("Loading account data")
		loadAccountsData()
		log.Debug("Number of routines: ", len(c))
		for _, elm := range accountData { // iterate over all accounts and load list of URLs
			log.Debug("scannerRoutine(), Loading account data: ", elm.AccountName)
			for _, url := range elm.URLList {
				log.Debug("scannerRoutine(), Loading url data: ", url.URL)
				go urlScan(url.URL, c)
			}
		}

		time.Sleep(config.Client.Period * time.Second)
	}

}

// func scannerRoutine() {
// 	log.Info("Starting URL scanner routine")

// 	for _, elm := range accountData { // iterate over all accounts and load list of URLs
// 		log.Debug("scannerRoutine(), Loading account data: ", elm.AccountName)
// 		for _, url := range elm.URLList {
// 			log.Debug("scannerRoutine(), Loading url data: ", url.URL)
// 			go urlScan(url.URL, c)
// 		}
// 	}

// 	for l := range c { // Infinite for loop to continusly scan URL listed in the channel 'c'
// 		go func(url string) {
// 			time.Sleep(config.Client.Period * time.Second)
// 			urlScan(url, c)
// 		}(l)
// 	}
// }

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
		SitesMap[url] = Site{url, 404, 0, t, 60} // Set HTTP STATUS CODE TO 404 and elapsed to '0'
		c <- url                                 // return result of scan the channel
		return
	}

	SitesMap[url] = Site{url, resp.StatusCode, elapsed, t, SitesMap[url].RefreshRateInSec} // Add URL to global 'u' map
	log.Debug("Map updated for the following element: ", SitesMap[url], "\tCurrent map size: ", len(SitesMap))

	c <- url // return result of scan the channel
}
