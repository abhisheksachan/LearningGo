package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang")

	days:= []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for d := 0; d < len(days); d++ {
		fmt.Printf("The day is %v\n", days[d])
	}

	for index, day := range days {
		fmt.Printf("The index is %v and the day is %v\n", index, day)
	}

	for index, day := range days {
		fmt.Printf("The day is %v at index %v\n", day, index)
	}

	rougevalue := 10
	for rougevalue < 20 {

		if rougevalue == 13 {
			goto mygoto
		}

		if rougevalue == 15 {
			rougevalue++
			continue
		}

		fmt.Printf("Rougue value is %v\n", rougevalue)
		rougevalue++
	}

	mygoto:
	fmt.Println("I am inside goto statement")

}
