package Is_test

import (
	"testing"
	"is.go/src"
)

func TestURI(t  *testing.T){	
	message := Is.URL("http://www.bloket.pro") 

	if !message {
		t.Error("URI cannot be determine by regex")
	}

	message = Is.URL("noturi")

	if message{
		t.Error("Oo noo! this is URI")
	}
}

func TestEmail(t *testing.T){
	message := Is.Email("mail@bloket.pro")

	if !message {
		t.Error("Emali cannot be determine by regex")
	}

	message = Is.Email("itsnotmail")

	if message{
		t.Error("is this realy email?")
	}
}