package main

import "fmt"

func main() {
	var myText = "test string 1"

	myText2 := "test string 2"

	myText3 := `long string
spanning multiple
lines`

	fmt.Println(myText)
	fmt.Println(myText2)
	fmt.Println(myText3)

}
