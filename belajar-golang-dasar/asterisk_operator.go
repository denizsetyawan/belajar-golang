package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main () {
	var address1 Address = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	var address2 *Address = &address1 //pointer

	address2.City = "Bandung"
	fmt.Println(address1) //berubah
	fmt.Println(address2)

	// address2 = &Address{"Surabaya", "Jawa Timur", "Indonesia"}
	*address2 = Address{"Surabaya", "Jawa Timur", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}