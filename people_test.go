package nationbuilder

import (
	"testing"
)

func TestPersonString(t *testing.T) {
	p := &Person{
		ID:        testID,
		FirstName: "Phileas",
		LastName:  "Fogg",
	}

	pString, expected := p.String(), "Person ID: 1, Name: Phileas Fogg"
	if pString != expected {
		t.Errorf("Expected person string to be %s but saw %s", expected, pString)
	}
}
