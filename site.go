package nationbuilder

import "fmt"

type Site struct {
	ID     int
	Name   string
	Slug   string
	Domain string
}

func (s *Site) String() string {
	return fmt.Sprintf("Site: %s", s.Name)
}

type SitePage struct {
	Results []*Site `json:"results"`
	Pagination
}

func (n *NationbuilderClient) GetSites(options *Options) (sites *SitePage, result *Result) {
	req := n.getRequest("GET", "/sites", options)
	result = n.retrieve(req, &sites)

	return
}
