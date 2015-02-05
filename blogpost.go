package nationbuilder

type BlogPost struct {
	Id                int      `json:"id"`
	Name              string   `json:"name"`
	Title             string   `json:"title"`
	Headline          string   `json:"headline"`
	Slug              string   `json:"slug"`
	Status            string   `json:"status"`
	ContentBeforeFlip string   `json:"content_before_flip"`
	ContentAfterFlip  string   `json:"content_after_flip"`
	PublishedAt       string   `json:"published_at"`
	Tags              []string `json:"tags"`
}
