package nationbuilder

import (
	"net/url"
	"strconv"
)

// Options describes available options for the API
// In particular the limit and pagination options for paginated resource endpoints
type Options struct {
	Limit     int
	PageToken string
	PageNonce string
	queryOpts url.Values
}

// Set arbitrary query options
func (o *Options) SetQueryOption(key string, value string) {
	if o.queryOpts == nil {
		o.queryOpts = make(url.Values)
	}
	o.queryOpts.Set(key, value)
}

func (o *Options) setQuery(u *url.URL) {
	q := u.Query()

	if o.queryOpts == nil {
		o.queryOpts = make(url.Values)
	}

	for k, v := range o.queryOpts {
		for _, s := range v {
			q.Add(k, s)
		}
	}

	if o.Limit != 0 {
		q.Set("limit", strconv.Itoa(o.Limit))
	}

	if o.PageToken != "" && o.PageNonce != "" {
		q.Set("__token", o.PageToken)
		q.Set("__nonce", o.PageNonce)
	}

	u.RawQuery = q.Encode()
}

// Instantiate an Options object
func NewOptions() *Options {
	return &Options{
		queryOpts: make(url.Values),
	}
}
