package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int

	agePointer = &age

	fmt.Println("Age: ", *agePointer)

	editToAdultYears(agePointer)
	fmt.Println("Adult Years: ", age)
}

func editToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
