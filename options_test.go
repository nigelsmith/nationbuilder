package nationbuilder

import (
	"net/url"
	"testing"
)

const testQueryParam = "testQueryParam"
const testQueryValue = "testQueryValue"
const testURL = "https://testslug.nationbuilder.com/api/v1/foo"

func TestSetQueryOptions(t *testing.T) {
	o := NewOptions()
	o.SetQueryOption(testQueryParam, testQueryValue)
	v := o.queryOpts.Get(testQueryParam)
	if v != testQueryValue {
		t.Errorf("Expected query param %s to be set with value %s but saw %s instead", testQueryParam, testQueryValue, v)
	}

	// contains nil queryOpts - should be initialised if nil
	o = &Options{}
	o.SetQueryOption(testQueryParam, testQueryValue)
	v = o.queryOpts.Get(testQueryParam)
	if v != testQueryValue {
		t.Errorf("Expected query param %s to be set with value %s but saw %s instead", testQueryParam, testQueryValue, v)
	}
}

func TestSetQuery(t *testing.T) {
	u, err := url.Parse(testURL)
	if err != nil {
		t.Fatal(err.Error())
	}

	o := NewOptions()
	o.SetQueryOption(testQueryParam, testQueryValue)
	o.setQuery(u)

	expectedURL := testURL + "?" + testQueryParam + "=" + testQueryValue
	if u.String() != expectedURL {
		t.Errorf("Expected SetQuery to make url %s but saw %s", expectedURL, u.String())
	}

	testToken := "testToken"
	testNonce := "testNonce"

	o = &Options{
		Limit:     100,
		PageToken: testToken,
		PageNonce: testNonce,
	}

	u, err = url.Parse(testURL)
	if err != nil {
		t.Fatal(err.Error())
	}

	o.setQuery(u)
	expectedURL = testURL + "?__nonce=" + testNonce + "&__token=" + testToken + "&limit=100"
	if u.String() != expectedURL {
		t.Errorf("Expected SetQuery to make url %s but saw %s", expectedURL, u.String())
	}

}
