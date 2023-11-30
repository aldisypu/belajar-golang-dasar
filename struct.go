package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var aldi Customer
	fmt.Println(aldi)

	aldi.Name = "Aldi Syah"
	aldi.Address = "Indonesia"
	aldi.Age = 30
	fmt.Println(aldi)
	fmt.Println(aldi.Name)
	fmt.Println(aldi.Address)
	fmt.Println(aldi.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 30}
	fmt.Println(budi)

	budi.sayHello("Agus")
	aldi.sayHello("Agus")
	joko.sayHello("Agus")
}
