package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := Pembagian(100, 0)

	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Hasil", err.Error())
	}
}