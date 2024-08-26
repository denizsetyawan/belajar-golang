package main

import "fmt"

type validationError struct {
	message string
}

func (v *validationError) Error() string {
	return v.message
}

type notFoundError struct {
	message string
}

func (n *notFoundError) Error() string {
	return n.message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}

	if id != "deniz" {
		return &notFoundError{"not found"}
	}

	return nil
}

func main() {
	err := SaveData("deniz", nil)

	if err != nil {
		// terjadi eror
		// if validationError, ok := err.(*validationError); ok {
		// 	fmt.Println("validation :" ,validationError.Error())
		// } else if notFoundError, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found :",notFoundError.Error())
		// } else {
		// 	fmt.Println("unknown error")
		// }

		// switch
		switch finalError := err.(type) {
			case *validationError:
				fmt.Println("validation :", finalError.Error())
			case *notFoundError:
				fmt.Println("not found :", finalError.Error())
			default:
				fmt.Println("unknown error")
		}
	} else {
		fmt.Println("success")
	}
}