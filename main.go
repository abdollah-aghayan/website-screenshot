package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func screenshot(w http.ResponseWriter, r *http.Request) {
	// Check request type
	if r.Method != "GET" {
		http.Error(w, "Recivied non GET request", 404)

		return
	}

	// check url
	u := r.URL.Query().Get("url")
	if u == "" {
		http.Error(w, "url is required", 400)
		return
	}

	// validate received url
	_, err := url.ParseRequestURI(u)
	if err != nil {
		http.Error(w, "Please pass valid url", 400)

		return
	}

	// create a random name for image
	name := generateRandomKey(keyLength)

	// take screenshout
	filePath, err := TackScreenShot(u, name, imageFolder, fileExtention)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Write image to response
	err = WriteFile(w, filePath)
	if err != nil {
		http.Error(w, "internal error", 500)
	}

	err = Remove(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return

}

func main() {
	// Set route
	http.HandleFunc("/screenshot", screenshot)

	// load env
	t := os.Getenv("TIMEOUT")
	if t != "" {
		value, err := strconv.Atoi(t)
		if err == nil {
			timeout = time.Duration(value) * time.Second
		}
	}

	// Start server on specified port
	fmt.Println("Starting server...")
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
