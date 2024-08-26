package main

import "fmt"

func main() {
	var names [4]string

	names[0] = "Deniz"
	names[1] = "Denis"
	names[2] = "Deni"
	names[3] = "Denny"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])

	var values = [...]int{1, 2, 3}
	fmt.Println(values)
	fmt.Println(len(values))
}