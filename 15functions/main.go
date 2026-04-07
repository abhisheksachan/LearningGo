package main

import "fmt"

// 1. Basic function with multiple return values
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// 2. Named returns (cleaner for multiple returns)
func divideNamed(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b
	return // implicit return of named values
}

// 3. Variadic function (accepts variable number of arguments)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 4. Function returning a function (closure)
func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// Rectangle struct for method examples
type Rectangle struct {
	width, height float64
}

// Method with value receiver (doesn't modify the struct)
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Method with pointer receiver (can modify the struct)
func (r *Rectangle) Scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

func main() {
	fmt.Println("=== Functions in Go ===")

	// Multiple returns
	q, r := divide(17, 3)
	fmt.Printf("17 ÷ 3 = %d remainder %d\n", q, r)

	// Named returns
	q2, r2 := divideNamed(20, 7)
	fmt.Printf("20 ÷ 7 = %d remainder %d\n", q2, r2)

	// Variadic function
	fmt.Printf("Sum of 1,2,3,4,5 = %d\n", sum(1, 2, 3, 4, 5))
	fmt.Printf("Sum of 10,20 = %d\n", sum(10, 20))

	// Function returning function
	add5 := makeAdder(5)
	fmt.Printf("5 + 3 = %d\n", add5(3))
	fmt.Printf("5 + 10 = %d\n", add5(10))

	// Anonymous function
	func() {
		fmt.Println("This is an anonymous function!")
	}()

	// Methods
	fmt.Println("\n=== Methods ===")
	rect := Rectangle{width: 10, height: 5}
	fmt.Printf("Rectangle: %.1fx%.1f\n", rect.width, rect.height)
	fmt.Printf("Area: %.2f\n", rect.Area())

	rect.Scale(2)
	fmt.Printf("After scaling by 2: %.1fx%.1f\n", rect.width, rect.height)
	fmt.Printf("New area: %.2f\n", rect.Area())
}