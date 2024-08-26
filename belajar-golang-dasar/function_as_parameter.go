package main

import "fmt"

type Filter func(string) string
// func sayHellowithFilter(name string, filter func(string) string) {
func sayHellowithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHellowithFilter("Denis", spamFilter)

	filter := spamFilter
	sayHellowithFilter("anjing", filter)
}