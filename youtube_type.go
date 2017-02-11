package utube

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"net/url"
	"strings"
)

type YoutubeParam interface {
	Params() url.Values
}

type Part struct {
	value []string
}

func (this *Part) ShowAuditDetails() *Part {
	this.Append("auditDetails")
	return this
}

func (this *Part) ShowBrandingSettings() *Part {
	this.Append("brandingSettings")
	return this
}

func (this *Part) ShowContentDetails() *Part {
	this.Append("contentDetails")
	return this
}

func (this *Part) ShowContentOwnerDetails() *Part {
	this.Append("contentOwnerDetails")
	return this
}

func (this *Part) ShowId() *Part {
	this.Append("id")
	return this
}

func (this *Part) ShowInVideoPromotion() *Part {
	this.Append("invideoPromotion")
	return this
}

func (this *Part) ShowLocalizations() *Part {
	this.Append("localizations")
	return this
}

func (this *Part) ShowPlayer() *Part {
	this.Append("player")
	return this
}

func (this *Part) ShowSnippet() *Part {
	this.Append("snippet")
	return this
}

func (this *Part) ShowStatistics() *Part {
	this.Append("statistics")
	return this
}

func (this *Part) ShowStatus() *Part {
	this.Append("status")
	return this
}

func (this *Part) ShowTopicDetails() *Part {
	this.Append("topicDetails")
	return this
}

func (this *Part) Append(v string) *Part {
	if this.value == nil {
		this.value = make([]string, 0, 0)
	}
	if v != "" {
		this.value = append(this.value, v)
	}
	return this
}

func (this *Part) Values() string {
	if len(this.value) == 0 {
		this.ShowSnippet()
	}
	return strings.Join(this.value, ",")
}

type PageInfo struct {
	TotalResults   int64 `json:"totalResults"`
	ResultsPerPage int64 `json:"resultsPerPage"`
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

type RelatedPlayLists struct {
	Favorites    string `json:"favorites"`
	Uploads      string `json:"uploads"`
	WatchHistory string `json:"watchHistory"`
	WatchLater   string `json:"watchLater"`
}
