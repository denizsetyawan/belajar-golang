package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("deniz")

	if data == nil {
		fmt.Println("data is nil")
	} else {
		fmt.Println(data["name"])
	}
}