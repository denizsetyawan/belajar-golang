package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (c Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", c.Name)
}

func main() {
	var deniz Customer
	deniz.Name = "Deniz"
	deniz.Address = "jogja"
	deniz.Age = 22
	fmt.Println(deniz)
	fmt.Println(deniz.Name)

	deni := Customer{
		Name:    "Denis",
		Address: "bantul",
		Age:     23,
	}
	
	denny := Customer{"Denny", "solo", 24}

	fmt.Println(deni)
	fmt.Println(denny)

	deniz.sayHello("Denis")
}