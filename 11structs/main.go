package main

import "fmt"

func main() {
	fmt.Println("Structs in golang ")
	// no inheritance in golang

	user := User{"Abhishek", "abhishek@gmail.com", true, 100}
	fmt.Println(user)
	fmt.Printf("Abhishek details are: %+v\n", user)
	fmt.Printf("Name is: %v and email is: %v ", user.Name, user.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
