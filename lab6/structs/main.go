package main

import "fmt"

func main() {
	type Person struct {
		Name string
		Age  int
	}

	nanodano := Person{Name: "Zaid", Age: 19}
	fmt.Println(nanodano)

	type Hacker struct {
		Person  Person
		Favlang string
	}

	fmt.Println(nanodano)

	hacker := Hacker{
		Person:  nanodano,
		Favlang: "Go",
	}

	fmt.Println(hacker)
	fmt.Println(hacker.Person.Name)
}
