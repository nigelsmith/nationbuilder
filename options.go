package nationbuilder

import (
	"net/url"
	"strconv"
)

type Options struct {
	Limit     int
	PageToken string
	PageNonce string
	Etag      string
	queryOpts url.Values
}

func (o *Options) SetQueryOption(key string, value string) {
	if o.queryOpts == nil {
		o.queryOpts = make(url.Values)
	}
	o.queryOpts.Set(key, value)
}

func (o *Options) setQuery(u *url.URL) {
	q := u.Query()

	for k, v := range o.queryOpts {
		for _, s := range v {
			q.Add(k, s)
		}
	}

	var limit int
	if o.Limit == 0 {
		limit = defaultLimit
	} else {
		limit = o.Limit
	}

	q.Set("limit", strconv.Itoa(limit))

	if o.PageToken != "" && o.PageNonce != "" {
		q.Set("__token", o.PageToken)
		q.Set("__nonce", o.PageNonce)
	}

	u.RawQuery = q.Encode()
}

func NewOptions() *Options {
	return &Options{
		queryOpts: make(url.Values),
	}
}
