package main

import "fmt"

func main() {
	name := "Deniz"

	switch name {
	case "Deniz":
		fmt.Println("Hello Deniz")
	case "Denis":
		fmt.Println("Hello Denis")
	default:
		fmt.Println("Hello World")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Terlalu Panjang")
	default:
		fmt.Println("OK")
	}

	length := len(name)

	switch {
	case length > 5:
		fmt.Println("Terlalu Panjang")
	default:
		fmt.Println("OK")
	}
}