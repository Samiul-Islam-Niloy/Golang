package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//no inheritance in golang; No super or parent

	niloy := User{"Niloy", "niloy@go.dev", true, 17}
	fmt.Println(niloy)
	fmt.Printf("niloy details are: %+v\n", niloy)
	fmt.Printf("Name is %v and email is %v", niloy.Name, niloy.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
