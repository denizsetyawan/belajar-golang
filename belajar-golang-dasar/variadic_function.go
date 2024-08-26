package main

import (
	"fmt"
)

func total(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	fmt.Println(total(1, 2, 3, 4, 5))
	fmt.Println(total(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(total(numbers...))
}