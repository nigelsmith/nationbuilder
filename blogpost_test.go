package nationbuilder

import (
	"testing"
)

func TestBlogPost_String(t *testing.T) {
	post := &BlogPost{
		Page: Page{
			ID:   1,
			Name: "Test Name",
		},
	}
	s, expected := post.String(), "ID: 1, Blog Post: Test Name"
	if s != expected {
		t.Errorf("Expected %s but saw %s", expected, s)
	}
}
