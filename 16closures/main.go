package main

import "fmt"

// 1. Simple closure - capturing variables from outer scope
func simpleClosureExample() {
	message := "Hello from closure!"
	
	// Anonymous function that captures 'message'
	greet := func() {
		fmt.Println(message)
	}
	
	greet() // Prints: Hello from closure!
}

// 2. Closure that modifies captured variable
func counterExample() {
	count := 0
	
	// increment closes over 'count' and modifies it
	increment := func() int {
		count++
		return count
	}
	
	// Each call increments the same 'count' variable
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
	fmt.Println(increment()) // 3
}

// 3. Function factory - returns a closure with captured parameter
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// 4. Closure with state - ledger example
func makeLedger() func(string, int) int {
	balance := 0
	
	return func(operation string, amount int) int {
		switch operation {
		case "deposit":
			balance += amount
		case "withdraw":
			balance -= amount
		}
		return balance
	}
}

// 5. Closure in loops - common gotcha (in Go < 1.22)
// Note: Go 1.22+ changed loop variable semantics - each iteration gets its own copy
func closureInLoopsGotcha() {
	fmt.Println("=== Closure in Loops - Gotcha (Go < 1.22) ===")
	
	// In Go 1.21 and earlier, this would capture the same 'i' variable for all closures
	// In Go 1.22+, each iteration creates a new 'i', so this "just works"
	functions := make([]func(), 5)
	
	for i := 0; i < 5; i++ {
		functions[i] = func() {
			fmt.Print(i, " ")
		}
	}
	
	// In Go 1.22+: prints "0 1 2 3 4 " (each closure has its own i)
	// In Go < 1.22: would print "5 5 5 5 5 " (all closures share same i)
	fmt.Print("Go 1.22+ output: ")
	for _, f := range functions {
		f()
	}
	fmt.Println()
}

// 5b. Closure in loops - correct approach
func closureInLoopsCorrect() {
	fmt.Println("=== Closure in Loops - Correct ===")
	
	functions := make([]func(), 5)
	
	for i := 0; i < 5; i++ {
		// Capture by creating a new variable for each iteration
		i := i // Shadow i with a new variable
		functions[i] = func() {
			fmt.Print(i, " ")
		}
	}
	
	// Now each closure has its own copy
	fmt.Print("Correct output: ")
	for _, f := range functions {
		f()
	}
	fmt.Println()
}

// 6. Closure returning multiple values
func makeBankAccount(initialBalance float64) func(string, float64) (float64, string) {
	balance := initialBalance
	
	return func(operation string, amount float64) (float64, string) {
		var message string
		
		switch operation {
		case "deposit":
			balance += amount
			message = fmt.Sprintf("Deposited $%.2f, new balance: $%.2f", amount, balance)
		case "withdraw":
			if amount > balance {
				message = fmt.Sprintf("Insufficient funds. Balance: $%.2f", balance)
			} else {
				balance -= amount
				message = fmt.Sprintf("Withdrew $%.2f, new balance: $%.2f", amount, balance)
			}
		}
		
		return balance, message
	}
}

func main() {
	fmt.Println("=== Closures in Go ===\n")

	// 1. Simple closure
	fmt.Println("--- Simple Closure ---")
	simpleClosureExample()
	
	fmt.Println()

	// 2. Closure that modifies variable
	fmt.Println("--- Counter Closure ---")
	counterExample()
	
	fmt.Println()

	// 3. Function factory
	fmt.Println("--- Function Factory ---")
	double := makeMultiplier(2)
	triple := makeMultiplier(3)
	
	fmt.Printf("Double of 5: %d\n", double(5))  // 10
	fmt.Printf("Triple of 5: %d\n", triple(5))  // 15
	
	fmt.Println()

	// 4. Closure with state - ledger
	fmt.Println("--- Ledger Closure ---")
	ledger := makeLedger()
	
	fmt.Printf("Deposit 100: $%d\n", ledger("deposit", 100))
	fmt.Printf("Withdraw 30: $%d\n", ledger("withdraw", 30))
	fmt.Printf("Deposit 50: $%d\n", ledger("deposit", 50))
	fmt.Printf("Current balance: $%d\n", ledger("withdraw", 0))
	
	fmt.Println()

	// 5. Closure in loops - gotcha and correct
	closureInLoopsGotcha()
	closureInLoopsCorrect()
	
	fmt.Println()

	// 6. Bank account with closure
	fmt.Println("--- Bank Account Closure ---")
	account := makeBankAccount(1000)
	
	newBalance, msg := account("deposit", 500)
	fmt.Println(msg, "New balance:", newBalance)
	
	newBalance, msg = account("withdraw", 200)
	fmt.Println(msg, "New balance:", newBalance)
	
	newBalance, msg = account("withdraw", 2000)
	fmt.Println(msg, "New balance:", newBalance)
}