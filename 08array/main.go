package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in Golang")

	var fruits [3]string
	fruits[0] = "Apple Big"
	//fruits[1] = "Banana"
	fruits[2] = "Orange"

	fmt.Println("Fruits - ", fruits)
	fmt.Println("Fruits array length- ", len(fruits))

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegetables - ", vegList)
	fmt.Println("Vegetables array length- ", len(vegList))
}
