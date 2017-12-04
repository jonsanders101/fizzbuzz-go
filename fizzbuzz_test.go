package fizzbuzz

import (
	"testing"
)

func TestPrintFizzBuzz(t *testing.T) {
	resultString := FizzBuzz(15)
	expectedString := "FizzBuzz:\n1\n2\nFizz\n4\nBuzz\nFizz\n7\n8\nFizz\nBuzz\n11\nFizz\n13\n14\nFizzBuzz\n"

	if resultString != expectedString {
		t.Error("Expecting:", resultString, "got:", expectedString)
	}
}
