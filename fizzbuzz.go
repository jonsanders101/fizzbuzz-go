package fizzbuzz

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
