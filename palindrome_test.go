package palindrome

import (
	"testing"
)

// Test negative scenarios
func TestNegatives(t *testing.T) {
	str := "Gladys"
	actual, err := IsPalindrome(str)

	if actual {
		t.Fatalf(`String '%s', is not palindrome`, str)
	}

	if err != nil {
		t.Fatalf(`Expected no error, but got %s`, err)
	}

	str = "SomethingNot"
	actual, err = IsPalindrome(str)

	if actual {
		t.Fatalf(`String '%s', is not palindrome`, str)
	}

	if err != nil {
		t.Fatalf(`Expected no error, but got %s`, err)
	}
}

// Test positive scenarios
func TestPositives(t *testing.T) {
	str := "NooN"
	actual, err := IsPalindrome(str)

	if !actual {
		t.Fatalf(`String '%s', is palindrome`, str)
	}

	if err != nil {
		t.Fatalf(`Expected no error, but got %s`, err)
	}

	str = "RacecaR"
	actual, err = IsPalindrome(str)

	if !actual {
		t.Fatalf(`String '%s', is palindrome`, str)
	}

	if err != nil {
		t.Fatalf(`Expected no error, but got %s`, err)
	}

	str = " KadaK   "
	actual, err = IsPalindrome(str)

	if !actual {
		t.Fatalf(`String '%s', is palindrome`, str)
	}

	if err != nil {
		t.Fatalf(`Expected no error, but got %s`, err)
	}
}

// Test invalid scenarios
func TestInvalid(t *testing.T) {
	actual, err := IsPalindrome("")

	if actual {
		t.Fatalf(`Empty string is invalid`)
	}

	if err == nil {
		t.Fatalf(`Expected error, but got no error`)
	}

	actual, err = IsPalindrome("      ")

	if actual {
		t.Fatalf(`Empty string, is invalid`)
	}

	if err == nil {
		t.Fatalf(`Expected error, but got no error`)
	}
}
