package nationbuilder

import (
	"testing"
)

func TestBasicPage_String(t *testing.T) {
	p := &BasicPage{
		Page: Page{
			Name: "Test Page",
		},
	}

	s, expected := p.String(), "Basic Page: Test Page"
	if s != expected {
		t.Errorf("Expected %s but saw %s", expected, s)
	}

}
