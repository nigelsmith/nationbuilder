package nationbuilder

import (
	"fmt"
	"net/url"
)

// Nation represents an instance of a Nationbuilder Nation
type Nation struct {
	Slug    string
	Key     string
	baseUrl *url.URL
	sites   map[string]*Site
}

func (n *Nation) Site(siteSlug string) *Site {

	if s, ok := n.sites[siteSlug]; !ok {
		u := n.baseUrl
		u.Path += "sites/" + siteSlug
		s := &Site{
			url: u,
		}
		n.sites[siteSlug] = s

		return s
	} else {
		return s
	}

}

func (n *Nation) getBaseURL() *url.URL {
	return n.baseUrl
}

// NationbuilderError represents the structure returned by Nationbuilder when an error is encountered
type NationbuilderError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (n *NationbuilderError) Error() string {
	return fmt.Sprintf("Code %s: %s", n.Code, n.Message)
}

func NewNation(slug, key string) (*Nation, error) {
	n := &Nation{
		Slug:  slug,
		Key:   key,
		sites: make(map[string]*Site),
	}
	u, err := url.Parse("https://" + slug + ".nationbuilder.com/api/" + apiVersion + "/")
	if err != nil {
		return nil, err
	}

	n.baseUrl = u

	q := n.baseUrl.Query()
	q.Set("access_token", n.Key)
	n.baseUrl.RawQuery = q.Encode()

	return n, nil
}
