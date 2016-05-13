package main

import (
	"net/http"

	"log"
	"fmt"
	"net/url"
	"time"
)

func main() {

	port := "8080"
	log.Print("Starting shortener, listening at :", port)

	// register health check
	launch := time.Now()
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Print("hit healthcheck endpoint\n")
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintf(w, `{ "name":"URL Shortener", "version":"v0.1", "port":"%s", "started":"%s"}`, port, launch.Format(time.RFC3339))
	})

	// default behavior see below
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}


// This is the URL shortener : redirect an URL to another
// To be invoked as http://domain:port/?url=XXXXXXX
// example :  http://localhost:8080/?url=spark://rooms/4131c373-81db-3731-837a-3d1eedcf76a3
func handler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		log.Print("Expecting GET method as I am an URL shortener")
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-type", "application/json")
		fmt.Fprintf(w, `{ "message":"I am an URL shortener, expecting a GET" }`)
		return
	}

	log.Print("Parsing URL :", req.RequestURI)
	params, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		log.Print("Could not parse URL: ", req.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-type", "application/json")
		fmt.Fprintf(w, `{ "message":"Could not parse URL" }`)
		return
	}
	log.Print("Parsed")
	urls := params["url"];
	if urls == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-type", "application/json")
		log.Print(`{ "message":"No url found to shorten, please specify url query parameter with url?" }`)
		fmt.Fprintf(w, `{ "message":"url parameter not found" }`)
		return
	}
	url := urls[0]
	log.Print("Url: " + url)
	if (url == "") {
		log.Print("Could not find query parameter 'url' in:", req.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-type", "application/json")
		fmt.Fprintf(w, `{ "message":"No url found to shorten, please specify url query parameter with url?" }`)
		return
	}


	// URL redirect
	log.Print("redirecting to " + url);
	http.Redirect(w, req, url, http.StatusMovedPermanently)

	return
}
