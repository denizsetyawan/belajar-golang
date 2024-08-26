package main

import "fmt"

func main() {
	names := []string{"Deniz", "Denis", "Deni", "Denny"}

	slice1 := names[0:2]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	days := [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "Changed Saturday"
	daySlice1[1] = "Changed Sunday"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "New Day")
	daySlice2[0] = "Changed New Day"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Deniz"
	newSlice[1] = "Denis"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Deni")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Changed Deniz"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}