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

type YoutubeId struct {
	Kind       string `json:"kind"`
	ChannelId  string `json:"channelId"`
	VideoId    string `json:"videoId"`
	PlaylistId string `json:"playlistId"`
}

func (this *YoutubeId) IsVideo() bool {
	return this.Kind == "youtube#video"
}

func (this *YoutubeId) IsChannel() bool {
	return this.Kind == "youtube#channel"
}

func (this *YoutubeId) IsPlaylist() bool {
	return this.Kind == "youtube#playlist"
}

func (this *YoutubeId) GetId() string {
	if this.IsVideo() {
		return this.VideoId
	} else if this.IsChannel() {
		return this.ChannelId
	}
	return this.PlaylistId
}