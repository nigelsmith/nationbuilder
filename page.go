package nationbuilder

// Common page fields
type Page struct {
	// Page slug - computer from name with mangling if a collision occurs. Optional field.
	Slug string `json:"slug,omitempty"`
	// Path of the page
	Path string `json:"path,omitempty"`
	// Page display status - 'published', 'drafted'.  Required field.
	Status string `json:"status,omitempty"`
	// The slug of the site to which the basic page belongs
	SiteSlug string `json:"site_slug,omitempty"`
	// Page name - it's from this that other values like the slug and headline are computed. Required field.
	Name string `json:"name,omitempty"`
	// Page headline.  Optional field.
	Headline string `json:"headline,omitempty"`
	// HTML page title - computed from the name if not provided. Optional field.
	Title string `json:"title,omitempty"`
	// Page excerpt - used for visiting bots. Optional field.
	Excerpt string `json:"excerpt,omitempty"`
	// Numeric ID of the person who authored the page
	AuthorID int `json:"author_id,omitempty"`
	// Page publication date in format described by DateFormat
	PublishedAt string `json:"published_at,omitempty"`
	// The external ID, if any, that a page has (e.g. a previous wordpress page). Optional field.
	ExternalID string `json:"external_id,omitempty"`
	// Page tags
	Tags []string `json:"tags,omitempty"`
	// Page ID
	ID int `json:"id,omitempty"`
}
