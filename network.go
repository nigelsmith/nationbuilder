package nationbuilder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	// Current rate limit for NB is 10 reqs / second
	rateLimit = 10
)

// Result provides details of an API request
// Implements type Error
type Result struct {
	StatusCode int
	Err        error
	NationErr  *NationError
	Body       []byte
	Etag       string
}

func (r *Result) Error() string {
	if r.NationErr != nil {
		var validationErrors string
		if len(r.NationErr.ValidationError) > 0 {
			validationErrors = strings.Join(r.NationErr.ValidationError, " - ")
		}
		return fmt.Sprintf("StatusCode: %d, Code: %s, Message: %s, ErrorMsg: %s, ValidationErrors: %s",
			r.StatusCode, r.NationErr.Code, r.NationErr.Message, r.NationErr.ErrorMessage, validationErrors)
	}
	return r.Err.Error()
}

func (r *Result) processResponse(expectedStatus int, dst interface{}) {
	if r.StatusCode != expectedStatus {
		err := json.Unmarshal(r.Body, &r.NationErr)
		if err != nil {
			r.Err = err
		}

		return
	}

	if dst != nil {
		err := json.Unmarshal(r.Body, dst)
		if err != nil {
			r.Err = err

			return
		}
	}
}

func (r *Result) HasError() bool {
	return r.Err != nil || r.NationErr != nil
}

type NationError struct {
	Code            string   `json:"code"`
	Message         string   `json:"message"`
	ErrorMessage    string   `json:"error"`
	ValidationError []string `json:"validation_errors"`
}

type apiRequest struct {
	url    string
	method string
	etag   string
	body   []byte
}

func (n *NationbuilderClient) retrieve(r *apiRequest, dst interface{}) *Result {
	res := n.sendRequest(r)
	if res.Err != nil {
		return res
	}

	res.processResponse(http.StatusOK, dst)

	return res
}

func (n *NationbuilderClient) create(data interface{}, r *apiRequest, dst interface{}, expectedStatus int) *Result {
	b, err := json.Marshal(data)
	if err != nil {
		return &Result{
			Err: err,
		}
	}

	r.body = b

	res := n.sendRequest(r)
	if res.Err != nil {
		return res
	}

	// NB sometimes returns 200 on object creation and sometimes 201...
	res.processResponse(expectedStatus, dst)

	return res
}

func (n *NationbuilderClient) delete(r *apiRequest) *Result {
	res := n.sendRequest(r)
	if res.Err != nil {
		return res
	}

	res.processResponse(http.StatusNoContent, nil)

	return res
}

func (n *NationbuilderClient) sendRequest(request *apiRequest) *Result {
	req, err := http.NewRequest(request.method, request.url, bytes.NewReader(request.body))
	if err != nil {
		return &Result{
			Err: err,
		}
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if request.etag != "" {
		if debug {
			log.Println("Set etag of: " + request.etag)
		}
		req.Header.Set("If-None-Match", request.etag)
	}

	if debug {
		log.Printf("Making %s request to %s", request.method, request.url)
	}

	resp, err := n.c.Do(req)
	if err != nil {
		return &Result{
			Err: err,
		}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Result{
			StatusCode: resp.StatusCode,
			Err:        err,
		}
	}
	resp.Body.Close()

	if debug {
		log.Printf("StatusCode %d for %s to %s", resp.StatusCode, request.method, request.url)
	}

	return &Result{
		StatusCode: resp.StatusCode,
		Body:       body,
		Etag:       resp.Header.Get("Etag"),
	}
}

// func sendApiCall(apiReq *apiRequest, ticker *time.Ticker, wg *sync.WaitGroup, results chan<- *Result) {
// 	defer wg.Done()

// 	req, err := http.NewRequest(apiReq.method, apiReq.url, bytes.NewReader(apiReq.body))
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Accept", "application/json")

// 	if apiReq.etag != "" {
// 		if debug {
// 			log.Println("Set etag of: " + apiReq.etag)
// 		}
// 		req.Header.Set("If-None-Match", apiReq.etag)
// 	}

// 	<-ticker.C

// 	resp, err := c.Do(req)
// 	if err != nil {
// 		results <- &Result{
// 			Err: err,
// 		}
// 		return
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		results <- &Result{
// 			Err:        err,
// 			StatusCode: resp.StatusCode,
// 		}
// 		return
// 	}
// 	resp.Body.Close()

// 	results <- &Result{
// 		StatusCode: resp.StatusCode,
// 		Etag:       resp.Header.Get("Etag"),
// 		Body:       body,
// 	}
// }

// func processRequests(requests ...*apiRequest) []*Result {
// 	ticker := time.NewTicker(1e9 / rateLimit)
// 	resultChan := make(chan *Result)
// 	results := make([]*Result, 0)
// 	wg := new(sync.WaitGroup)

// 	defer ticker.Stop()
// 	wg.Add(len(requests))

// 	go func() {
// 		wg.Wait()
// 		close(resultChan)
// 	}()

// 	for _, req := range requests {
// 		go sendApiCall(req, ticker, wg, resultChan)
// 	}

// 	for r := range resultChan {
// 		results = append(results, r)
// 	}

// 	return results

// }
