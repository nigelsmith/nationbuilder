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
type Sites struct {
	Results []*Site `json:"results"`
	Pagination
}

// Retrieve sites
func (n *NationbuilderClient) GetSites(options *Options) (sites *Sites, result *Result) {
	req := n.getRequest("GET", "/sites", options)
	result = n.retrieve(req, &sites)

	return
}
