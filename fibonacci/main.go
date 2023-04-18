package main

import "fmt"

func main() {
	const n = 5
	fmt.Println("Calculating %d sequence\n", n)
}

func fibonacci(n int, old, new, calc string) int {

	if n == 0 {
		return
	}

	fibonacci(n-1, old, calc, new)

	fmt.Printf("  n  ", old, new)

	fibonacci(n-1, calc, new, old)

}

// Fill me in!
// When n is 0 or 1, fibonacci(n) should just be n.
// When n is any larger number,
// fibonacci(n) should be the sum of the previous two fibonacci numbers.

//	hanoi(n-1, from, spare, to)
//	fmt.Println("  Move 1 disc from %s to %s\n", from, to)
//	hanoi(n-1, spare, to, from)
