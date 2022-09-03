package main

import (
	"fmt"
	"sort"
	"time"
)

/*
1.	Write a program include functions
a.	To add five employees (include EmployeeID, Designation, Date of Birth, Date of Joining, Salary) in your company
b.	To sort employees as per designation, as per age, as per date of joining
c.	To change salary of few employees across all designations
d.	To display list of employees along with date when they will retire (Age of retirement to be considered as 60 years)

*/

type Employee struct {
	EmployeeID    int
	Designation   string
	DateOfBirth   string
	DateOfJoining string
	Salary        int
}

type ByDate []time.Time

func (e Employee) String() string {
	return fmt.Sprintf("%d %s %s %s %d", e.EmployeeID, e.Designation, e.DateOfBirth, e.DateOfJoining, e.Salary)
}

func AddEmployee() *[]Employee {
	e := []Employee{}
	for i := 0; i < 5; i++ {
		fmt.Println("Enter EmployeeID: ")
		var EmployeeID int
		fmt.Scanln(&EmployeeID)
		fmt.Println("Enter Designation: ")
		var Designation string
		fmt.Scanln(&Designation)
		fmt.Println("Enter DateOfBirth: ")
		var DateOfBirth string
		fmt.Scanln(&DateOfBirth)
		fmt.Println("Enter DateOfJoining: ")
		var DateOfJoining string
		fmt.Scanln(&DateOfJoining)
		fmt.Println("Enter Salary: ")
		var Salary int
		fmt.Scanln(&Salary)
		e = append(e, Employee{EmployeeID, Designation, DateOfBirth, DateOfJoining, Salary})
	}
	return &e
}

func SortEmployee(employee []Employee, by string) {
	switch by {
	case "Designation":
		sort.Slice(employee, func(i, j int) bool {
			return employee[i].Designation < employee[j].Designation
		})
	case "Age":
		sort.Slice(employee, func(i, j int) bool {
			return CalculateAge(employee[i].DateOfBirth).Before(CalculateAge(employee[j].DateOfBirth))
		})
	case "DateOfJoining":
		sort.Slice(employee, func(i, j int) bool {
			return CalculateDateofJoining(employee[i].DateOfJoining).Before(CalculateDateofJoining(employee[j].DateOfJoining))
		})
	case "Salary":
		sort.Slice(employee, func(i, j int) bool {
			return employee[i].Salary < employee[j].Salary
		})
	default:
		fmt.Println("Invalid Sort")
	}
}

func CalculateAge(DateOfBirth string) time.Time {
	const shortForm = "Jan/02/2006"
	t, _ := time.Parse(shortForm, DateOfBirth)
	return t
}

func CalculateDateofJoining(DateOfJoining string) time.Time {
	const shortForm = "Jan/02/2006"
	t, _ := time.Parse(shortForm, DateOfJoining)
	return t
}

func ChangeSalary(employee []Employee) {
	for i := 0; i < len(employee); i++ {
		if employee[i].Designation == "Developer" {
			employee[i].Salary = employee[i].Salary + 10000
		} else if employee[i].Designation == "Tester" {
			employee[i].Salary = employee[i].Salary + 5000
		} else if employee[i].Designation == "Designer" {
			employee[i].Salary = employee[i].Salary + 3000
		} else if employee[i].Designation == "Manager" {
			employee[i].Salary = employee[i].Salary + 20000
		} else {
			fmt.Println("Invalid Designation")
		}
	}
}

func DisplayEmployee(employee []Employee) {
	for _, value := range employee {
		fmt.Println(value)
	}
}

func main() {
	employee := []Employee{}
	employee = *AddEmployee()
	DisplayEmployee(employee)
	fmt.Println("Which field do you want to sort by?")
	fmt.Println("1. Designation")
	fmt.Println("2. Age")
	fmt.Println("3. DateOfJoining")
	fmt.Println("4. Salary")
	var by string
	fmt.Scanln(&by)
	SortEmployee(employee, by)
	DisplayEmployee(employee)
	ChangeSalary(employee)
	println("Changed Salary")
	DisplayEmployee(employee)
}
