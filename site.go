package nationbuilder

import "net/url"

type Site struct {
	url *url.URL
}

func (s *Site) BasicPages() *BasicPageCollectionEndpoint {
	u := s.url
	u.Path += "pages/basic_pages"

	bp := newBasicPageCollectionEndpoint()
	bp.url = u

	return bp
}
