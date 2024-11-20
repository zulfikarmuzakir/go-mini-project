package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	squareRoot := int(math.Sqrt(float64(n)))

	for i := 2; i <= squareRoot; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var numbers []int = make([]int, 100)

	for i := 0; i < 100; i++ {
		numbers[i] = i + 1
	}

	for i := 99; i >= 0; i-- {
		num := numbers[i]

		if isPrime(num) {
			continue
		}

		if num%3 == 0 && num%5 == 0 {
			fmt.Printf("FooBar ")
		} else if num%3 == 0 {
			fmt.Printf("Foo ")
		} else if num%5 == 0 {
			fmt.Printf("Bar ")
		} else {
			fmt.Printf("%d ", num)
		}
	}

	// fmt.Println(numbers)
}
