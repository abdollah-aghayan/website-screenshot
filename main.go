package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

const (
	keyLength     = 20
	imageFolder   = "/tmp"
	fileExtention = "png"
)

func screenshot(w http.ResponseWriter, r *http.Request) {
	// Check request type
	if r.Method != "GET" {
		http.Error(w, "Recivied non GET request", 404)

		return
	}

	u := r.URL.Query().Get("url")
	if u == "" {
		http.Error(w, "url is required", 400)

		return
	}

	_, err := url.ParseRequestURI(u)
	if err != nil {
		http.Error(w, "Please pass valid url", 400)

		return
	}

	name := generateRandomKey(keyLength)

	filePath, err := TackScreenShot(u, name, imageFolder, fileExtention)

	if err != nil {
		http.Error(w, err.Error(), 500)

		return
	}

	err = Writ(w, filePath)
	if err != nil {
		http.Error(w, "internal error", 500)
	}

	return

}

func main() {
	http.HandleFunc("/screenshot", screenshot)

	// load env
	t := os.Getenv("TIMEOUT")
	if t != "" {
		value, err := strconv.Atoi(t)
		if err == nil {
			timeout = time.Duration(value) * time.Second
		}
	}

	fmt.Println("Starting server...")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
