package nationbuilder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	// Current rate limit for NB is 10 reqs / second
	rateLimit = 10
)

var (
	c = new(http.Client)
)

// Result provides details of an API request
// Implements type Error
type result struct {
	statusCode int
	err        error
	body       []byte
	url        string
}

func (r *result) Error() string {
	return r.err.Error()
}

type apiResults []*result

func (a *apiResults) getErrors() (errorCollection, bool) {
	e := make([]error)
	errFound := false

	for _, r := range a {
		if r.err != nil {
			e = append(e, r.err)
			errFound = true
		}
	}

	return errorCollection(e), errFound
}

type apiRequest struct {
	url    string
	body   []byte
	method string
}

// Addresser returns a string URL for a resource
type addresser interface {
	GetURL() string
}

// Creator describes a type capable of being used to
// create a Nationbuilder resource
type creator interface {
	addresser
	GetBody() ([]byte, error)
}

func create(creators ...creator) apiResults {

	reqs := make([]*apiRequest)
	results := make([]*result)

	for _, c := range creators {
		b, err := c.GetBody()
		if err != nil {
			results = append(results, &result{
				err: err,
			})
			continue
		}

		reqs = append(reqs, &apiRequest{
			url:    c.GetURL(),
			method: "POST",
			body:   b,
		})
	}

	r := processRequests(reqs)
	results = append(results, r)

	return apiResults(results)
}

func expectStatus(statusCode int, res *result) {
	// At this point nationbuilder is indicating an error
	if res.statusCode != statusCode {
		nbErr := &NationbuilderError{}
		err := json.Unmarshal(res.body, nbErr)
		if err != nil {
			res.err = err
		} else {
			res.err = nbErr
		}
	}
}

func retrieve(a addresser) *result {
	r := &apiRequest{
		url:    a.GetURL(),
		method: "GET",
	}

	res := processRequests(r)[0]
	if res.err != nil {
		return res
	}

	expectStatus(http.StatusOK, res)

	return res
}

func sendApiCall(apiReq *apiRequest, ticker *time.Ticker, wg *sync.WaitGroup, results chan<- *result) {
	defer wg.Done()

	req, err := http.NewRequest(apiReq.method, apiReq.url, bytes.NewReader(apiReq.body))
	if err != nil {
		log.Fatal(err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	<-ticker.C

	r := &result{
		url: apiReq.url,
	}

	resp, err := c.Do(req)
	if err != nil {
		r.err = err
		results <- r
		return
	}

	r.statusCode = resp.StatusCode

	r.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		r.err = err
		results <- r
		return
	}
	resp.Body.Close()

	results <- r
}

func processRequests(requests ...*apiRequest) []*result {
	ticker := time.NewTicker(1e9 / rateLimit)
	resultChan := make(chan *result)
	results := make([]*result, 0)
	wg := new(sync.WaitGroup)

	defer ticker.Stop()
	wg.Add(len(requests))

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for _, req := range requests {
		go sendApiCall(req, ticker, wg, resultChan)
	}

	for r := range resultChan {
		results = append(results, r)
	}

	return results

}
