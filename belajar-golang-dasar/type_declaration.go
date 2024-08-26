package main

import "fmt"

func main()  {
	type noKTP string
	var noKTP1 noKTP = "123123123123"

	var contoh string = "123123123124"
	var contoh2 = noKTP(contoh)

	fmt.Println(noKTP1)
	fmt.Println(contoh2)
}