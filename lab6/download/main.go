package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// download a file
func main() {
	nameFile, err := os.Create("download/devdungeon.html")
	if err != nil {
		log.Fatal(err)
	}

	defer nameFile.Close()

	url := "https://www.devdungeon.com/archive"
	
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	numBytes, err := io.Copy(nameFile, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Downloaded %d bytes\n", numBytes)
}
