package nationbuilder

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const apiKey = "testkey"
const slug = "testslug"
const siteSlug = "testSite"
const testName = "testName"
const testID = 1

var c *NationbuilderClient

var (
	basicPagesURL   = fmt.Sprintf("/api/v1/sites/%s/pages/basic_pages", siteSlug)
	basicPageURL    = fmt.Sprintf("/api/v1/sites/%s/pages/basic_pages/%d", siteSlug, testID)
	blogPostsURL    = fmt.Sprintf("/api/v1/sites/%s/pages/blogs/%d/posts", siteSlug, testID)
	blogPostURL     = fmt.Sprintf("/api/v1/sites/%s/pages/blogs/%d/posts/%d", siteSlug, testID, testID)
	blogsURL        = fmt.Sprintf("/api/v1/sites/%s/pages/blogs", siteSlug)
	blogURL         = fmt.Sprintf("/api/v1/sites/%s/pages/blogs/%d", siteSlug, testID)
	calendarsURL    = fmt.Sprintf("/api/v1/sites/%s/pages/calendars", siteSlug)
	calendarURL     = fmt.Sprintf("/api/v1/sites/%s/pages/calendars/%d", siteSlug, testID)
	peopleURL       = "/api/v1/people"
	personURL       = fmt.Sprintf("/api/v1/people/%d", testID)
	peopleNearbyURL = "/api/v1/people/nearby"
	peopleSearchURL = "/api/v1/people/search"
	sitesURL        = "/api/v1/sites"
)

func sitesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		err := json.NewEncoder(w).Encode(&Sites{
			Results: []*Site{
				&Site{
					Name: testName,
				},
			},
		})
		if err != nil {
			log.Fatal(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func peopleSearchHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		q := r.URL.Query()
		fName := q.Get("first_name")
		lName := q.Get("last_name")
		if fName == "" || lName == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err := json.NewEncoder(w).Encode(&People{
			Results: []*Person{
				&Person{
					FirstName: "Wally",
					LastName:  "Waldo",
				},
			},
		})
		if err != nil {
			log.Fatal(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func peopleNearbyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		q := r.URL.Query()
		loc := q.Get("location")
		if loc == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		dist := q.Get("distance")
		if dist == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err := json.NewEncoder(w).Encode(&People{
			Results: []*Person{
				&Person{
					FirstName: "Wally",
					LastName:  "Waldo",
				},
			},
		})
		if err != nil {
			log.Fatal(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		person := &Person{
			FirstName: "Phileas",
			LastName:  "Fogg",
		}
		err := json.NewEncoder(w).Encode(&personWrap{person})
		if err != nil {
			log.Fatal(err.Error())
		}
	case "PUT":
		pw := &personWrap{}
		err := json.NewDecoder(r.Body).Decode(pw)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = json.NewEncoder(w).Encode(pw)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "DELETE":
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		people := &People{
			Results: []*Person{
				&Person{
					ID:        testID,
					FirstName: "Phileas",
					LastName:  "Fogg",
				},
			},
		}
		err := json.NewEncoder(w).Encode(people)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "POST":
		pw := &personWrap{}
		err := json.NewDecoder(r.Body).Decode(pw)
		if err != nil {
			log.Fatal(err.Error())
		}
		pw.Person.ID = testID
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(pw)
		if err != nil {
			log.Fatal(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func calendarHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		cw := &calendarWrap{
			Calendar: &Calendar{
				Page: Page{
					Name: testName,
					ID:   testID,
				},
			},
		}
		err := json.NewEncoder(w).Encode(cw)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "PUT":
		cw := &calendarWrap{}
		err := json.NewDecoder(r.Body).Decode(cw)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = json.NewEncoder(w).Encode(cw)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "DELETE":
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func calendarsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c := &Calendars{
			Results: []*Calendar{
				&Calendar{
					Page: Page{
						Name: testName,
						ID:   testID,
					},
				},
			},
		}
		err := json.NewEncoder(w).Encode(c)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "POST":
		cw := &calendarWrap{}
		err := json.NewDecoder(r.Body).Decode(cw)
		if err != nil {
			log.Fatal(err.Error())
		}
		cw.Calendar.ID = testID
		err = json.NewEncoder(w).Encode(cw)
		if err != nil {
			log.Fatal(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

var basicPagesHandler = func(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		pages := &BasicPages{
			Results: []*BasicPage{
				&BasicPage{
					Page: Page{
						Name: testName,
					},
				},
			},
		}
		err := json.NewEncoder(w).Encode(pages)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	if req.Method == "POST" {
		bpw := &basicPageWrap{}
		err := json.NewDecoder(req.Body).Decode(bpw)
		if err != nil {
			log.Fatal(err.Error())
		}

		if bpw.BasicPage.Name != testName {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			err := json.NewEncoder(w).Encode(bpw)
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	}
}

var basicPageHandler = func(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "PUT":
		bpw := &basicPageWrap{}
		err := json.NewDecoder(req.Body).Decode(bpw)
		if err != nil {
			log.Fatal(err.Error())
		}
		page := bpw.BasicPage
		if page.Name != testName {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			err := json.NewEncoder(w).Encode(bpw)
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	case "DELETE":
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

var blogPostsHandler = func(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		posts := &BlogPosts{
			Results: []*BlogPost{
				&BlogPost{
					Page: Page{
						Name: testName,
						ID:   testID,
					},
				},
			},
		}
		err := json.NewEncoder(w).Encode(posts)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "POST":
		pw := &blogPostWrap{}
		err := json.NewDecoder(req.Body).Decode(pw)
		if err != nil {
			log.Fatal(err.Error())
		}
		pw.BlogPost.ID = testID
		err = json.NewEncoder(w).Encode(pw)
		if err != nil {
			log.Fatal(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func blogPostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		p := &BlogPost{
			Page: Page{
				Name: testName,
			},
		}
		err := json.NewEncoder(w).Encode(&blogPostWrap{p})
		if err != nil {
			log.Fatal(err.Error())
		}
	case "PUT":
		pw := &blogPostWrap{}
		err := json.NewDecoder(r.Body).Decode(pw)
		if err != nil {
			log.Fatal(err.Error())
		}
		if pw.BlogPost.Name != testName {
			w.WriteHeader(http.StatusBadRequest)
		}
		err = json.NewEncoder(w).Encode(pw)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "DELETE":
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func blogsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		blogs := &Blogs{
			Results: []*Blog{
				&Blog{
					Page: Page{
						Name: testName,
					},
				},
			},
		}

		err := json.NewEncoder(w).Encode(blogs)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "POST":
		bw := &blogWrap{}
		err := json.NewDecoder(r.Body).Decode(bw)
		if err != nil {
			log.Fatal(err.Error())
		}
		bw.Blog.ID = testID
		err = json.NewEncoder(w).Encode(bw)
		if err != nil {
			log.Fatal(err.Error())
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		bw := &blogWrap{}
		err := json.NewDecoder(r.Body).Decode(bw)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = json.NewEncoder(w).Encode(bw)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "DELETE":
		w.WriteHeader(http.StatusNoContent)
	}
}

func init() {
	client, err := NewNationbuilderClient(slug, apiKey)
	if err != nil {
		log.Fatal(err.Error())
	}
	c = client

	apiMux := http.NewServeMux()
	apiMux.HandleFunc(basicPagesURL, basicPagesHandler)
	apiMux.HandleFunc(basicPageURL, basicPageHandler)
	apiMux.HandleFunc(blogPostsURL, blogPostsHandler)
	apiMux.HandleFunc(blogPostURL, blogPostHandler)
	apiMux.HandleFunc(blogsURL, blogsHandler)
	apiMux.HandleFunc(blogURL, blogHandler)
	apiMux.HandleFunc(calendarsURL, calendarsHandler)
	apiMux.HandleFunc(calendarURL, calendarHandler)
	apiMux.HandleFunc(peopleURL, peopleHandler)
	apiMux.HandleFunc(personURL, personHandler)
	apiMux.HandleFunc(peopleNearbyURL, peopleNearbyHandler)
	apiMux.HandleFunc(peopleSearchURL, peopleSearchHandler)
	apiMux.HandleFunc(sitesURL, sitesHandler)

	server := httptest.NewServer(apiMux)

	u, err := url.Parse(server.URL)
	if err != nil {
		log.Fatal(err.Error())
	}
	u.Path += "/api/v1"
	c.baseURL = &nationbuilderURL{*u}
}

func TestNationbuilderURLExtendPath(t *testing.T) {
	n, err := NewNationbuilderClient(slug, apiKey)
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedPath := "/api/v1"
	url := n.baseURL

	if url.u.Path != expectedPath {
		t.Errorf("Expected path %s but saw %s", expectedPath, url.u.Path)
	}

	url.extendPath("/foo")
	expectedPath = "/api/v1/foo"

	if url.u.Path != expectedPath {
		t.Errorf("Expected extended path to be %s but saw %s", expectedPath, url.u.Path)
	}

	url.extendPath("bar")
	expectedPath = "/api/v1/foo/bar"
	if url.u.Path != expectedPath {
		t.Errorf("Expected extended path to be %s but saw %s", expectedPath, url.u.Path)
	}
}

func TestNewNationbuilderClient(t *testing.T) {
	c, err := NewNationbuilderClient(slug, apiKey)
	if err != nil {
		t.Fatal(err.Error())
	}

	if c.Slug != slug {
		t.Errorf("expected client to have slug %s but saw %s", slug, c.Slug)
	}

	if c.ApiKey != apiKey {
		t.Errorf("expected client to have api key %s but saw %s", apiKey, c.ApiKey)
	}

}

func TestBasicPagesGet(t *testing.T) {
	pages, result := c.GetBasicPages(siteSlug, nil)
	if result.HasError() {
		t.Error(result.Error())
	}

	if result.StatusCode != http.StatusOK {
		t.Errorf("Retrieving basic pages - expect status %d but saw %d", http.StatusOK, result.StatusCode)
	}

	if pages != nil {
		if len(pages.Results) != 1 {
			t.Errorf("Expected number of results to be %d but received %d results", 1, len(pages.Results))
		}

		if pages.Results[0].Name != testName {
			t.Errorf("Expected name of %s but saw %s", testName, pages.Results[0].Name)
		}

	} else {
		t.Error("Expected page of basic pages but received no response")
	}
}

func TestBasicPagesCreate(t *testing.T) {
	bp := &BasicPage{
		Page: Page{
			Name: testName,
		},
	}
	page, result := c.CreateBasicPage(siteSlug, bp, nil)
	if result.HasError() {
		t.Errorf("Unexpected error - code %d, Err: %s", result.StatusCode, result.Error())
	}

	if page == nil {
		t.Error("Expected new page but received a nil response")
	} else {
		if page.Name != testName {
			t.Errorf("Expected page name to be %s but saw %s", testName, page.Name)
		}
	}

	if result.StatusCode != http.StatusOK {
		t.Errorf("NB BasicPages returns a 200 on page creation, not 201 - make sure test reflects that: saw code %d", result.StatusCode)
	}
}

func TestBasicPageUpdate(t *testing.T) {
	bp := &BasicPage{
		Page: Page{
			Name: testName,
		},
	}
	page, result := c.UpdateBasicPage(siteSlug, testID, bp, nil)
	if result.HasError() {
		t.Errorf("Unexpected error - code %d, Err: %s", result.StatusCode, result.Error())
	}

	if page == nil {
		t.Error("Expected new page but received a nil response")
	} else {
		if page.Name != testName {
			t.Errorf("Expected page name to be %s but saw %s", testName, page.Name)
		}
	}

	if result.StatusCode != http.StatusOK {
		t.Errorf("NB BasicPages returns a 200 on page creation, not 201 - make sure test reflects that: saw code %d", result.StatusCode)
	}
}

func TestBasicPageDelete(t *testing.T) {
	result := c.DeleteBasicPage(siteSlug, testID)
	if result.HasError() {
		t.Errorf("Unexpected error - code %d, Err: %s", result.StatusCode, result.Error())
	}

	if result.StatusCode != http.StatusNoContent {
		t.Errorf("NB delete returns a 204 on resource deletion - make sure test reflects that: saw code %d", result.StatusCode)
	}
}

func TestBlogPostsGet(t *testing.T) {
	posts, result := c.GetBlogPosts(siteSlug, testID, nil)
	if result.HasError() {
		t.Errorf("Unexpected error - code %d, Err: %s", result.StatusCode, result.Error())
	}

	if posts == nil {
		t.Error("Unexpected nil posts response")
		t.SkipNow()
	}

	if len(posts.Results) != 1 {
		t.Error("Check API test - should return 1 post")
	}

	pName := posts.Results[0].Name
	if pName != testName {
		t.Errorf("Expected name of %s but saw %s", testName, pName)
	}
}

func TestBlogPostsCreate(t *testing.T) {
	p := &BlogPost{
		Page: Page{
			Name: testName,
		},
	}
	post, result := c.CreateBlogPost(siteSlug, testID, p, nil)
	if result.HasError() {
		t.Errorf("Unexpected error creating blog post: %s", result.Error())
		t.SkipNow()
	}

	if post == nil {
		t.Error("Unexpeced nil post - check test api")
		t.SkipNow()
	}

	name, expected := post.Name, testName
	if name != expected {
		t.Errorf("Expected post to have name %s but saw %s: check API test server", expected, name)
	}
}

func TestBlogPostGet(t *testing.T) {
	post, result := c.GetBlogPost(siteSlug, testID, testID, nil)
	if result.HasError() {
		t.Errorf("Unexpected error creating blog post: %s", result.Error())
		t.SkipNow()
	}

	name, expected := post.Name, testName
	if name != expected {
		t.Errorf("Expected %s but saw %s", expected, name)
	}
}

func TestBlogPostUpdate(t *testing.T) {
	p := &BlogPost{
		Page: Page{
			Name: testName,
		},
	}
	post, result := c.UpdateBlogPost(siteSlug, testID, testID, p, nil)
	if result.HasError() {
		t.Errorf("Unexpected error updating blog post: %s", result.Error())
		t.SkipNow()
	}

	if post == nil {
		t.Error("Unexpeced nil post - check test api")
		t.SkipNow()
	}

	name, expected := post.Name, testName
	if name != expected {
		t.Errorf("Expected post to have name %s but saw %s: check API test server", expected, name)
	}
}

func TestBlogPostDelete(t *testing.T) {
	result := c.DeleteBlogPost(siteSlug, testID, testID)
	if result.HasError() {
		t.Errorf("Unexpected error deleting blog post: %s", result.Error())
		t.SkipNow()
	}

	if result.StatusCode != http.StatusNoContent {
		t.Errorf("NB delete returns a 204 on resource deletion - make sure test reflects that: saw code %d", result.StatusCode)
	}
}

func TestBlogsGet(t *testing.T) {
	blogs, result := c.GetBlogs(siteSlug, nil)
	if result.HasError() {
		t.Errorf("Unexpected error retrieving blog: %s", result.Error())
		t.SkipNow()
	}

	if blogs == nil {
		t.Error("Unexpected nil blogs response")
		t.SkipNow()
	}

	if len(blogs.Results) != 1 {
		t.Error("Check API test - should return 1")
	}

	pName := blogs.Results[0].Name
	if pName != testName {
		t.Errorf("Expected name of %s but saw %s", testName, pName)
	}
}

func TestBlogsCreate(t *testing.T) {
	blog := &Blog{
		Page: Page{
			Name: testName,
		},
	}

	newBlog, result := c.CreateBlog(siteSlug, blog, nil)
	if result.HasError() {
		t.Errorf("Unexpected error creating blog: %s", result.Error())
		t.SkipNow()
	}

	if newBlog == nil {
		t.Error("Unexpected nil newblog response")
		t.SkipNow()
	} else {
		name, expected := newBlog.Name, testName
		if name != expected {
			t.Errorf("Expected new blog to have name %s but saw %s - check test API", testName, name)
		}
	}
}

func TestBlogUpdate(t *testing.T) {
	blog := &Blog{
		Page: Page{
			Name: testName,
		},
	}

	updatedBlog, result := c.UpdateBlog(siteSlug, testID, blog, nil)
	if result.HasError() {
		t.Errorf("Unexpected error retrieving blog: %s", result.Error())
		t.SkipNow()
	}

	if updatedBlog == nil {
		t.Error("Unexpected nil updatedBlog response")
		t.SkipNow()
	} else {
		name, expected := updatedBlog.Name, testName
		if name != expected {
			t.Errorf("Expected updated blog to have name %s but saw %s - check test API", testName, name)
		}
	}
}

func TestBlogDelete(t *testing.T) {
	result := c.DeleteBlog(siteSlug, testID)
	if result.HasError() {
		t.Errorf("Unexpected error deleting blog: %s", result.Error())
		t.SkipNow()
	}

	if result.StatusCode != http.StatusNoContent {
		t.Errorf("NB delete returns a 204 on resource deletion - make sure test reflects that: saw code %d", result.StatusCode)
	}
}

func TestCalendarsGet(t *testing.T) {
	calendars, result := c.GetCalendars(siteSlug, nil)
	if result.HasError() {
		t.Errorf("Unexpected error retrieving calendars: %s", result.Error())
		t.SkipNow()
	}

	if calendars == nil {
		t.Error("Unexpected nil calendars response")
		t.SkipNow()
	}

	if len(calendars.Results) != 1 {
		t.Error("Check API test - should return 1")
	}

	pName := calendars.Results[0].Name
	if pName != testName {
		t.Errorf("Expected name of %s but saw %s", testName, pName)
	}
}

func TestCalendarCreate(t *testing.T) {
	cal := &Calendar{
		Page: Page{
			Name: testName,
		},
	}
	newCal, result := c.CreateCalendar(siteSlug, cal, nil)
	if result.HasError() {
		t.Errorf("Unexpected error creating calendar: %s", result.Error())
		t.SkipNow()
	}

	if newCal == nil {
		t.Error("Unexpected nil calendar response")
		t.SkipNow()
	} else {
		name, expected := newCal.Name, testName
		if name != expected {
			t.Errorf("Expected new calendar to have name %s but saw %s - check test API", testName, name)
		}
	}
}

func TestCalendarUpdate(t *testing.T) {
	cal := &Calendar{
		Page: Page{
			Name: testName,
		},
	}

	updatedCal, result := c.UpdateCalendar(siteSlug, testID, cal, nil)
	if result.HasError() {
		t.Errorf("Unexpected error updating calendar: %s", result.Error())
		t.SkipNow()
	}

	if updatedCal == nil {
		t.Error("Unexpected nil updatedCal response")
		t.SkipNow()
	} else {
		name, expected := updatedCal.Name, testName
		if name != expected {
			t.Errorf("Expected updated blog to have name %s but saw %s - check test API", testName, name)
		}
	}
}

func TestCalendarDelete(t *testing.T) {
	result := c.DeleteCalendar(siteSlug, testID)
	if result.HasError() {
		t.Errorf("Unexpected error deleting calendar: %s", result.Error())
		t.SkipNow()
	}

	if result.StatusCode != http.StatusNoContent {
		t.Errorf("NB delete returns a 204 on resource deletion - make sure test reflects that: saw code %d", result.StatusCode)
	}
}

func TestCalendarGet(t *testing.T) {
	cal, result := c.GetCalendar(siteSlug, testID, nil)
	if result.HasError() {
		t.Errorf("Unexpected error retrieving calendar: %s", result.Error())
		t.SkipNow()
	}

	name, expected := cal.Name, testName
	if name != expected {
		t.Errorf("Expected %s but saw %s", expected, name)
	}
}

func TestPeopleGet(t *testing.T) {
	people, result := c.GetPeople(nil)
	if result.HasError() {
		t.Error(result.Error())
		t.SkipNow()
	}

	if result.StatusCode != http.StatusOK {
		t.Errorf("Retrieving people - expect status %d but saw %d", http.StatusOK, result.StatusCode)
	}

	if people != nil {
		if len(people.Results) != 1 {
			t.Errorf("Expected number of results to be %d but received %d results", 1, len(people.Results))
		}

		p := people.Results[0]
		if p.FirstName != "Phileas" && p.LastName != "Fogg" {
			t.Errorf("Expected name of Phileas Fogg but saw %s %s", p.FirstName, p.LastName)
		}

	} else {
		t.Error("Expected page of people but received no response")
	}
}

func TestPersonCreate(t *testing.T) {
	p := &Person{
		FirstName: "Phileas",
		LastName:  "Fogg",
	}
	newP, result := c.CreatePerson(p, nil)
	if result.HasError() {
		t.Errorf("Unexpected error creating person: %s", result.Error())
		t.SkipNow()
	}

	if newP == nil {
		t.Error("Unexpected nil person response")
		t.SkipNow()
	} else {
		firstName, lastName, expectedFirstName, expectedLastName := newP.FirstName, newP.LastName, "Phileas", "Fogg"
		if firstName != expectedFirstName || lastName != expectedLastName {
			t.Errorf("Expected new person to have name %s %s but saw %s %s - check test API", expectedFirstName, expectedLastName, firstName, lastName)
		}
	}
}

func TestPersonGet(t *testing.T) {
	person, result := c.GetPerson(testID, nil)
	if result.HasError() {
		t.Errorf("Unexpected error retrieving person: %s", result.Error())
		t.SkipNow()
	}

	if person == nil {
		t.Error("Unexpected nil person response")
		t.SkipNow()
	}
}

func TestPersonUpdate(t *testing.T) {
	p := &Person{
		FirstName: "Slarti",
		LastName:  "Bartfast",
	}
	updatedPerson, result := c.UpdatePerson(testID, p, nil)
	if result.HasError() {
		t.Errorf("Unexpected error updating person: %s", result.Error())
		t.SkipNow()
	}

	if updatedPerson == nil {
		t.Error("Unexpected nil person response")
		t.SkipNow()
	} else {
		firstName, lastName, expectedFirstName, expectedLastName := updatedPerson.FirstName, updatedPerson.LastName, "Slarti", "Bartfast"
		if firstName != expectedFirstName || lastName != expectedLastName {
			t.Errorf("Expected updated person to have name %s %s but saw %s %s - check test API", expectedFirstName, expectedLastName, firstName, lastName)
		}
	}
}

func TestPersonDelete(t *testing.T) {
	result := c.DeletePerson(testID, nil)
	if result.HasError() {
		t.Errorf("Unexpected error deleting person: %s", result.Error())
		t.SkipNow()
	}

	if result.StatusCode != http.StatusNoContent {
		t.Errorf("NB delete returns a 204 on resource deletion - make sure test reflects that: saw code %d", result.StatusCode)
	}
}

func TestNearbyPeople(t *testing.T) {
	people, result := c.NearbyPeople(-78.301233, 14.129705, 10, nil)
	if result.HasError() {
		t.Errorf("Unexpected error finding nearby people: %s", result.Error())
		t.SkipNow()
	}

	if people == nil {
		t.Error("Unexpected nil people response")
		t.SkipNow()
	}
}

func TestSearchPeople(t *testing.T) {
	opts := &PeopleSearchOptions{
		FirstName: "Wally",
		LastName:  "Waldo",
	}
	people, result := c.SearchPeople(opts, nil)

	if result.HasError() {
		t.Errorf("Unexpected error searching people: %s", result.Error())
		t.SkipNow()
	}

	if people == nil {
		t.Error("Unexpected nil people response")
		t.SkipNow()
	}
}

func TestSitesGet(t *testing.T) {
	sites, result := c.GetSites(nil)
	if result.HasError() {
		t.Errorf("Unexpected error fetching sites: %s", result.Error())
		t.SkipNow()
	}

	if sites == nil {
		t.Error("Unexpected nil sites response")
	}
}
