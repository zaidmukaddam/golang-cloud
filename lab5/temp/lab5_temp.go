package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

type Employee struct {
	EmployeeID       int
	Designation      string
	DateOfBirth      string
	DateOfJoining    string
	Salary           int
	DateOfRetirement string
}

func (e Employee) String() string {
	return fmt.Sprintf("%d %s %s %s %d %s", e.EmployeeID, e.Designation, e.DateOfBirth, e.DateOfJoining, e.Salary, e.DateOfRetirement)
}

type ByDate []time.Time

func UnderScore(s string) {
	for i := 0; i < 52; i++ {
		print("-")
	}
	fmt.Println()
	fmt.Println(s)
	for i := 0; i < 52; i++ {
		print("-")
	}
	println()
}

func AddEmployee() *[]Employee {
	e := []Employee{}
	for i := 0; i < 5; i++ {
		UnderScore("Employee " + strconv.Itoa(i+1))
		var EmployeeID int
		EmployeeID = i + 1
		fmt.Print("Enter Designation: ")
		var Designation string
		fmt.Scan(&Designation)
		fmt.Print("Enter DateOfBirth: ")
		var DateOfBirth string
		fmt.Scan(&DateOfBirth)
		fmt.Print("Enter DateOfJoining: ")
		var DateOfJoining string
		fmt.Scan(&DateOfJoining)
		fmt.Print("Enter Salary: ")
		var Salary int
		fmt.Scan(&Salary)
		e = append(e, Employee{EmployeeID, Designation, DateOfBirth, DateOfJoining, Salary, ""})
		print()
	}
	return &e
}

func SortEmployee(employee []Employee) {
	var by string
	UnderScore("Select -> ")

	fmt.Println("Which field do you want to sort by?")
	fmt.Println("1. Designation")
	fmt.Println("2. Age")
	fmt.Println("3. DateOfJoining")
	fmt.Println("4. Salary")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&by)

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

func CalculateDateOfRetirement(DateOfBirth string) string {
	const shortForm = "Jan/02/2006"
	t, _ := time.Parse(shortForm, DateOfBirth)
	t = t.AddDate(60, 0, 0)
	return t.Format(shortForm)
}

func DisplayEmployeeRetirement(employee []Employee) {
	for i := 0; i < len(employee); i++ {
		employee[i].DateOfRetirement = CalculateDateOfRetirement(employee[i].DateOfBirth)
	}
	DisplayEmployee(employee)
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
	SortEmployee(employee)
	UnderScore("Sorted Employee")
	DisplayEmployee(employee)
	UnderScore("Changed Salary wirh Date of Retirement")
	DisplayEmployeeRetirement(employee)
}
