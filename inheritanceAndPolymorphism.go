package main

import (
	"fmt"
)

type Human interface {
	printInfo()
}

func info(h Human) {
	h.printInfo()
}

type Employee struct {
	firstName string
	lastName  string
	salary    int
}

func (e Employee) printInfo() {
	fmt.Println(e.firstName, e.lastName, e.salary)
}

type Developer struct {
	Employee
	progLanguage string
}

func (d Developer) printInfo() {
	d.Employee.printInfo()
	fmt.Println(d.progLanguage)
}

type Manager struct {
	Employee
	workTime int
}

func (m Manager) printInfo() {
	m.Employee.printInfo()
	fmt.Println(m.workTime)
}

func main() {
	e := Employee{
		firstName: "Sabata",
		lastName:  "Abramishvili",
		salary:    3000,
	}
	m := Manager{
		Employee: Employee{"Elguja", "kvizhinadze", 3000},
		workTime: 8,
	}
	d := Developer{
		Employee:     Employee{"Avtandil", "chachanidze", 2000},
		progLanguage: "C++",
	}
	var arr [3]Human = [3]Human{}
	arr[0] = &e
	arr[1] = &m
	arr[2] = &d
	for i := 0; i < 3; i++ {
		info(arr[i])
	}
}
