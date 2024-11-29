package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func generateInfinitePrimes() {
	n := 1
	for {
		if isPrime(n) {
			fmt.Printf("%d ", n)
		}
		n++
		// time.Sleep(100 * time.Millisecond) // Adding small delay for readability
	}
}

func generatePrimes(start, end int) []int {
	sieve := make([]bool, end+1)
	for i := 2; i <= end; i++ {
		sieve[i] = true
	}

	for i := 2; i*i <= end; i++ {
		if sieve[i] {
			for j := i * i; j <= end; j += i {
				sieve[j] = false
			}
		}
	}

	var primes []int
	for i := start; i <= end; i++ {
		if i >= 2 && sieve[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func printUsage() {
	fmt.Println("Prime Number Generator and Checker")
	fmt.Println("\nUsage:")
	fmt.Println("  Check if a number is prime:")
	fmt.Println("    --check=N     Check if N is a prime number")
	fmt.Println("\n  Generate prime numbers in range:")
	fmt.Println("    --start=N     Start of range (default: 1)")
	fmt.Println("    --end=N       End of range (default: 100)")
	fmt.Println("\n  No arguments: Will print help and then generate infinite prime numbers")
	fmt.Println("\nExamples:")
	fmt.Println("  Check if 17 is prime:")
	fmt.Println("    prime --check=17")
	fmt.Println("\n  Generate primes between 10 and 50:")
	fmt.Println("    prime --start=10 --end=50")
}

func main() {
	start := flag.Int("start", -1, "Start of range")
	end := flag.Int("end", -1, "End of range")
	check := flag.Int("check", -1, "Check if a number is prime")
	flag.Parse()

	if len(os.Args) == 1 {
		printUsage()
		fmt.Println("\nGenerating infinite prime numbers (Press Ctrl+C to stop):")
		generateInfinitePrimes()
		return
	}

	if *check != -1 {
		result := isPrime(*check)
		if result {
			fmt.Printf("%d is a prime number\n", *check)
		} else {
			fmt.Printf("%d is not a prime number\n", *check)
		}
		return
	}

	if *start != -1 && *end != -1 {
		if *start < 1 || *end < *start {
			fmt.Println("Invalid range. Start must be >= 1 and end must be >= start")
			return
		}

		primes := generatePrimes(*start, *end)
		fmt.Printf("Prime numbers between %d and %d are:\n", *start, *end)
		for _, prime := range primes {
			fmt.Printf("%d ", prime)
		}
		fmt.Println()
	}
}
