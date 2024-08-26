package main

import "fmt"

func getCompletedName() (firstName, lastName string) {
	firstName = "Deniz"
	lastName = "Setiawan"

	return firstName, lastName
}

func main() {
	a, b := getCompletedName()

	fmt.Println(a)
	fmt.Println(b)
}