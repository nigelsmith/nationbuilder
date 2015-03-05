package nationbuilder

import (
	"fmt"
	"net/http"
)

// A nationbuilder blog - the model for which doesn't vary from the page type
type Blog struct {
	Page
}

func (b *Blog) String() string {
	return fmt.Sprintf("Blog: %s", b.Name)
}

// A 'page' of Nationbuilder Blogs
type Blogs struct {
	Results []*Blog `json:"results"`
	Pagination
}

// Wrapper around blogs
type blogWrap struct {
	Blog *Blog `json:"blog"`
}

// Retrieve a page of blogs from the specified site
func (n *NationbuilderClient) GetBlogs(siteSlug string, options *Options) (blogs *Blogs, result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/blogs", siteSlug)
	req := n.getRequest("GET", u, options)
	result = n.retrieve(req, &blogs)

	return
}

// Create a blog for the specified site
func (n *NationbuilderClient) CreateBlog(siteSlug string, blog *Blog, options *Options) (newBlog *Blog, result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/blogs", siteSlug)
	req := n.getRequest("POST", u, options)
	bw := &blogWrap{}
	result = n.create(&blogWrap{blog}, req, bw, http.StatusOK)
	newBlog = bw.Blog

	return
}

// Update a blog
func (n *NationbuilderClient) UpdateBlog(siteSlug string, id int, blog *Blog, options *Options) (updatedBlog *Blog, result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/blogs/%d", siteSlug, id)
	req := n.getRequest("PUT", u, options)
	bw := &blogWrap{}
	result = n.create(&blogWrap{blog}, req, bw, http.StatusOK)
	updatedBlog = bw.Blog

	return
}

// Delete a blog
func (n *NationbuilderClient) DeleteBlog(siteSlug string, id int) (result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/blogs/%d", siteSlug, id)
	req := n.getRequest("DELETE", u, nil)
	result = n.delete(req)

	return
}
