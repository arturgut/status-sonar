package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func metrics(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "# HELP Version: 0.1 Alpha - https://github.com/arturgut/urlchecker\n")
	fmt.Fprintf(w, "# HELP url_checker. Label: HTTP Response code. Value: Request duration in Ms\n")
	for key, value := range scanResultsMap {
		// Prometheus format: <metric name>{<label name>=<label value>, ...}
		fmt.Fprintf(w, "url_checker{ url='%v', http_status_code=%v } %v\n", key, value.ResponseCode, value.DurationInMs)
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	log.Debug("In /list handler")
	enableCors(&w) // enableCors(&w)

	if r.URL.Path != "/api/list" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte(mapToJSON()))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func remove(w http.ResponseWriter, r *http.Request) {
	log.Debug("In /remove handler")

	if r.URL.Path != "/api/remove" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {

	case "GET":
		for k, v := range r.URL.Query() {
			log.Debug("remove(): Received a GET request", k, v)

			// Remove URL new url here
			delete(scanResultsMap, k)
			log.Info("remove(): Removed item ", scanResultsMap, "from map.\tCurrent map items: ", len(scanResultsMap))
			w.Write([]byte("Item has been removed!"))
			w.Write([]byte(mapToJSON()))
		}
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}

func add(w http.ResponseWriter, r *http.Request) {
	log.Debug("In /add handler")

	if r.URL.Path != "/api/add" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {

	case "GET":
		for k, v := range r.URL.Query() {
			log.Debug("Received a GET request", k, v)
			url, err := url.Parse(k) // URL Parse

			if val, ok := scanResultsMap[k]; ok { // Check if
				//do something here
				log.Info("URL already known. Skipping. URL: ", scanResultsMap[k], val)
				w.Write([]byte("URL already known. Skipping"))
				return
			}

			if err != nil {
				log.Error("There something wrone with this URL", err)
				return
			}
			if url.Scheme == "http" || url.Scheme == "https" { // Run some basic URL parsing
				urlScan(k, c)
				w.Write([]byte("Site has ben successfully added to the list of URL's\n"))
				w.Write([]byte(mapToJSON()))
				log.Info(k, "URL has ben successfully added to the list of URL's\n")
				log.Debug("Elements in map:", len(scanResultsMap))
			} else {
				w.Write([]byte(url.Scheme + " is not at valid scheme. Expecting http or https!"))
				log.Error("URL has no valid scheme. Expecting http or https!")
				return
			}

		}
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}

func startServer() {
	log.Info("INFO: Starting HTTP server on port:", config.Server.Port, ".\n You should be able to access the URL at http://localhost:", config.Server.Port)
	http.HandleFunc("/metrics", metrics)
	http.HandleFunc("/api/add", add)
	http.HandleFunc("/api/remove", remove)
	http.HandleFunc("/api/list", list)

	serverPort := ":" + strconv.FormatInt(int64(config.Server.Port), 10) // Format server port to be type string
	http.ListenAndServe(serverPort, nil)                                 // Start server
}

func mapToJSON() []byte {

	var sliceOfScanResult []ScanResult
	for _, value := range scanResultsMap { // Convert Map of structs to slice
		sliceOfScanResult = append(sliceOfScanResult, value)
	}

	data, err := json.MarshalIndent(sliceOfScanResult, "", "   ") // Print JSON
	if err != nil {
		fmt.Println("Error during marshalling")
	}
	fmt.Printf("%s\n", data)
	return data
}
