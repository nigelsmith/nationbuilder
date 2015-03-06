package nationbuilder

import (
	"testing"
)

func TestAttachmentString(t *testing.T) {
	a := &Attachment{
		FileName: "test",
	}

	s, expected := a.String(), "Attachment: test"
	if s != expected {
		t.Errorf("Expected string method output of %s but saw %s", expected, s)
	}
}
