package nationbuilder

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const apiVersion = "v1"
const debug = true
const defaultLimit = 50

type nationbuilderURL struct {
	u url.URL
}

func (n *nationbuilderURL) setQuery(key string, val string) {
	q := n.u.Query()
	q.Set(key, val)
	n.u.RawQuery = q.Encode()
}

func (n *nationbuilderURL) setLimit(limit int) {
	n.setQuery("limit", strconv.Itoa(limit))
}

func (n *nationbuilderURL) setToken(token string) {
	n.setQuery("access_token", token)
}

func (n *nationbuilderURL) extendPath(path string) {
	if len(path) > 0 {
		if string(path[0]) != "/" {
			n.u.Path += "/"
		}

		n.u.Path += path
	}
}

func (n *nationbuilderURL) String() string {
	return n.u.String()
}

type NationbuilderClient struct {
	Slug    string
	ApiKey  string
	baseURL *nationbuilderURL
	c       *http.Client
}

func (n *NationbuilderClient) getRequest(method string, path string, options *Options) *apiRequest {
	b := *n.baseURL
	b.u.Path += path

	if options != nil {
		options.setQuery(&b.u)
	}

	return &apiRequest{
		url:    b.String(),
		method: method,
	}

}

func (n *NationbuilderClient) SetClient(c *http.Client) {
	n.c = c
}

func NewNationbuilderClient(slug string, key string) (*NationbuilderClient, error) {
	u, err := url.Parse(fmt.Sprintf("https://%s.nationbuilder.com/api/%s", slug, apiVersion))
	if err != nil {
		return nil, err
	}

	nbURL := &nationbuilderURL{
		u: *u,
	}

	nbURL.setToken(key)

	return &NationbuilderClient{
		Slug:    slug,
		ApiKey:  key,
		baseURL: nbURL,
		c:       &http.Client{},
	}, nil
}
