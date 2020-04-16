package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

var (
	timeout   = 30 * time.Second
	chromeAdd = "/usr/bin/chromium-browser"
)

// TackScreenShot take a screenshot form the given url
func TackScreenShot(url, name, root, extension string) (string, error) {
	path := root + "/" + name + "." + extension

	arg := []string{
		`--screenshot=` + path,

		"--headless", "--no-sandbox", "--disable-gpu",
		"--disable-dev-shm-usage",
		"--window-size=1920,1080",
		// Stop getting error
		"--disable-software-rasterizer",
		// Disable various background network services, including extension updating,
		//   safe browsing service, upgrade detector, translate, UMA
		"--disable-background-networking",
		// Disable installation of default apps on first run
		"--disable-default-apps",
		// Disable all chrome extensions entirely
		"--disable-extensions",
		// Disable syncing to a Google account
		"--disable-sync",
		// Disable chrome pop-up notifications which cover the page
		"--disable-notifications",
		// Disable built-in Google Translate service
		"--disable-translate",
		// Hide scrollbars on generated images/PDFs
		"--hide-scrollbars",
		// Disable reporting to UMA, but allows for collection
		"--metrics-recording-only",
		// Mute audio
		"--mute-audio",
		// Skip first run wizards
		"--no-first-run",
		// Disable fetching safebrowsing lists, likely redundant due to disable-background-networking
		"--safebrowsing-disable-auto-update",
		// set user data path
		"--user-data-dir=/data",
		url,
	}

	// get a context to run the command in
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, chromeAdd, arg...)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)

		return "", errors.New("Can not take a screeshot! timeout")
	}

	return path, nil
}
