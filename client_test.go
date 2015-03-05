package nationbuilder

import (
	"net/http"
	"testing"
)

func TestNationbuilderClientSetClient(t *testing.T) {
	n, err := NewNationbuilderClient(siteSlug, apiKey)
	if err != nil {
		t.Fatal(err.Error())
	}

	if n.c != http.DefaultClient {
		t.Error("Expected client to have default http client")
	}

	client := &http.Client{}
	n.SetClient(client)

	if n.c != client {
		t.Error("Expected client to have new http client")
	}

}
