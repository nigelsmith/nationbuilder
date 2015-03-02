package nationbuilder

import (
	"fmt"
	"net/http"
)

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

// Wrapper around basic page
type BlogWrap struct {
	Blog *Blog `json:"blog"`
}

// Retrieve a page of Basic Pages from the specified site
func (n *NationbuilderClient) GetBlogs(siteSlug string, options *Options) (blogs *Blogs, result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/blogs", siteSlug)
	req := n.getRequest("GET", u, options)
	result = n.retrieve(req, &blogs)

	return
}

// Create a Basic Page for the specified site
func (n *NationbuilderClient) CreateBlog(siteSlug string, blog *Blog, options *Options) (newBlog *Blog, result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/blogs", siteSlug)
	req := n.getRequest("POST", u, options)
	bw := &BlogWrap{}
	result = n.create(&BlogWrap{blog}, req, bw, http.StatusOK)
	newBlog = bw.Blog

	return
}

// Update a Basic Page
func (n *NationbuilderClient) UpdateBlog(siteSlug string, id int, blog *Blog, options *Options) (updatedBlog *Blog, result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/blogs/%d", siteSlug, id)
	req := n.getRequest("PUT", u, options)
	bw := &BlogWrap{}
	result = n.create(&BlogWrap{blog}, req, bw, http.StatusOK)
	updatedBlog = bw.Blog

	return
}

// Delete a Basic Page
func (n *NationbuilderClient) DeleteBlog(siteSlug string, id int) (result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/blogs/%d", siteSlug, id)
	req := n.getRequest("DELETE", u, nil)
	result = n.delete(req)

	return
}
