package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 3; i++ {
		fmt.Println("i: ", i)
	}

	m := 5

	for m < 10 {
		fmt.Println("m: ", m)
		m++
	}
}
