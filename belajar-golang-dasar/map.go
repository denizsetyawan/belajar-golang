package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}

	// person["name"] = "Deniz"
	// person["address"] = "jogja"

	person := map[string]string{
		"name":    "Deniz",
		"address": "jogja",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Deniz Setiawan"
	book["ups"] = "Salah"

	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)
}