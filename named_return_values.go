package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Aldi"
	middleName = "Syah"
	lastName = "Putra"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
