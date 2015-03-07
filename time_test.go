package nationbuilder

import (
	"bytes"
	"encoding/json"
	"testing"
)

const testTime = "2015-03-04T12:45:28+00:00"

func TestDateString(t *testing.T) {
	n, err := NewDate(testTime)
	if err != nil {
		t.Fatal(err.Error())
	}

	s, expected := n.String(), testTime
	if s != expected {
		t.Errorf("Expected Date string to be %s but saw %s", expected, s)
	}
}

func TestDateMarshal(t *testing.T) {
	n, err := NewDate(testTime)
	if err != nil {
		t.Fatal(err.Error())
	}

	b, err := json.Marshal(n)
	if err != nil {
		t.Fatal(err.Error())
	}

	actual, expected := b, []byte(`"`+testTime+`"`)
	if !bytes.Equal(actual, expected) {
		t.Errorf("Expected marshalled Date to be %v but saw %v", expected, actual)
	}
}

func TestDateUnmarshal(t *testing.T) {
	j := []byte(`"` + testTime + `"`)
	nDate := &Date{}

	err := json.Unmarshal(j, nDate)
	if err != nil {
		t.Fatal(err.Error())
	}

	actual, expected := nDate.String(), testTime
	if actual != expected {
		t.Errorf("Expected unmarshalling Date to produce %s but saw %s", expected, actual)
	}
}
