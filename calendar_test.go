package nationbuilder

import (
	"fmt"
	"testing"
)

func TestCalendar_String(t *testing.T) {
	c := &Calendar{
		Page: Page{
			Name: testName,
			ID:   testID,
		},
	}

	s, expected := c.String(), fmt.Sprintf("Calendar ID %d - %s", testID, testName)
	if s != expected {
		t.Errorf("Expected calendar string to be %s but saw %s", expected, s)
	}
}
