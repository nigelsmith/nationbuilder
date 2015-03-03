package nationbuilder

import "fmt"

// Site represent a single nationbuilder site
type Site struct {
	ID     int
	Name   string
	Slug   string
	Domain string
}

func (s *Site) String() string {
	return fmt.Sprintf("Site: %s", s.Name)
}

// A paginated collection of sites
type SitePage struct {
	Results []*Site `json:"results"`
	Pagination
}

// Retrieve sites
func (n *NationbuilderClient) GetSites(options *Options) (sites *SitePage, result *Result) {
	req := n.getRequest("GET", "/sites", options)
	result = n.retrieve(req, &sites)

	return
}
