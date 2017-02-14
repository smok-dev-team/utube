package utube

import (
	"fmt"
	"net/url"
)

// GetChannelsParam https://developers.google.cn/youtube/v3/docs/channels/list
type GetChannelsParam struct {
	// Required parameters
	Part Part

	// Filters (specify exactly one of the following parameters)
	CategoryId    string
	ForUsername   string
	Id            string
	ManagedByMe   int // authorized request 1 true, 2 false
	Mine          int // authorized request 1 true, 2 false
	MySubscribers int // authorized request 1 true, 2 false

	// Optional parameters
	HL                     string
	OnBehalfOfContentOwner string // authorized request

	MaxResults int
	PageToken  string
}

func (this GetChannelsParam) Params() url.Values {
	var v = url.Values{}

	v.Add("part", this.Part.Values())

	if len(this.CategoryId) > 0 {
		v.Add("categoryId", this.CategoryId)
	}

	if len(this.ForUsername) > 0 {
		v.Add("forUsername", this.ForUsername)
	}

	if len(this.Id) > 0 {
		v.Add("id", this.Id)
	}

	if this.ManagedByMe == 1 {
		v.Add("managedByMe", "true")
	} else if this.ManagedByMe == 2 {
		v.Add("managedByMe", "false")
	}

	if this.Mine == 1 {
		v.Add("mine", "true")
	} else if (this.Mine) == 2 {
		v.Add("mine", "false")
	}

	if this.MySubscribers == 1 {
		v.Add("mySubscribers", "true")
	} else if this.MySubscribers == 2 {
		v.Add("mySubscribers", "false")
	}

	if len(this.HL) > 0 {
		v.Add("hl", this.HL)
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

type Channels struct {
	Kind          string     `json:"kind"`
	ETag          string     `json:"etag"`
	NextPageToken string     `json:"nextPageToken"`
	PrevPageToken string     `json:"prevPageToken"`
	RegionCode    string     `json:"regionCode"`
	PageInfo      *PageInfo  `json:"pageInfo"`
	Items         []*Channel `json:"items,omitempty"`
}

type Channel struct {
	Kind             string                    `json:"kind"`
	ETag             string                    `json:"etag"`
	Id               string                    `json:"id"`
	Snippet          *ChannelSnippet           `json:"snippet,omitempty"`
	ContentDetails   *ChannelContentDetails    `json:"contentDetails,omitempty"`
	Statistics       *ChannelContentStatistics `json:"statistics,omitempty"`
	TopicDetails     *TopicDetails             `json:"topicDetails,omitempty"`
	Status           *ChannelStatus            `json:"status,omitempty"`
	BrandingSettings *ChannelBrandingSettings  `json:"brandingSettings,omitempty"`
}

type ChannelSnippet struct {
	Title           string      `json:"title"`
	Description     string      `json:"description"`
	CustomUrl       string      `json:"customUrl"`
	PublishedAt     string      `json:"publishedAt"`
	Thumbnails      *Thumbnails `json:"thumbnails,omitempty"`
	DefaultLanguage string      `json:"defaultLanguage"`
	Localized       *Localized  `json:"localized,omitempty"`
	Country         string      `json:"country"`
}

type ChannelContentDetails struct {
	RelatedPlayLists *RelatedPlayLists `json:"relatedPlaylists,omitempty"`
}

type ChannelContentStatistics struct {
	ViewCount             string `json:"viewCount"`
	CommentCount          string `json:"commentCount"`
	SubscriberCount       string `json:"subscriberCount"`
	HiddenSubscriberCount bool   `json:"hiddenSubscriberCount"`
	VideoCount            string `json:"videoCount"`
}

type ChannelStatus struct {
	PrivacyStatus     string `json:"privacyStatus"`
	IsLinked          bool   `json:"isLinked"`
	LongUploadsStatus string `json:"longUploadsStatus"`
}

type ChannelBrandingSettings struct {
	Channel *ChannelBrandingChannel `json:"channel,omitempty"`
	Image   *ChannelBrandingImage   `json:"image,omitempty"`
	Hints   []*ChannelHint          `json:"hints,omitempty"`
}

type ChannelBrandingChannel struct {
	Title               string `json:"title"`
	Description         string `json:"description"`
	ShowRelatedChannels bool   `json:"showRelatedChannels"`
	ShowBrowseView      bool   `json:"showBrowseView"`
	UnsubscribedTrailer string `json:"unsubscribedTrailer"`
	ProfileColor        string `json:"profileColor"`
}

type ChannelBrandingImage struct {
	BannerImageUrl               string `json:"bannerImageUrl"`
	BannerMobileImageUrl         string `json:"bannerMobileImageUrl"`
	BannerTabletLowImageUrl      string `json:"bannerTabletLowImageUrl"`
	BannerTabletImageUrl         string `json:"bannerTabletImageUrl"`
	BannerTabletHdImageUrl       string `json:"bannerTabletHdImageUrl"`
	BannerTabletExtraHdImageUrl  string `json:"bannerTabletExtraHdImageUrl"`
	BannerMobileLowImageUrl      string `json:"bannerMobileLowImageUrl"`
	BannerMobileMediumHdImageUrl string `json:"bannerMobileMediumHdImageUrl"`
	BannerMobileHdImageUrl       string `json:"bannerMobileHdImageUrl"`
	BannerMobileExtraHdImageUrl  string `json:"bannerMobileExtraHdImageUrl"`
	BannerTvImageUrl             string `json:"bannerTvImageUrl"`
	BannerTvLowImageUrl          string `json:"bannerTvLowImageUrl"`
	BannerTvMediumImageUrl       string `json:"bannerTvMediumImageUrl"`
	BannerTvHighImageUrl         string `json:"bannerTvHighImageUrl"`
}

type ChannelHint struct {
	Property string `json:"property"`
	Value    string `json:"value"`
}
