package main

import "fmt"

// seperti class. Define data dengan valuenya
// anggap saja ini tipe data baru
type Customer struct {
	Name, Address string
	Age           int
}

// memasukan struct ke dalam sebuah function
// dimana seolah-olah, si structnya ini punya sebuah function
func (customer Customer) sayHello(name string) {
	fmt.Println("HELLO ", name, "MY NAME IS ", customer.Name)
}

func main() {
	// first way to get type data struct
	var customer Customer
	customer.Name = "DJALAL KURNIA"
	customer.Address = "LEBAK"
	customer.Age = 20
	fmt.Println(customer)

	customer.sayHello("BUDI")

	//// second way
	//customer2 := Customer{
	//	Name:    "DJALAL 2",
	//	Address: "LEBAK 2",
	//	Age:     22,
	//}
	//fmt.Println(customer2)
	//
	//// thrid way
	//customer3 := Customer{"DJALAL 3", "LEBAK 3", 23}
	//fmt.Println(customer3)
}
