package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var fruitList = []string{"Apple", "Orange", "Banana", "Grape"}
	fmt.Printf("Tyoe of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Peach")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 200
	highScores[1] = 450
	highScores[2] = 300
	highScores[3] = 600

	highScores = append(highScores, 50, 70, 90)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	// how to remove value from slice based on index
	var courses = []string{"java", "go", "javascript", "python", "swift", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
