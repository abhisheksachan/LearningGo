package main

import "fmt"

func main() {
	fmt.Println("Structs in golang ")
	// no inheritance in golang

	user := User{"Abhishek", "abhishek@gmail.com", true, 30}
}

type user struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
