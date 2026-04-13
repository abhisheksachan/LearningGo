package main

import "fmt"

// 1. Simple defer - executes at end of function
func simpleDeferExample() {
	fmt.Println("Start")
	
	defer fmt.Println("This runs last (deferred)")
	
	fmt.Println("Middle")
	fmt.Println("End")
	// Output: Start, Middle, End, This runs last (deferred)
}

// 2. Multiple defers - LIFO (Last In First Out) order
func multipleDeferExample() {
	defer fmt.Println("Defer 1 - First to defer")
	defer fmt.Println("Defer 2 - Second to defer")
	defer fmt.Println("Defer 3 - Third to defer (executes first)")
	
	fmt.Println("Main function body")
	// Output: Main function body, Defer 3, Defer 2, Defer 1
}

// 3. Defer with file operations (cleanup pattern)
func fileOperationExample() {
	filename := "test.txt"
	fmt.Printf("Opening file: %s\n", filename)
	
	// Simulate file handling with defer cleanup
	defer fmt.Printf("Closing file: %s\n", filename)
	
	fmt.Println("Reading/writing file...")
	// Actual file operations would go here
}

// 4. Panic - terminates program execution
func panicExample() {
	fmt.Println("Before panic")
	
	panic("Something went wrong! This will stop execution.")
	
	fmt.Println("This won't execute") // Never reached
}

// 5. Recover - catches panic and continues execution
func recoverExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()
	
	fmt.Println("About to panic...")
	panic("Oops! An error occurred")
	fmt.Println("This won't print") // Won't execute
	
	// After recover, execution continues in calling function
}

// 6. Panic with defer
func panicWithDeferExample() {
	defer fmt.Println("This defer runs even after panic!")
	
	fmt.Println("Before panic")
	panic("Critical error")
	
	fmt.Println("Won't reach here")
}

// 7. Recover in deferred function
func recoverInDeferExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered: %v\n", r)
		}
	}()
	
	defer fmt.Println("Defer 1 - runs before recover check")
	
	fmt.Println("Starting function")
	panic("Something failed!")
}

// 8. Defer with value capture
func deferWithValueCapture() {
	x := 10
	
	// Deferred function captures the VALUE of x at defer time
	defer fmt.Printf("Deferred: x = %d\n", x)
	
	x = 20
	fmt.Printf("Current: x = %d\n", x)
	// Output: Current: x = 20, Deferred: x = 20 (value captured at defer, not execution)
}

// 9. Defer with function call - captures reference
func deferWithFunctionCall() {
	x := 10
	
	// Defer wraps a function that captures x by reference
	defer func() {
		fmt.Printf("Deferred function: x = %d\n", x)
	}()
	
	x = 30
	fmt.Printf("Current: x = %d\n", x)
	// Output: Current: x = 30, Deferred function: x = 30
}

// 10. Custom error handling with defer and recover
func divideWithRecovery(a, b int) (result int, err string) {
	defer func() {
		if r := recover(); r != nil {
			result = 0
			err = fmt.Sprintf("Error: %v", r)
		}
	}()
	
	if b == 0 {
		panic("Division by zero!")
	}
	
	result = a / b
	return
}

// 11. Defer for resource cleanup (database connection pattern)
func databaseConnectionExample() {
	fmt.Println("Connecting to database...")
	
	defer func() {
		fmt.Println("Closing database connection")
	}()
	
	fmt.Println("Executing queries...")
	// Database operations would go here
	
	fmt.Println("Query completed")
	// Connection closes here via defer
}

func main() {
	fmt.Println("=== Defer, Panic, Recover in Go ===\n")

	// 1. Simple defer
	fmt.Println("--- 1. Simple Defer ---")
	simpleDeferExample()
	
	fmt.Println("\n--- 2. Multiple Defers (LIFO) ---")
	multipleDeferExample()
	
	fmt.Println("\n--- 3. Defer with Cleanup ---")
	fileOperationExample()
	
	fmt.Println("\n--- 4. Panic (caught below) ---")
	// Wrap in function to prevent program crash
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered: %v\n", r)
			}
		}()
		panicExample()
	}()
	
	fmt.Println("\n--- 5. Recover Example ---")
	recoverExample()
	
	fmt.Println("\n--- 6. Panic with Defer ---")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered: %v\n", r)
			}
		}()
		panicWithDeferExample()
	}()
	
	fmt.Println("\n--- 7. Recover in Defer ---")
	recoverInDeferExample()
	
	fmt.Println("\n--- 8. Defer with Value Capture ---")
	deferWithValueCapture()
	
	fmt.Println("\n--- 9. Defer with Reference Capture ---")
	deferWithFunctionCall()
	
	fmt.Println("\n--- 10. Division with Recovery ---")
	result, err := divideWithRecovery(10, 2)
	fmt.Printf("10 / 2 = %d (error: %s)\n", result, err)
	
	result, err = divideWithRecovery(10, 0)
	fmt.Printf("10 / 0 = %d (error: %s)\n", result, err)
	
	fmt.Println("\n--- 11. Database Connection Example ---")
	databaseConnectionExample()
}
