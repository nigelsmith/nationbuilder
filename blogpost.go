package nationbuilder

import (
	"fmt"
	"net/http"
)

// Nationbuilder's blog post type
type BlogPost struct {
	Page
	// HTML content to be displayed as 'teaser' text
	ContentBeforeFlip string `json:"content_before_flip,omitempty"`
	// HTML content to be concatenated to the teaser text
	ContentAfterFlip string `json:"content_after_flip,omitempty"`
}

func (b *BlogPost) String() string {
	return fmt.Sprintf("ID: %d, Blog Post: %s", b.ID, b.Name)
}

// A paginated collection of blog posts
type BlogPosts struct {
	Results []*BlogPost `json:"results"`
	Pagination
}

type blogPostWrap struct {
	BlogPost *BlogPost `json:"blog_post"`
}

// Retrieve a page of Blog Posts for the given site and blog id
func (n *Client) GetBlogPosts(siteSlug string, id int, options *Options) (blogPosts *BlogPosts, result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/blogs/%d/posts", siteSlug, id)
	req := n.getRequest("GET", u, options)
	result = n.retrieve(req, &blogPosts)

	return
}

// Retrieve an individual Blog Post
func (n *Client) GetBlogPost(siteSlug string, blogID int, postID int, options *Options) (blogPost *BlogPost, result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/blogs/%d/posts/%d", siteSlug, blogID, postID)
	req := n.getRequest("GET", u, options)
	b := &blogPostWrap{}
	result = n.retrieve(req, b)
	blogPost = b.BlogPost

	return
}

// Create a Blog Post for the specified site and blog id
func (n *Client) CreateBlogPost(siteSlug string, id int, blogPost *BlogPost, options *Options) (newBlogPost *BlogPost, result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/blogs/%d/posts", siteSlug, id)
	req := n.getRequest("POST", u, options)
	bpw := &blogPostWrap{}
	result = n.create(&blogPostWrap{blogPost}, req, bpw, http.StatusOK)
	newBlogPost = bpw.BlogPost

	return
}

// Update a Blog Post for the specified site and blog id
func (n *Client) UpdateBlogPost(siteSlug string, blogID int, postID int, blogPost *BlogPost, options *Options) (newBlogPost *BlogPost, result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/blogs/%d/posts/%d", siteSlug, blogID, postID)
	req := n.getRequest("PUT", u, options)
	bpw := &blogPostWrap{}
	result = n.create(&blogPostWrap{blogPost}, req, bpw, http.StatusOK)
	newBlogPost = bpw.BlogPost

	return
}

// Delete a Blog Post
func (n *Client) DeleteBlogPost(siteSlug string, blogID int, postID int, options *Options) (result *Result) {
	u := fmt.Sprintf("/sites/%s/pages/blogs/%d/posts/%d", siteSlug, blogID, postID)
	req := n.getRequest("DELETE", u, options)
	result = n.delete(req)

	return
}
