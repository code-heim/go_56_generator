package main

import (
	"fmt"
	"math"
)

// Function that returns the next prime number after a given integer
func nextPrime(n int) int {
	if n < 2 {
		return 2
	}
	for {
		n++
		if isPrime(n) {
			return n
		}
	}
}

// Helper function to check if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Prime generator function
func primeGenerator(done <-chan bool, operation func(int) int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		num := 1 // Initial value to find the first prime
		for {
			num = operation(num)
			select {
			case <-done:
				fmt.Println("Received done signal, stopping generator.")
				return
			case ch <- num:
			}
		}
	}()
	return ch
}

func main() {
	done := make(chan bool)

	i := 0
	for prime := range primeGenerator(done, nextPrime) {
		i++
		fmt.Println("Prime #", i, ": ", prime)
		if i == 10 {
			break
		}
	}

	// Signal the generator to stop
	done <- true
}
