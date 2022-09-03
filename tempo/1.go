package main

import (
	"fmt"
)

type employeee struct {
	ID            int
	designation   string
	yearOfBirth   int
	yearOfJoining int
	salary        float64
}

func e(employee employeee) {
	fmt.Println("Employee ID: ", employee.ID)
	fmt.Println("Employee designation: ", employee.designation)
	fmt.Println("Employee year of birth: ", employee.yearOfBirth)
	fmt.Println("Employee year of joining: ", employee.yearOfJoining)
	fmt.Println("Employee salary: ", employee.salary)
	println("")
}

func employeeInputs() (int, string, int, int, float64) {
	fmt.Println("Enter the employee ID: ")
	var empID int
	fmt.Scanln(&empID)
	fmt.Println("Enter the employee designation: ")
	var empDesignation string
	fmt.Scanln(&empDesignation)
	fmt.Println("Enter the employee year of birth: ")
	var empYearOfBirth int
	fmt.Scanln(&empYearOfBirth)
	fmt.Println("Enter the employee year of joining: ")
	var empYearOfJoining int
	fmt.Scanln(&empYearOfJoining)
	fmt.Println("Enter the employee salary: ")
	var empSalary float64
	fmt.Scanln(&empSalary)

	return empID, empDesignation, empYearOfBirth, empYearOfJoining, empSalary
}

func main() {
	var one, two, three, four, five employeee

	employees := []employeee{one, two, three, four, five}

	for _, employee := range employees {
		employee.ID, employee.designation, employee.yearOfBirth, employee.yearOfJoining, employee.salary = employeeInputs()
		e(employee)
	}
}
