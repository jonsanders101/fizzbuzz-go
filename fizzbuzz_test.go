package fizzbuzz

import (
	"testing"
)

func TestReturnsFizzFor3(t *testing.T) {
	resultString, resultNum := Fizzbuzz(3)

	if resultString != "Fizz" {
		t.Log("Expected 'Fizz', got: ", resultString)
		t.Fail()
	}
	if resultNum != 3 {
		t.Log("Expected 3, got: ", resultNum)
		t.Fail()
	}
}

func TestReturnsRegularNumber(t *testing.T) {
	resultString, resultNum := Fizzbuzz(1)

	if resultString != "Regular ol' number" {
		t.Log("Expected 'Regular ol' number', got: ", resultString)
		t.Fail()
	}
	if resultNum != 1 {
		t.Log("Expected 1, got: ", resultNum)
		t.Fail()
	}
}

func TestReturnsBuzzFor5(t *testing.T) {
	resultString, resultNum := Fizzbuzz(5)

	if resultString != "Buzz" {
		t.Log("Expected 'Buzz', got: ", resultString)
		t.Fail()
	}
	if resultNum != 5 {
		t.Log("Expected 5, got: ", resultNum)
		t.Fail()
	}
}

func TestReturnsFizzBuzzFor15(t *testing.T) {
	resultString, resultNum := Fizzbuzz(15)

	if resultString != "FizzBuzz" {
		t.Log("Expected 'FizzBuzz', got: ", resultString)
		t.Fail()
	}
	if resultNum != 15 {
		t.Log("Expected 15, got: ", resultNum)
		t.Fail()
	}
}