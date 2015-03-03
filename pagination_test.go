package nationbuilder

import (
	"testing"
)

func TestGetOptionsForURL(t *testing.T) {
	url := "https://testslug.nationbuilder.com/api/v1/blogs?limit=20&foo=bar"
	o, err := getOptionsforURL(url)
	if err != nil {
		t.Fatal(err.Error())
	}

	if o.Limit != 20 {
		t.Errorf("Expected limit of %d but saw %d", 20, o.Limit)
	}

	if o.PageNonce != "" {
		t.Errorf("Expected empty page nonce but saw %s", o.PageNonce)
	}

	if o.PageToken != "" {
		t.Errorf("Expected empty page token but saw %s", o.PageToken)
	}

	v := o.queryOpts.Get("foo")
	if v != "bar" {
		t.Errorf("Expected query param foo to have value bar but saw %s", v)
	}
}
