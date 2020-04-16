package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
)

func Writ(w http.ResponseWriter, path string) error {

	//Check if file exists and open
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close() //Close after function return

	//Get the file size
	FileStat, _ := file.Stat() //Get info from file
	fileName := file.Name()
	fileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fileSize)

	//Send the file

	//We read 512 bytes from the file already, so we reset the offset back to 0
	file.Seek(0, 0)
	io.Copy(w, file) //'Copy' the file to the client

	return nil
}

func remove(path string) error {
	err := os.Remove(path)

	if err != nil {
		return err
	}

	return nil
}
