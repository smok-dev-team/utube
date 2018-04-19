package utube

import (
	"fmt"
	"net/url"
)

type SearchParam struct {
	// Required parameters
	Part Part

	// Filters (specify 0 or 1 of the following parameters)
	ForContentOwner  int // authorized request
	ForDeveloper     int // authorized request
	ForMine          int // authorized request
	RelatedToVideoId string

	// Optional parameters
	ChannelId              string
	ChannelType            string
	EventType              string
	Location               string
	LocationRadius         string
	Type                   string
	Order                  string
	OnBehalfOfContentOwner string // authorized request
	PublishedAfter         string
	PublishedBefore        string
	Q                      string
	RegionCode             string
	RelevanceLanguage      string
	SafeSearch             string
	TopicId                string
	VideoCaption           string
	VideoCategoryId        string
	VideoDefinition        string
	VideoDimension         string
	VideoDuration          string
	VideoEmbeddable        string
	VideoLicense           string
	VideoSyndicated        string
	VideoType              string

	MaxResults int
	PageToken  string
}

func (this SearchParam) Params() url.Values {
	var v = url.Values{}

	v.Add("part", this.Part.Values())

	if this.ForContentOwner == 1 {
		v.Add("forContentOwner", "true")
	} else if this.ForContentOwner == 2 {
		v.Add("forContentOwner", "false")
	}

	if this.ForDeveloper == 1 {
		v.Add("forDeveloper", "true")
	} else if this.ForDeveloper == 2 {
		v.Add("forDeveloper", "false")
	}

	if this.ForMine == 1 {
		v.Add("forMine", "true")
	} else if this.ForMine == 2 {
		v.Add("forMine", "false")
	}

	if len(this.RelatedToVideoId) > 0 {
		v.Add("relatedToVideoId", this.RelatedToVideoId)
	}

	if len(this.ChannelId) > 0 {
		v.Add("channelId", this.ChannelId)
	}

	if len(this.ChannelType) > 0 {
		v.Add("channelType", this.ChannelType)
	}

	if len(this.EventType) > 0 {
		v.Add("eventType", this.EventType)
	}

	if len(this.Location) > 0 {
		v.Add("location", this.Location)
	}

	if len(this.LocationRadius) > 0 {
		v.Add("locationRadius", this.LocationRadius)
	}

	if len(this.Order) == 0 {
		this.Order = "date"
	}
	v.Add("order", this.Order)

	if len(this.Type) == 0 {
		this.Type = "video,channel,playlist"
	}
	v.Add("type", this.Type)

	if len(this.OnBehalfOfContentOwner) > 0 {
		v.Add("onBehalfOfContentOwner", this.OnBehalfOfContentOwner)
	}

	if len(this.PublishedAfter) > 0 {
		v.Add("publishedAfter", this.PublishedAfter)
	}

	if len(this.PublishedBefore) > 0 {
		v.Add("publishedBefore", this.PublishedBefore)
	}

	if len(this.Q) > 0 {
		v.Add("q", this.Q)
	}

	if len(this.RegionCode) > 0 {
		v.Add("regionCode", this.RegionCode)
	}

	if len(this.RelevanceLanguage) > 0 {
		v.Add("relevanceLanguage", this.RelevanceLanguage)
	}

	if len(this.SafeSearch) > 0 {
		v.Add("safeSearch", this.SafeSearch)
	}

	if len(this.TopicId) > 0 {
		v.Add("topicId", this.TopicId)
	}

	if len(this.VideoCaption) > 0 {
		v.Add("videoCaption", this.VideoCaption)
	}

	if len(this.VideoCategoryId) > 0 {
		v.Add("videoCategoryId", this.VideoCategoryId)
	}

	if len(this.VideoDefinition) > 0 {
		v.Add("videoDefinition", this.VideoDefinition)
	}

	if len(this.VideoDimension) > 0 {
		v.Add("videoDimension", this.VideoDimension)
	}

	if len(this.VideoDuration) > 0 {
		v.Add("videoDuration", this.VideoDuration)
	}

	if len(this.VideoEmbeddable) > 0 {
		v.Add("videoEmbeddable", this.VideoEmbeddable)
	}

	if len(this.VideoLicense) > 0 {
		v.Add("videoLicense", this.VideoLicense)
	}

	if len(this.VideoSyndicated) > 0 {
		v.Add("videoSyndicated", this.VideoSyndicated)
	}

	if len(this.VideoType) > 0 {
		v.Add("videoType", this.VideoType)
	}

	if this.MaxResults > 0 {
		v.Add("maxResults", fmt.Sprintf("%d", this.MaxResults))
	}

	if len(this.PageToken) > 0 {
		v.Add("pageToken", this.PageToken)
	}
	return v
}

type SearchResults struct {
	Kind          string          `json:"kind"`
	ETag          string          `json:"etag"`
	NextPageToken string          `json:"nextPageToken"`
	PrevPageToken string          `json:"prevPageToken"`
	RegionCode    string          `json:"regionCode"`
	PageInfo      *PageInfo       `json:"pageInfo"`
	Items         []*SearchResult `json:"items,omitempty"`
}

type SearchResult struct {
	Kind    string         `json:"kind"`
	ETag    string         `json:"etag"`
	Id      *ResourceId    `json:"id"`
	Snippet *SearchSnippet `json:"snippet,omitempty"`
}

func (this *SearchResult) IsVideo() bool {
	return this.Id.IsVideo()
}

func (this *SearchResult) IsChannel() bool {
	return this.Id.IsChannel()
}

func (this *SearchResult) IsPlaylist() bool {
	return this.Id.IsPlaylist()
}

func (this *SearchResult) GetId() string {
	return this.Id.GetId()
}

type SearchSnippet struct {
	Title                string      `json:"title"`
	Description          string      `json:"description"`
	ChannelId            string      `json:"channelId"`
	ChannelTitle         string      `json:"channelTitle"`
	PublishedAt          string      `json:"publishedAt"`
	LiveBroadcastContent string      `json:"liveBroadcastContent"`
	Thumbnails           *Thumbnails `json:"thumbnails"`
}
