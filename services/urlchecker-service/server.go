package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func reloadConfig(w http.ResponseWriter, r *http.Request) {
	log.Debug("In reload config")
	if r.URL.Path != "/api/reload-config" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {

	case "GET":
		enableCors(&w)
		log.Debug("Reload data")
		loadAccountsData()
		w.Write([]byte("Configuration has been reloaded"))
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

}

func metrics(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "# HELP Version: 0.1 Alpha - https://github.com/arturgut/urlchecker\n")
	fmt.Fprintf(w, "# HELP url_checker. Label: HTTP Response code. Value: Request duration in Ms\n")
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	log.Debug("Hostname:", name)
	fmt.Fprintf(w, "# Hostname: %v \n", name)

	for key, value := range SitesMap {
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
	// w.Write([]byte(mapToJSONv2()))
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
			delete(SitesMap, k)
			log.Info("remove(): Removed item ", SitesMap, "from map.\tCurrent map items: ", len(SitesMap))
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

			if val, ok := SitesMap[k]; ok { // Check if
				//do something here
				log.Info("URL already known. Skipping. URL: ", SitesMap[k], val)
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
				log.Debug("Elements in map:", len(SitesMap))
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
	http.HandleFunc("/api/reload-config", reloadConfig)

	serverPort := ":" + strconv.FormatInt(int64(config.Server.Port), 10) // Format server port to be type string
	http.ListenAndServe(serverPort, nil)                                 // Start server
}

func mapToJSON() []byte {

	var sliceOfSite []Site
	for _, value := range SitesMap { // Convert Map of structs to slice
		sliceOfSite = append(sliceOfSite, value)
	}

	data, err := json.MarshalIndent(sliceOfSite, "", "   ") // Print JSON
	if err != nil {
		fmt.Println("Error during marshalling")
	}
	fmt.Printf("%s\n", data)
	return data
}

func mapToJSONv2() []byte {

	for _, elm := range accountData { // iterate over all accounts and load list of URLs
		log.Debug("scannerRoutine(), Loading account data: ", elm.AccountName)
		for _, url := range elm.URLList {
			log.Debug("scannerRoutine(), Loading url data: ", url.URL)
			go urlScan(url.URL, c)
		}
	}

	return nil
}
