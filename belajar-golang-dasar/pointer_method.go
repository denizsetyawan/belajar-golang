package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) sayHello() {
	man.Name = "Mr. " + man.Name
}

func main() {
	deniz := Man{"deniz"}
	deniz.sayHello()
	fmt.Println(deniz)
}