package main

import "fmt"

func goodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	contoh := goodBye
	fmt.Println(contoh("Deniz"))
}