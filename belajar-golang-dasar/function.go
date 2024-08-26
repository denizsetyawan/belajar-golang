package main

import "fmt"

func main()  {
	sayHello("Deniz")
	result := getHello("Deniz")

	fmt.Println(result)
}

func sayHello(name string)  {
	fmt.Println("Hello", name)
}

// function return value
func getHello(name string) string  {
	return "Hello " + name
}

// function return multiple values
func getHello2(name string) (string, string)  {
	return "Hello ", name
}