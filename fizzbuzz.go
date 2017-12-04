package fizzbuzz

import (
	"fmt"
)

func FizzBuzz(times int) string {
	results := "FizzBuzz:\n"

	for fizzBuzzCount := 1; fizzBuzzCount <= times; fizzBuzzCount++ {
		result := ""
		if fizzBuzzCount%3 == 0 {
			result += "Fizz"
		}
		if fizzBuzzCount%5 == 0 {
			result += "Buzz"
		}
		if result != "" {
			results += (result + "\n")
			continue
		}
		results += fmt.Sprintf("%d\n", fizzBuzzCount)
	}

	return results
}
