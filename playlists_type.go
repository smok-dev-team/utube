package utube

import (
	"fmt"
	"net/url"
)

// GetPlaylistsParam https://developers.google.cn/youtube/v3/docs/playlists/list
type GetPlaylistsParam struct {
	// Required parameters
	Part string

	// Filters (specify exactly one of the following parameters)
	ChannelId string
	Id        string
	Mine      int // authorized request 1 true, 2 false

	// Optional parameters
	HL                            string
	OnBehalfOfContentOwner        string // authorized request
	OnBehalfOfContentOwnerChannel string // authorized request

	MaxResults int
	PageToken  string
}

func (this GetPlaylistsParam) Params() url.Values {
	var v = url.Values{}

	if len(this.Part) == 0 {
		this.Part = "snippet"
	}
	v.Add("part", this.Part)

	if len(this.ChannelId) > 0 {
		v.Add("channelId", this.ChannelId)
	}

	if len(this.Id) > 0 {
		v.Add("id", this.Id)
	}

	if this.Mine == 1 {
		v.Add("mine", "true")
	} else if (this.Mine) == 2 {
		v.Add("mine", "false")
	}

	if len(this.HL) > 0 {
		v.Add("hl", this.HL)
	}

	if len(this.OnBehalfOfContentOwner) > 0 {
		v.Add("onBehalfOfContentOwner", this.OnBehalfOfContentOwner)
	}

	if len(this.OnBehalfOfContentOwnerChannel) > 0 {
		v.Add("onBehalfOfContentOwnerChannel", this.OnBehalfOfContentOwnerChannel)
	}

	if this.MaxResults > 0 {
		v.Add("maxResults", fmt.Sprintf("%d", this.MaxResults))
	}

	if len(this.PageToken) > 0 {
		v.Add("pageToken", this.PageToken)
	}

	return v
}

type Playlists struct {
	Kind          string      `json:"kind"`
	ETag          string      `json:"etag"`
	NextPageToken string      `json:"nextPageToken"`
	PrevPageToken string      `json:"prevPageToken"`
	RegionCode    string      `json:"regionCode"`
	PageInfo      *PageInfo   `json:"pageInfo"`
	Items         []*Playlist `json:"items,omitempty"`
}

type Playlist struct {
	Kind    string           `json:"kind"`
	ETag    string           `json:"etag"`
	Id      string           `json:"id"`
	Snippet *PlaylistSnippet `json:"snippet,omitempty"`
}

type PlaylistSnippet struct {
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	ChannelId    string      `json:"channelId"`
	ChannelTitle string      `json:"channelTitle"`
	PublishedAt  string      `json:"publishedAt"`
	Thumbnails   *Thumbnails `json:"thumbnails"`
	Localized    *Localized  `json:"localized"`
}
