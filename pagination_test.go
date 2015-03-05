package nationbuilder

import (
	"fmt"
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

	o, err = getOptionsforURL("")
	if o != nil && err != nil {
		t.Error("Expected empty URL string to return nil options and error")
	}

	o, err = getOptionsforURL(":")
	if err == nil {
		t.Error("Expected error for non-sensical URL")
	}

	o, err = getOptionsforURL("/foo?limit=flibble")
	if err == nil {
		t.Error("Expected error for non integer limit")
	}
}

func TestNextAndPrev(t *testing.T) {
	tLimit, tNonce, tToken := 10, "nonce", "token"
	u := fmt.Sprintf("/foo?limit=%d&__nonce=%s&__token=%s", tLimit, tNonce, tToken)
	p := &Pagination{
		PrevURL: u,
		NextURL: u,
	}

	o, err := p.Next()
	if err != nil {
		t.Errorf("expected nil error but saw %s", err.Error())
		t.SkipNow()
	}

	if o.Limit != tLimit {
		t.Errorf("expected limit of %d but saw %d", tLimit, o.Limit)
	}

	if o.PageNonce != tNonce {
		t.Errorf("expected nonce of %s but saw %s", tNonce, o.PageNonce)
	}

	if o.PageToken != tToken {
		t.Errorf("expected token of %s but saw %s", tToken, o.PageToken)
	}

	o, err = p.Prev()
	if err != nil {
		t.Errorf("expected nil error but saw %s", err.Error())
		t.SkipNow()
	}

	if o.Limit != tLimit {
		t.Errorf("expected limit of %d but saw %d", tLimit, o.Limit)
	}

	if o.PageNonce != tNonce {
		t.Errorf("expected nonce of %s but saw %s", tNonce, o.PageNonce)
	}

	if o.PageToken != tToken {
		t.Errorf("expected token of %s but saw %s", tToken, o.PageToken)
	}
}
