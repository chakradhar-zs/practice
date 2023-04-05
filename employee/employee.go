package main

import (
	"strconv"
	"time"
)

type employee struct {
	firstname, lastname, dob string
}

var emp = employee{}

func GreetEmployee(f, l string) (s string) {

	emp.firstname, emp.lastname = f, l
	if emp.lastname == "" {
		s = "Hello, " + emp.firstname
	} else {
		s = "Hello, " + emp.firstname + " " + emp.lastname
	}

	return
}

func CalculateAge(dob string) (age int) {
	emp.dob = dob
	l := len(dob)
	yr, _ := strconv.Atoi(dob[l-4 : l])
	return time.Now().Year() - yr
}
