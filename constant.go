package main

import "fmt"

func main() {
	const (
		firstName string = "Aldi"
		lastName         = "Putra"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// error
	// firstName = "Budi"
	// lastName = "Joko"
}
