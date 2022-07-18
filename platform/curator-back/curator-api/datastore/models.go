package datastore

type ExternalLink struct {
	ID  string `json:"id"`
	Url string `json:"url"`
}
type CuratedRecord struct {
	ID            string         `json:"id"`
	Headline      string         `json:"headline"`
	Description   string         `json:"description"`
	ExternalLinks []ExternalLink `json:"external-links"`
}
