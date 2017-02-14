package utube

import (
	"fmt"
	"net/url"
)

// GetVideosParam https://developers.google.cn/youtube/v3/docs/videos/list
type GetVideosParam struct {
	// Required parameters
	Part Part

	// Filters (specify exactly one of the following parameters)
	Chart    string
	Id       string
	MyRating string // authorized request

	// Optional parameters
	HL                     string
	MaxHeight              int
	MaxWidth               int
	RegionCode             string
	VideoCategoryId        string
	OnBehalfOfContentOwner string

	MaxResults int
	PageToken  string
}

func (this GetVideosParam) Params() url.Values {
	var v = url.Values{}

	v.Add("part", this.Part.Values())

	if len(this.Chart) > 0 {
		v.Add("chart", this.Chart)
	}

	if len(this.Id) > 0 {
		v.Add("id", this.Id)
	}

	if len(this.MyRating) > 0 {
		v.Add("myRating", this.MyRating)
	}

	if len(this.HL) > 0 {
		v.Add("hl", this.HL)
	}

	if this.MaxHeight > 0 {
		v.Add("maxHeight", fmt.Sprintf("%d", this.MaxHeight))
	}

	if this.MaxWidth > 0 {
		v.Add("maxWidth", fmt.Sprintf("%d", this.MaxWidth))
	}

	if len(this.RegionCode) > 0 {
		v.Add("regionCode", this.RegionCode)
	}

	if len(this.VideoCategoryId) > 0 {
		v.Add("videoCategoryId", this.VideoCategoryId)
	}

	if len(this.OnBehalfOfContentOwner) > 0 {
		v.Add("onBehalfOfContentOwner", this.OnBehalfOfContentOwner)
	}

	if this.MaxResults > 0 {
		v.Add("maxResults", fmt.Sprintf("%d", this.MaxResults))
	}

	if len(this.PageToken) > 0 {
		v.Add("pageToken", this.PageToken)
	}

	return v
}

type Videos struct {
	Kind          string    `json:"kind"`
	ETag          string    `json:"etag"`
	NextPageToken string    `json:"nextPageToken"`
	PrevPageToken string    `json:"prevPageToken"`
	PageInfo      *PageInfo `json:"pageInfo"`
	Items         []*Video  `json:"items,omitempty"`
}

type Video struct {
	Kind           string               `json:"kind"`
	ETag           string               `json:"etag"`
	Id             string               `json:"id"`
	Snippet        *VideoSnippet        `json:"snippet,omitempty"`
	ContentDetails *VideoContentDetails `json:"contentDetails,omitempty"`
	Status         *VideoStatus         `json:"status,omitempty"`
	Statistics     *VideoStatistics     `json:"statistics"`
	Player         *Player              `json:"player"`
	TopicDetails   *TopicDetails        `json:"topicDetails,omitempty"`
}

type VideoSnippet struct {
	Title                string      `json:"title"`
	Description          string      `json:"description"`
	ChannelId            string      `json:"channelId"`
	ChannelTitle         string      `json:"channelTitle"`
	CategoryId           string      `json:"categoryId"`
	PublishedAt          string      `json:"publishedAt"`
	Thumbnails           *Thumbnails `json:"thumbnails,omitempty"`
	Tags                 []string    `json:"tags"`
	LiveBroadcastContent string      `json:"liveBroadcastContent"`
	Localized            *Localized  `json:"localized,omitempty"`
	DefaultAudioLanguage string      `json:"defaultAudioLanguage"`
}

type VideoContentDetails struct {
	Duration          string             `json:"duration"`
	Dimension         string             `json:"dimension"`
	Definition        string             `json:"definition"`
	Caption           string             `json:"caption"`
	LicensedContent   bool               `json:"licensedContent"`
	RegionRestriction *RegionRestriction `json:"regionRestriction,omitempty"`
	Projection        string             `json:"projection"`
}

type RegionRestriction struct {
	Blocked []string `json:"blocked"`
}

type VideoStatus struct {
	UploadStatus        string `json:"uploadStatus"`
	PrivacyStatus       string `json:"privacyStatus"`
	License             string `json:"license"`
	Embeddable          bool   `json:"embeddable"`
	PublicStatsViewable bool   `json:"publicStatsViewable"`
}

type VideoStatistics struct {
	ViewCount     string `json:"viewCount"`
	LikeCount     string `json:"likeCount"`
	DislikeCount  string `json:"dislikeCount"`
	FavoriteCount string `json:"favoriteCount"`
	CommentCount  string `json:"commentCount"`
}

// GetVideoCategoriesParam https://developers.google.cn/youtube/v3/docs/videoCategories/list
type GetVideoCategoriesParam struct {
	// Required parameters
	Part Part

	// Filters (specify exactly one of the following parameters)
	Id         string
	RegionCode string

	// Optional parameters
	HL string
}

func (this GetVideoCategoriesParam) Params() url.Values {
	var v = url.Values{}

	v.Add("part", this.Part.Values())

	if len(this.Id) > 0 {
		v.Add("id", this.Id)
	}

	if len(this.RegionCode) > 0 {
		v.Add("regionCode", this.RegionCode)
	}

	if len(this.HL) > 0 {
		v.Add("hl", this.HL)
	}

	return v
}

type VideoCategories struct {
	Kind          string           `json:"kind"`
	ETag          string           `json:"etag"`
	NextPageToken string           `json:"nextPageToken"`
	PrevPageToken string           `json:"prevPageToken"`
	PageInfo      *PageInfo        `json:"pageInfo"`
	Items         []*VideoCategory `json:"items,omitempty"`
}

type VideoCategory struct {
	Kind    string                `json:"kind"`
	ETag    string                `json:"etag"`
	Id      string                `json:"id"`
	Snippet *VideoCategorySnippet `json:"snippet,omitempty"`
}

type VideoCategorySnippet struct {
	Title      string `json:"title"`
	ChannelId  string `json:"channelId"`
	Assignable bool   `json:"assignable"`
}
