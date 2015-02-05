package nationbuilder

import "errors"

const apiVersion = "v1"

var (
	errNoResponse = errors.New("No response from nationbuilder")
)

// func (a Api) CreateBlogPosts(site string, blogID string, posts ...*BlogPost) ([]*BlogPostResult, error) {
// 	requests := make([]*apiRequest, 0)
// 	u := a.getBaseURL()
// 	u.Path += "/" + site + "/pages/blogs/" + blogID + "/posts"

// 	for _, p := range posts {
// 		bpc := &BlogPostContainer{p}
// 		body, err := bpc.AsJSON()
// 		if err != nil {
// 			return nil, err
// 		}
// 		requests = append(requests, &apiRequest{u.String(), body, "POST"})
// 	}

// 	results := processRequests(requests...)
// 	if len(results) == 0 {
// 		return nil, errNoResponse
// 	}

// 	blogPosts := make([]*BlogPostResult, 0)
// 	for _, r := range results {
// 		if r.err != nil {
// 			return nil, r.err
// 		}
// 		if r.statusCode != http.StatusOK {
// 			apiErr := &ApiError{}
// 			err := json.Unmarshal(r.body, apiErr)
// 			if err != nil {
// 				return nil, err
// 			}
// 			log.Printf("Nationbuilder Error\nCode: %s\tMessage: %s\nHTTP Statuscode: %d\n", apiErr.Code, apiErr.Message, r.statusCode)
// 			return nil, errors.New("Nationbuilder Error encountered when creating blog post")
// 		}
// 		post, err := NewBlogPostResultFromJSON(r.body)
// 		if err != nil {
// 			return nil, err
// 		}
// 		blogPosts = append(blogPosts, post)
// 	}

// 	return blogPosts, nil

// }

// func (a Api) CreateBlog(blog BlogPage) {

// }

// func (a Api) GetBlogs(site string) (*BlogResult, error) {
// 	u := a.baseUrl
// 	u.Path += "/" + site + "/pages/blogs"

// 	results := processRequests(&apiRequest{
// 		method: "GET",
// 		url:    u.String(),
// 	})
// 	if len(results) == 0 {
// 		return nil, errNoResponse
// 	}

// 	r := results[0]
// 	if r.err != nil {
// 		return nil, r.err
// 	}

// 	return NewBlogResultFromJSON(r.body)

// }
