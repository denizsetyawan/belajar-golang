package main

import "fmt"

func getFullName() (string, string) {
	return "Deniz", "Setiawan"
}

func main()  {
	// firstName, lastName := getFullName()
	firstName, _ := getFullName()

	fmt.Println(firstName)
	// fmt.Println(lastName)
}