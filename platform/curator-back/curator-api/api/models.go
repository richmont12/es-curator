package api

type externalLink struct {
	ID  string `json:"id"`
	Url string `json:"url"`
}

type curatedRecord struct {
	ID            string         `json:"id"`
	Headline      string         `json:"headline"`
	Description   string         `json:"description"`
	ExternalLinks []externalLink `json:"external-links"`
}
