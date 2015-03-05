package nationbuilder

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"testing"
)

func TestResultError(t *testing.T) {
	e := errors.New("Test Error")
	r := &Result{
		Err:        e,
		StatusCode: 400,
	}

	errS := r.Error()
	if errS != "Test Error" {
		t.Errorf("Expected error string to be Test Error but saw %s", errS)
	}

	r.Err = nil
	errS = r.Error()
	if errS != "" {
		t.Errorf("Expected empty error string for nil error but saw %s", errS)
	}

	r.NationErr = &NationError{
		Code:         "Error Code",
		Message:      "Test Message",
		ErrorMessage: "Test Error Message",
	}

	errS, expected := r.Error(), "StatusCode: 400, Code: Error Code, Message: Test Message, ErrorMsg: Test Error Message, ValidationErrors: "
	if errS != expected {
		t.Errorf("Expected NationError string of %s but saw %s", expected, errS)
	}

	r.NationErr.ValidationError = []string{"1", "2", "3", "4"}
	expected += "1 - 2 - 3 - 4"
	errS = r.Error()
	if errS != expected {
		t.Errorf("Expected NationError string with validation errors of %s but saw %s", expected, errS)
	}
}

func TestResultProcessResponse(t *testing.T) {
	r := &Result{
		StatusCode: 400,
	}
	nErr := &NationError{
		Code:         "Test Code",
		Message:      "Test Message",
		ErrorMessage: "Error Message",
	}

	b, err := json.Marshal(nErr)
	if err != nil {
		log.Fatal(err.Error())
	}
	r.Body = b

	r.processResponse(http.StatusOK, nil)
	if r.NationErr == nil {
		t.Error("Expected non-nil NationErr")
		t.SkipNow()
	}

	if r.NationErr.Code != nErr.Code ||
		r.NationErr.Message != nErr.Message ||
		r.NationErr.ErrorMessage != nErr.ErrorMessage {

		t.Error("Expected nation error to match supplied body")
	}

	blog := &Blog{
		Page: Page{
			Name: "Test Blog",
		},
	}
	b, err = json.Marshal(blog)
	if err != nil {
		log.Fatal(err.Error())
	}

	r = &Result{
		StatusCode: 200,
		Body:       b,
	}
	dstBlog := &Blog{}
	r.processResponse(http.StatusOK, dstBlog)

	if r.Err != nil {
		t.Error("Expected nil error")
		t.SkipNow()
	}

	if r.NationErr != nil {
		t.Error("Expected nil nation error")
		t.SkipNow()
	}

	if dstBlog.Name != "Test Blog" {
		t.Errorf("Expected blog to be set with name %s but saw %s", "Test Blog", dstBlog.Name)
	}

}
