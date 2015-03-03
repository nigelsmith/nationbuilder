package nationbuilder

import (
	"testing"
)

const apiKey = "testkey"
const slug = "testslug"

func TestNationbuilderURLExtendPath(t *testing.T) {
	n, err := NewNationbuilderClient(slug, apiKey)
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedPath := "/api/v1"
	url := n.baseURL

	if url.u.Path != expectedPath {
		t.Errorf("Expected path %s but saw %s", expectedPath, url.u.Path)
	}

	url.extendPath("/foo")
	expectedPath = "/api/v1/foo"

	if url.u.Path != expectedPath {
		t.Errorf("Expected extended path to be %s but saw %s", expectedPath, url.u.Path)
	}

	url.extendPath("bar")
	expectedPath = "/api/v1/foo/bar"
	if url.u.Path != expectedPath {
		t.Errorf("Expected extended path to be %s but saw %s", expectedPath, url.u.Path)
	}
}

func TestNewNationbuilderClient(t *testing.T) {
	c, err := NewNationbuilderClient(slug, apiKey)
	if err != nil {
		t.Fatal(err.Error())
	}

	if c.Slug != slug {
		t.Errorf("expected client to have slug %s but saw %s", slug, c.Slug)
	}

	if c.ApiKey != apiKey {
		t.Errorf("expected client to have api key %s but saw %s", apiKey, c.ApiKey)
	}

}
