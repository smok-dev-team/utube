package utube

import (
	"fmt"
	"net/url"
	"strings"
)

// https://developers.google.cn/youtube/v3/docs/commentThreads/list
type GetCommentThreadsParam struct {
	// Required parameters
	Part Part

	// Filters (specify exactly one of the following parameters)
	AllThreadsRelatedToChannelId string
	ChannelId                    string
	Id                           string
	VideoId                      string

	// Optional parameters
	MaxResults       int
	PageToken        string
	ModerationStatus string
	Order            string
	SearchTerms      string
	TextFormat       string
}

func (this GetCommentThreadsParam) Params() url.Values {
	var v = url.Values{}

	v.Add("part", this.Part.Values())

	if len(this.AllThreadsRelatedToChannelId) > 0 {
		v.Add("allThreadsRelatedToChannelId", this.AllThreadsRelatedToChannelId)
	}

	if len(this.ChannelId) > 0 {
		v.Add("channelId", this.ChannelId)
	}

	if len(this.Id) > 0 {
		v.Add("id", this.Id)
	}

	if len(this.VideoId) > 0 {
		v.Add("videoId", this.VideoId)
	}

	if this.MaxResults > 0 {
		v.Add("maxResults", fmt.Sprintf("%d", this.MaxResults))
	}

	if len(this.PageToken) > 0 {
		v.Add("pageToken", this.PageToken)
	}

	if len(this.ModerationStatus) > 0 {
		v.Add("moderationStatus", this.ModerationStatus)
	}

	if len(this.Order) > 0 {
		v.Add("order", this.Order)
	}

	if len(this.SearchTerms) > 0 {
		v.Add("searchTerms", this.SearchTerms)
	}

	if len(this.TextFormat) > 0 {
		v.Add("textFormat", this.TextFormat)
	}
	return v
}

// https://developers.google.cn/youtube/v3/docs/comments/list
type GetCommentsParams struct {
	// Required parameters
	Part Part

	// Filters (specify exactly one of the following parameters)
	Ids      []string
	ParentId string

	// Optional parameters
	MaxResults int
	PageToken  string
	TextFormat string
}

func (this *GetCommentsParams) AppendId(id string) {
	this.Ids = append(this.Ids, id)
}

func (this GetCommentsParams) Params() url.Values {
	var v = url.Values{}

	v.Add("part", this.Part.Values())

	if len(this.ParentId) > 0 {
		v.Add("parentId", this.ParentId)
	}

	if len(this.Ids) > 0 {
		v.Add("id", strings.Join(this.Ids, ","))
	}

	if this.MaxResults > 0 {
		v.Add("maxResults", fmt.Sprintf("%d", this.MaxResults))
	}

	if len(this.PageToken) > 0 {
		v.Add("pageToken", this.PageToken)
	}

	if len(this.TextFormat) > 0 {
		v.Add("textFormat", this.TextFormat)
	}
	return v
}

type CommentThreads struct {
	Kind          string           `json:"kind"`
	ETag          string           `json:"etag"`
	NextPageToken string           `json:"nextPageToken"`
	PrevPageToken string           `json:"prevPageToken"`
	PageInfo      *PageInfo        `json:"pageInfo"`
	Items         []*CommentThread `json:"items,omitempty"`
}

type CommentThread struct {
	Kind    string                `json:"kind"`
	ETag    string                `json:"etag"`
	Id      string                `json:"id"`
	Snippet *CommentThreadSnippet `json:"snippet,omitempty"`
	Replies *CommentThreadReplies `json:"replies,omitempty"`
}

type CommentThreadSnippet struct {
	VideoId         string   `json:"videoId"`
	CanReply        bool     `json:"canReply"`
	TotalReplyCount int      `json:"totalReplyCount"`
	IsPublic        bool     `json:"isPublic"`
	TopLevelComment *Comment `json:"topLevelComment,omitempty"`
}

type Comments struct {
	Kind  string     `json:"kind"`
	ETag  string     `json:"etag"`
	Id    string     `json:"id"`
	Items []*Comment `json:"items"`
}

type Comment struct {
	Kind    string          `json:"kind"`
	ETag    string          `json:"etag"`
	Id      string          `json:"id"`
	Snippet *CommentSnippet `json:"snippet,omitempty"`
}

type CommentSnippet struct {
	AuthorDisplayName     string           `json:"authorDisplayName"`
	AuthorProfileImageUrl string           `json:"authorProfileImageUrl"`
	AuthorChannelUrl      string           `json:"authorChannelUrl"`
	AuthorChannelId       *AuthorChannelId `json:"authorChannelId"`
	VideoId               string           `json:"videoId"`
	TextDisplay           string           `json:"textDisplay"`
	TextOriginal          string           `json:"textOriginal"`
	CanRate               bool             `json:"canRate"`
	ViewerRating          string           `json:"viewerRating"`
	LikeCount             int              `json:"likeCount"`
	PublishedAt           string           `json:"publishedAt"`
	UpdatedAt             string           `json:"updatedAt"`
}

type AuthorChannelId struct {
	Value string `json:"value"`
}

type CommentThreadReplies struct {
	Comments []*Comment `json:"comments,omitempty"`
}
