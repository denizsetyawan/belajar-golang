package main

import "fmt"

func main() {
	name := "Denis"

	if name == "Deniz" {
		fmt.Println("Hello Deniz")
	} else if name == "Denis" {
		fmt.Println("Hello Denis")
	} else {
		fmt.Println("Hello World")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("OK")
	}
}