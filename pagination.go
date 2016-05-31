package nationbuilder

import (
	"net/url"
	"strconv"
)

// Paginated nationbuilder resources provide a previous and next url which
// can be called on subsequent requests.
// The convenience methods Next and Prev return a suitably configured Options object
// to be used with any further API calls.  A call which returns a nil reference
// indicates that there are no more prior or subsequent pages or that an error occurred.
// Errors should only occur if the paths returned by the nationbuilder API fail to parse.
type Pagination struct {
	PrevURL string `json:"prev"`
	NextURL string `json:"next"`
}

// Return an options object that will ask for the next page of results
func (p *Pagination) Next() (*Options, error) {
	return getOptionsforURL(p.NextURL)
}

// Return an options object that will ask for the previous page of results
func (p *Pagination) Prev() (*Options, error) {
	return getOptionsforURL(p.PrevURL)
}

func getOptionsforURL(pageURL string) (*Options, error) {
	if pageURL == "" {
		return nil, nil
	}

	u, err := url.Parse(pageURL)
	if err != nil {
		return nil, err
	}
	q := u.Query()

	var limit int
	l := q.Get("limit")
	if l != "" {
		limit, err = strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
	}

	return &Options{
		Limit:     limit,
		PageToken: q.Get("__token"),
		PageNonce: q.Get("__nonce"),
		queryOpts: q,
	}, nil
}
