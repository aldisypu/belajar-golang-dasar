package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Aldi"
	names[1] = "Syah"
	names[2] = "Putra"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
		90,
		80,
		95,
		100,
		110,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
}
