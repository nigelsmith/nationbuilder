package nationbuilder

import (
	"net/url"
	"strconv"
)

type Pagination struct {
	PrevURL string `json:"prev"`
	NextURL string `json:"next"`
}

func (p *Pagination) Next() (*Options, error) {
	return getOptionsforURL(p.NextURL)
}

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

	limit, err := strconv.Atoi(q.Get("limit"))
	if err != nil {
		return nil, err
	}

	return &Options{
		Limit:     limit,
		PageToken: q.Get("__token"),
		PageNonce: q.Get("__nonce"),
		queryOpts: q,
	}, nil
}
