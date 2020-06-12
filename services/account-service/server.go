package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

func updateAccountHandler(w http.ResponseWriter, r *http.Request) {

	var a Account

	log.Debug("updateAccountHandler(): ")
	if r.URL.Path != "/api/account/update" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	case "POST":
		log.Debug("updateAccountHandler(): CASE POST: ", r.PostForm)
		enableCors(&w)
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Debug("Post body: ", reqBody)

		var e = json.Unmarshal(reqBody, &a)
		if err != nil {
			log.Fatal(e)
		}
		log.Debug("Unmarshalled JSON Body: ", a)
		log.Debug("Account name: ", a.AccountName)
		dbUpdateAccount(a)

	case "OPTIONS":
		log.Debug("updateAccountHandler(): CASE OPTIONS: ", r.PostForm)
		enableCors(&w)
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Debug("Post body: ", reqBody)
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}

// add new account
func addAccountHandler(w http.ResponseWriter, r *http.Request) {

	var a Account
	a.URLList = make(map[string]ScanResult)
	t := time.Now()
	log.Debug("addAccountHandler(): ")
	enableCors(&w) // enable Cors support
	if r.URL.Path != "/api/account/add" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		log.Debug("addAccountHandler(): Received a GET request: ", r.URL.Query())
		for k, v := range r.URL.Query() {
			log.Debug("addAccountHandler(): GET request Key ", k)
			log.Debug("addAccountHandler(): GET request Value:", v[0])
			a.AccountName = string(v[0])
			a.URLList["google.com"] = ScanResult{"http://google.com", 200, 123, t, 60} // Add default scan result entry
			dbAddAccount(a)
			defer r.Body.Close()
			w.WriteHeader(http.StatusOK)
		}

	// TODO - Implement POST
	case "POST":
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", reqBody)
		w.Write([]byte("Received a POST request\n"))
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
	// w.Write([]byte(mapToJSON(a)))
}

func metrics(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "# HELP Version: 0.1 Alpha - https://github.com/arturgut/urlchecker\n")
	fmt.Fprintf(w, "# HELP account service. Label: HTTP Response code. Value: Request duration in Ms\n")
}

func listAccountURLsHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("listAccountHandler(): ")
	enableCors(&w) // enable Cors support
	if r.URL.Path != "/api/account/list" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		for _, v := range r.URL.Query() {
			a := dbGetAccount(v[0])
			if a.AccountName != "" {
				log.Debug("a.AccountName not empty: ", a.AccountName)
				m := a.URLList
				var sliceOfScanResult []ScanResult
				for _, value := range m { // Convert Map of structs to slice
					sliceOfScanResult = append(sliceOfScanResult, value)
				}
				d, err := json.MarshalIndent(sliceOfScanResult, "", "   ")
				if err != nil {
					fmt.Println("Error during marshalling")
				}
				w.Write([]byte(d))
				w.WriteHeader(http.StatusOK)
			} else {
				log.Debug("a.accountname is empty ", a.AccountName)
				w.WriteHeader(http.StatusPreconditionFailed)
				w.Write([]byte("Account not found"))
			}
		}
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

}

func listAccountDetailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("listAccountHandler(): ")
	enableCors(&w) // enable Cors support

	if r.URL.Path != "/api/account/get" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {

	case "GET":
		for _, v := range r.URL.Query() {
			a := dbGetAccount(v[0])
			if a.AccountName != "" {
				log.Debug("a.AccountName not empty: ", a.AccountName)
				d, err := json.MarshalIndent(a, "", "   ")
				if err != nil {
					fmt.Println("Error during marshalling")
				}
				w.Write([]byte(d))
				w.WriteHeader(http.StatusOK)
			} else {
				log.Debug("a.accountname is empty ", a.AccountName)
				w.WriteHeader(http.StatusPreconditionFailed)
				w.Write([]byte("Account not found"))
			}
		}
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}

func startServer() {
	log.Info("INFO: Starting HTTP server on port:", config.Server.Port, ".\n You should be able to access the URL at http://localhost:", config.Server.Port)
	http.DefaultClient.Timeout = time.Minute * 10
	http.HandleFunc("/metrics", metrics)
	http.HandleFunc("/api/account/list", listAccountURLsHandler)
	http.HandleFunc("/api/account/get", listAccountDetailsHandler)
	http.HandleFunc("/api/account/add", addAccountHandler)
	http.HandleFunc("/api/account/update", updateAccountHandler)

	serverPort := ":" + strconv.FormatInt(int64(config.Server.Port), 10) // Format server port to be type string
	http.ListenAndServe(serverPort, nil)                                 // Start server
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")

}

func mapToJSON(a Account) []byte {

	d, err := json.MarshalIndent(a, "", "   ") // Print JSON
	if err != nil {
		fmt.Println("Error during marshalling")
	}
	fmt.Printf("%s\n", d)
	return d
}
