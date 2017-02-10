package utube

import "net/url"

type YoutubeParam interface {
	Params() url.Values
}

type PageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}

type Thumbnail struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Thumbnails struct {
	Default  *Thumbnails `json:"default"`
	Medium   *Thumbnails `json:"medium"`
	High     *Thumbnails `json:"high"`
	Standard *Thumbnails `json:"standard"`
	Maxres   *Thumbnails `json:"maxres"`
}

type Localized struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
