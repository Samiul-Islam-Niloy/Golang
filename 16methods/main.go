package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//no inheritance in golang; No super or parent

	niloy := User{"Niloy", "niloy@go.dev", true, 17}
	fmt.Println(niloy)
	fmt.Printf("niloy details are: %+v\n", niloy)
	fmt.Printf("Name is %v and email is %v\n", niloy.Name, niloy.Email)
	niloy.GetStatus()
	niloy.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", niloy.Name, niloy.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
