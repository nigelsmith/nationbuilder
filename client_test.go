package nationbuilder

import (
	"net/http"
	"testing"
)

func TestClientSetHTTPClient(t *testing.T) {
	n, err := NewClient(siteSlug, apiKey)
	if err != nil {
		t.Fatal(err.Error())
	}

	if n.c != http.DefaultClient {
		t.Error("Expected client to have default http client")
	}

	client := &http.Client{}
	n.SetHTTPClient(client)

	if n.c != client {
		t.Error("Expected client to have new http client")
	}

}
