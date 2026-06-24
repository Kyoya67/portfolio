package model

type Work struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Body         string   `json:"body"`
	ThumbnailURL string   `json:"thumbnailUrl,omitempty"`
	GitHubURL    string   `json:"githubUrl,omitempty"`
	ProjectURL   string   `json:"projectUrl,omitempty"`
	Technologies []string `json:"technologies"`
	Published    bool     `json:"published"`
	CreatedAt    string   `json:"createdAt"`
	UpdatedAt    string   `json:"updatedAt"`
}

type Article struct {
	ID          string   `json:"id"`
	Source      string   `json:"source"`
	SourceID    string   `json:"sourceId"`
	Title       string   `json:"title"`
	URL         string   `json:"url"`
	Excerpt     string   `json:"excerpt"`
	Tags        []string `json:"tags"`
	PublishedAt string   `json:"publishedAt"`
	FetchedAt   string   `json:"fetchedAt"`
	IsPublished bool     `json:"isPublished"`
}

type Profile struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Bio       string   `json:"bio"`
	Skills    []string `json:"skills"`
	GitHub    string   `json:"github,omitempty"`
	X         string   `json:"x,omitempty"`
	LinkedIn  string   `json:"linkedin,omitempty"`
	Website   string   `json:"website,omitempty"`
	UpdatedAt string   `json:"updatedAt"`
}
