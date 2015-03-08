Nationbuilder
=============

A work in progress implementation of a client for the Nationbuilder remote API.

###Supported Endpoints

+ Basic Pages
+ Blogs
+ Blog Posts
+ Calendars
+ Site
+ Attachments
+ People

with more to come over time.

###Example Usage

Instantiate a client and fetch blog posts:

```go

myNation, err := nationbuilder.NewClient("nationSlug", "apiKey")
if err != nil {
	log.Fatal(err.Error())
}

blogID := 1
// Pass nil for options (default page size is set to 50)
blogPosts, result := myNation.GetBlogPosts("siteSlug", blogID, nil)
if result.HasError() {
	log.Fatal(result.Error())
}

for _, post := range blogPosts.Results {
	fmt.Println(post)
}
```

Full API documentation is available at: [godoc.org](https://godoc.org/github.com/nigelsmith/nationbuilder)