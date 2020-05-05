package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
)

// WriteFile return a file as response
func WriteFile(w http.ResponseWriter, path string) error {

	//Check if file exists and open
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	//Close after function return
	defer file.Close()

	//Get info from file
	FileStat, _ := file.Stat()
	fileName := file.Name()
	//Get file size as a string
	fileSize := strconv.FormatInt(FileStat.Size(), 10)

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fileSize)

	//We read 512 bytes from the file already, so we reset the offset back to 0
	file.Seek(0, 0)
	io.Copy(w, file) //'Copy' the file to the client

	return nil
}
