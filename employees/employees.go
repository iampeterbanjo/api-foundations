package main

import "fmt"

type Employee interface {
	getName() string
	getSalary() int
}

type Personnel struct {
	firstName string
	lastName  string
	salary    int
}

func (p Personnel) getSalary() int {
	return p.salary
}

func (p Personnel) getName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

type Receptionist struct {
	Personnel
	petty_cash int
}

func Review(e Employee) string {
	return fmt.Sprintf("Hi, my name is %v and I make %v a year", e.getName(), e.getSalary())
}

func main() {
	r := Receptionist{}
	r.firstName = "Money"
	r.lastName = "Penny"
	fmt.Println(Review(r))
}
