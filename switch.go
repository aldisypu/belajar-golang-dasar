package main

import "fmt"

func main() {
	name := "Aldi"

	switch name {
	case "Aldi":
		fmt.Println("Hello Aldi")
	case "Budi":
		fmt.Println("Hello Budi")
	default:
		fmt.Println("Hi, Boleh Kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	name = "PutraPutraPutra"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
