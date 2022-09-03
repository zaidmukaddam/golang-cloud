package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func printUsage() {
	fmt.Println("Usage: " + os.Args[0] + " <filename>")
	fmt.Println("Example: " + os.Args[0] + " diskimage.iso")
}

func checkArgs() string {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	return os.Args[1]
}

func main() {
	filename := checkArgs()

	// Open file for reading
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create new hasher, which is a writer interface
	hasher := md5.New()

	// Default buffer size for copying is 32*1024 or 32kb per copy
	_, err = io.Copy(hasher, file)
	if err != nil {
		log.Fatal(err)
	}

	// Now get the final sum or checksum.
	checksum := hasher.Sum(nil)

	fmt.Printf("Md5 checksum: %x\n", checksum)
}
