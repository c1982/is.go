package Is_test

import "testing"
import "is.go/src"

func TestLowercase(t *testing.T)  {		
	
	message := Is.Lowercase("NOT Lower")

	if message {
		t.Error("It's lowercase text")
	}

	message = Is.Lowercase("yes lower")

	if !message {
		t.Error("Yes it's not lowercase text")
	}	
}

func TestUppercase(t *testing.T){
	message := Is.Uppercase("not upper")

	if message {
		t.Error("Yeah! upper case dude.")
	}

	message = Is.Uppercase("YES UPPER")

	if !message {
		t.Error("Oha amk.")
	}
}

func TestCapitalized(t *testing.T){
	message := Is.Capitalized("This Is Capitalized  Text")

	if !message {
			t.Error("It's not capitalized text")
	}

	message = Is.Capitalized("this is not capitalized text")

	if message {
		t.Error("Oh Noo! Capitalized found.")
	}
}

func TestPalindrome(t *testing.T){
	message := Is.Palindrome("makam")

	if !message{
		t.Error("It's not palindrome")
	}

	message = Is.Palindrome("MakaM")

	if !message{
		t.Error("It's not realy palindrome")
	}
}

func TestInclude(t *testing.T){
	message := Is.Include("Follow the white rabbit","white")

	if !message{
		t.Error("white not included")
	}

	message = Is.Include("Choose your destiny","white")

	if message {
		t.Error("white is included")
	}
}