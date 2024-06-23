package main

import "fmt"

const Login string = "this is const"

func main() {
	var username string = "Abhishek"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var smallVal float32 = 2550.
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// dummy := 100
	fmt.Println(Login)
	fmt.Printf("Variable is of type: %T \n", Login)
}
