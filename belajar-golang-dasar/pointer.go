package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main () {
	// address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// address2 := address1 // copy value

	// address2.City = "Bandung"
	// fmt.Println(address1) // tidak berubah
	// fmt.Println(address2)

	var address1 Address = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	var address2 *Address = &address1 //pointer

	address2.City = "Bandung"
	fmt.Println(address1) //berubah
	fmt.Println(address2)
}