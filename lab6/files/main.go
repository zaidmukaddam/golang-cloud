package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	nameFile, err := os.Create("files/main.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(nameFile.Name())
	nameFile.Close()
	stat()
	delFile()
}

func stat() {
	file, err := os.Stat("files/main.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File name:", file.Name())
	fmt.Println("Size in bytes:", file.Size())
	fmt.Println("Permissions:", file.Mode())
	fmt.Println("Last modified:", file.ModTime())
	fmt.Println("Is Directory:", file.IsDir())
	fmt.Println("Is Regular File:", file.Mode().IsRegular())
	fmt.Printf("System interface type: %T\n", file.Sys())
	fmt.Printf("System info: %+v\n\n", file.Sys())
}

func copyFile() {

}

func delFile() {
	err := os.Remove("files/main.txt")
	if err != nil {
		log.Fatal(err)
	}
}
