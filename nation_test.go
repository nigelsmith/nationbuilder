package nationbuilder

import (
	"testing"
)

func TestNation(t *testing.T) {
	tSlug := "testslug"
	tKey := "1234"
	n := NewNation(tSlug, tKey)

	if n.Slug != tSlug {
		t.Errorf("Expected Nation to have slug %s but saw %s", tSlug, n.Slug)
	}

	if n.Key != tKey {
		t.Errorf("Expected Nation to have key %s but saw %s", tKey, n.Key)
	}

	baseUrl := n.getBaseURL().String()
	eURL := "https://" + tSlug + ".nationbuilder.com/api/v1/?access_token=" + tKey
	if baseUrl != eURL {
		t.Errorf("Expected a base URL of %s but saw %s", eURL, baseUrl)
	}

	s := n.Site("testSite")
	sURL := s.url.String()
	eSiteURL := "https://" + tSlug + ".nationbuilder.com/api/v1/sites/" + "testSite" + "?access_token=" + tKey
	if sURL != eSiteURL {
		t.Errorf("Expected site with URL %s to be created but saw %s", eSiteURL, sURL)
	}

}
