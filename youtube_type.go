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
	Default  *Thumbnail `json:"default"`
	Medium   *Thumbnail `json:"medium"`
	High     *Thumbnail `json:"high"`
	Standard *Thumbnail `json:"standard"`
	Maxres   *Thumbnail `json:"maxres"`
}

func (this *Thumbnails) GetThumbnailURL() string {
	if this.Maxres != nil {
		return this.Maxres.URL
	} else if this.Standard != nil {
		return this.Standard.URL
	} else if this.High != nil {
		return this.High.URL
	} else if this.Medium != nil {
		return this.Medium.URL
	} else if this.Default != nil {
		return this.Default.URL
	}
	return ""
}

type Localized struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ResourceId struct {
	Kind       string `json:"kind"`
	ChannelId  string `json:"channelId"`
	VideoId    string `json:"videoId"`
	PlaylistId string `json:"playlistId"`
}

func (this *ResourceId) IsVideo() bool {
	return this.Kind == "youtube#video"
}

func (this *ResourceId) IsChannel() bool {
	return this.Kind == "youtube#channel"
}

func (this *ResourceId) IsPlaylist() bool {
	return this.Kind == "youtube#playlist"
}

func (this *ResourceId) GetId() string {
	if this.IsVideo() {
		return this.VideoId
	} else if this.IsChannel() {
		return this.ChannelId
	}
	return this.PlaylistId
}