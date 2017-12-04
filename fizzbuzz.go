package fizzbuzz

import (
	"fmt"
)

func PrintFizzBuzz(times int) string {
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

func Fizzbuzz(number int) (string, int) {
	if number%15 == 0 {
		return "FizzBuzz", number
	}
	if number%3 == 0 {
		return "Fizz", number
	}
	if number%5 == 0 {
		return "Buzz", number
	}
	return "Regular ol' number", number
}
