package nationbuilder

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// BasicPage is an implementation of Nationbuilder's
// basic page type
type BasicPage struct {
	// The path at which to place the page. Must be unique, and there are some restrictions for namespace collisions.
	// (Optional- will be computed from name if not present)
	Slug string `json:"slug"`
	// The content html to have on the page. There is a whitelist of html elements and attributes that are allowed.
	Content string `json:"content"`
	// Valid values are 'published' or 'drafted'.  Required.
	Status string `json:"status"`
	// Internal name, how the page will be referred to in lists in the control panel (required)
	Name string `json:"name"`
	// Title of the page, shows up as tab name, for example (optional, defaults to the name)
	Title string `json:"title"`
	// Heading on the page (optional, defaults to the name)
	Headline string `json:"headline"`
	// Meta attribute for SEO - description (optional)
	Excerpt string `json:"excerpt"`
	// The unique identifier for this resource in an external service (optional)
	ExternalID string `json:"external_id"`
	// List of tags (optional)
	Tags []string `json:"tags"`
	// Page ID
	ID int `json:"id"`
}

// Nationbuilder returns content wrapped within an object
type BasicPageContainer struct {
	BasicPageContent *BasicPage `json:"basic_page"`
}

func newBasicPageContainerFromJSON(data []byte) (*BasicPage, error) {
	bp := &BasicPageContainer{}
	err := json.Unmarshal(data, bp)
	if err != nil {
		return nil, err
	}

	return bp.BasicPageContent, nil
}

type BasicPageEndpoint struct {
	url       url.URL
	basicPage *BasicPage
}

func (b *BasicPageEndpoint) GetURL() string {
	return b.url.String()
}

func (b *BasicPageEndpoint) GetBody() ([]byte, error) {
	return json.Marshal(&BasicPageContainer{
		BasicPageContent: b.basicPage,
	})
}

func (b *BasicPageEndpoint) Retrieve() *BasicPage {
	return b.basicPage
}

// func (b *BasicPageEndpoint) Update(bp *BasicPage) (*BasicPage, error) {

// }

// func (b *BasicPageEndpoint) Delete() error {

// }

func newBasicPageEndpoint(url url.URL, bp *BasicPage) *BasicPageEndpoint {
	return &BasicPageEndpoint{
		url:       url,
		basicPage: bp,
	}
}

// The basic page collection type
type BasicPageCollection struct {
	Page
	Results []*BasicPage `json:"results"`
}

// BasicPageEndpoint represents the basic page API
// Note - it is not, somewhat strangely, possible to
// directly retrieve a basic page by ID from the server.
// One must page through to find the right page
type basicPageCollectionEndpoint struct {
	url        url.URL
	limit      int
	pageNumber int
	pages      map[int]*BasicPageCollection
	basicPages map[int]*BasicPageEndpoint
}

func (b *basicPageCollectionEndpoint) GetURL() string {
	return b.url.String()
}

func (b *basicPageCollectionEndpoint) SetLimit(limit int) {
	b.limit = limit
	q := b.url.Query()
	q.Set("limit", strconv.Itoa(limit))

	b.url.RawQuery = q.Encode()
}

func (b *basicPageCollectionEndpoint) Retrieve() error {
	res := retrieve(b)
	if res.err != nil {
		return res.err
	}

	basicPages := &BasicPageCollection{}
	err := json.Unmarshal(res.body, basicPages)
	if err != nil {
		return err
	}

	b.pages[b.pageNumber] = basicPages

	for _, p := range basicPages.Results {
		u := b.url
		u.Path += fmt.Sprintf("/%d", p.ID)
		q := u.Query()
		q.Del("limit")
		u.RawQuery = q.Encode()
		b.basicPages[p.ID] = newBasicPageEndpoint(u, p)
	}

	return nil
}

func (b *basicPageCollectionEndpoint) Create(bp ...*BasicPage) ([]*BasicPage, error) {

	bpe := make([]*BasicPageEndpoint, 0)
	for _, p := range bp {
		bpe = append(bpe, newBasicPageEndpoint(b.url, p))
	}

	results := create(bpe...)

	if errors, ok := results.getErrors(); !ok {
		return errors
	} else {
		bps := make([]*BasicPage)
		for _, r := range results {
			p, err := NewBasicPageContainerFromJSON(r.body)
			if err != nil {
				// What?
			}
			b.basicPages[p.ID] = p
			bps = append(bps, p)
		}
		return nil
	}
}

func (b *basicPageCollectionEndpoint) BasicPage(id int) {

}

func newBasicPageCollectionEndpoint() *basicPageCollectionEndpoint {
	b := &basicPageCollectionEndpoint{
		url:        url.URL{},
		basicPages: make(map[int]*BasicPageEndpoint),
	}

	b.SetLimit(50)
}

// myNation := NewNation("slug", "key")
// pages := myNation.Site("myAwesomeSite").BasicPages()
// pages.Retrieve()
// pages.Next()
// pages.Prev()
// p, err := pages.BasicPage(7)
// p.Update(&BasicPage{})
// p.Delete()
