package lib

import "fmt"

func AssertEqual[T comparable](expected, actual T) {
	if expected != actual {
		fmt.Println("Test failed \n\texpected: ", expected, " got: ", actual)
	} else {
		fmt.Println("Test passed")
	}
}
