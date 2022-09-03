package main

import (
	"fmt"
	"sort"
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

func AddEmployee() *[]Employee {
	e := []Employee{}
	e = append(e, Employee{1, "Manager", "Jan/02/1992", "Aug/10/2020", 100000, ""})
	e = append(e, Employee{2, "Developer", "Dec/09/1995", "Mar/07/2017", 50000, ""})
	e = append(e, Employee{3, "Tester", "Jun/12/1991", "May/15/2014", 30000, ""})
	e = append(e, Employee{4, "Designer", "Oct/19/1990", "Feb/12/2021", 20000, ""})
	e = append(e, Employee{5, "Tester", "Aug/29/1990", "Jan/16/2016", 30000, ""})
	return &e
}

func UnderScore(s string) {
	for i := 0; i < 52; i++ {
		print("-")
	}
	fmt.Println()
	println(s)
	for i := 0; i < 52; i++ {
		print("-")
	}
	println()
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
	ChangeSalary(employee)
	UnderScore("Sorted Employee")
	DisplayEmployee(employee)
	UnderScore("Changed Salary wirh Date of Retirement")
	DisplayEmployeeRetirement(employee)
}
