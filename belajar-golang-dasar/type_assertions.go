package main

import "fmt"

func random () any {
	return "ok"
}

func main() {
	var result any = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("unknown", value)
	}
}