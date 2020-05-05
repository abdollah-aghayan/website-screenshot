package main

import "os"

// Remove remove a file
func Remove(path string) error {
	err := os.Remove(path)

	if err != nil {
		return err
	}

	return nil
}
