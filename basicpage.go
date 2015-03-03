package nationbuilder

import (
	"fmt"
	"net/http"
)

// BasicPage is an implementation of Nationbuilder's
// basic page type
type BasicPage struct {
	Page
	// HTML for the page - will be run through a whitelist filter.
	Content string `json:"content,omitempty"`
}

func (b *BasicPage) String() string {
	return fmt.Sprintf("Basic Page: %s", b.Name)
}

// Paginated collection of basic pages
type BasicPages struct {
	Results []*BasicPage `json:"results"`
	Pagination
}

// Wrapper around basic page
type basicPageWrap struct {
	BasicPage *BasicPage `json:"basic_page"`
}

// Retrieve a page of Basic Pages from the specified site
func (n *NationbuilderClient) GetBasicPages(siteSlug string, options *Options) (basicPages *BasicPages, result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/basic_pages", siteSlug)
	req := n.getRequest("GET", u, options)
	result = n.retrieve(req, &basicPages)

	return
}

// Create a Basic Page for the specified site
func (n *NationbuilderClient) CreateBasicPage(siteSlug string, basicPage *BasicPage, options *Options) (newBasicPage *BasicPage, result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/basic_pages", siteSlug)
	req := n.getRequest("POST", u, options)
	bpw := &basicPageWrap{}
	result = n.create(&basicPageWrap{basicPage}, req, bpw, http.StatusOK)
	newBasicPage = bpw.BasicPage

	return
}

// Update a Basic Page
func (n *NationbuilderClient) UpdateBasicPage(siteSlug string, id int, basicPage *BasicPage, options *Options) (newBasicPage *BasicPage, result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/basic_pages/%d", siteSlug, id)
	req := n.getRequest("PUT", u, options)
	bpw := &basicPageWrap{}
	result = n.create(&basicPageWrap{basicPage}, req, bpw, http.StatusOK)
	newBasicPage = bpw.BasicPage

	return
}

// Delete a Basic Page
func (n *NationbuilderClient) DeleteBasicPage(siteSlug string, id int) (result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/basic_pages/%d", siteSlug, id)
	req := n.getRequest("DELETE", u, nil)
	result = n.delete(req)

	return
}
