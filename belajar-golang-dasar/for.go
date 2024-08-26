package main

import "fmt"

func main() {
	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke", i)
	}

	fmt.Println("selesai perulangan")

	names := []string{"Deniz", "Denis", "Deni", "Denny"}
	for i := 0; i < len(names); i++ {
		fmt.Println((names[i]))
	}

	for index, value := range names {
		fmt.Println("Index", index, "=>", value)
	}

	for _, value := range names {
		fmt.Println(value)
	}
}