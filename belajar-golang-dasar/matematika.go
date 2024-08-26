package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = a + b

	var i = 10
	i += 10
	i += 5

	fmt.Println(i)
	fmt.Println(c)
}