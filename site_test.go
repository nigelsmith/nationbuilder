package nationbuilder

import (
	"testing"
)

func TestSiteString(t *testing.T) {
	site := &Site{
		Name: testName,
	}

	s, expected := site.String(), "Site: "+testName
	if s != expected {
		t.Errorf("Expected site string of %s but saw %s", expected, s)
	}
}
