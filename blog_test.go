package nationbuilder

import (
	"testing"
)

func TestBlog_String(t *testing.T) {
	blog := &Blog{
		Page: Page{
			Name: testName,
		},
	}

	s, expected := blog.String(), "Blog: testName"
	if s != expected {
		t.Errorf("Expected blog string to be %s but saw %s", expected, s)
	}
}
