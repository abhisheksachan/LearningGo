package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now().Nanosecond()
	fmt.Println(presentTime)

	//fmt.Println(presentTime.Format("01-02-2006 15:04:05"))

	createdDate := time.Date(1994, time.August, 4, 10, 10, 10, 0, time.UTC)
	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
