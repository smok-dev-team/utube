package utube

import (
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

func (this *Part) ShowReplies() *Part {
	this.Append("replies")
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

type Localization struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Localizations struct {
	HU     *Localization `json:"hu,omitempty"`
	TH     *Localization `json:"th,omitempty"`
	SR     *Localization `json:"sr,omitempty"`
	ZU     *Localization `json:"zu,omitempty"`
	PT_PT  *Localization `json:"pt-PT,omitempty"`
	SW     *Localization `json:"sw,omitempty"`
	TA     *Localization `json:"ta,omitempty"`
	MK     *Localization `json:"mk,omitempty"`
	KY     *Localization `json:"ky,omitempty"`
	IW     *Localization `json:"iw,omitempty"`
	PL     *Localization `json:"pl,omitempty"`
	SR_XA  *Localization `json:"sr-XA,omitempty"`
	PT     *Localization `json:"pt,omitempty"`
	ZH_TW  *Localization `json:"zh-TW,omitempty"`
	ES_419 *Localization `json:"es-419,omitempty"`
	NL     *Localization `json:"nl,omitempty"`
	ID     *Localization `json:"id,omitempty"`
	SK     *Localization `json:"sk,omitempty"`
	SL     *Localization `json:"sl,omitempty"`
	HI     *Localization `json:"hi,omitempty"`
	RU     *Localization `json:"ru,omitempty"`
	MR     *Localization `json:"mr,omitempty"`
	EL     *Localization `json:"el,omitempty"`
	PA     *Localization `json:"pa,omitempty"`
	LT     *Localization `json:"lt,omitempty"`
	AR     *Localization `json:"ar,omitempty"`
	FI     *Localization `json:"fi,omitempty"`
	NE     *Localization `json:"ne,omitempty"`
	Ro     *Localization `json:"ro,omitempty"`
	KO     *Localization `json:"ko,omitempty"`
	ML     *Localization `json:"ml,omitempty"`
	BN     *Localization `json:"bn,omitempty"`
	SV     *Localization `json:"sv,omitempty"`
	UR     *Localization `json:"ur,omitempty"`
	FIL    *Localization `json:"fil,omitempty"`
	KK     *Localization `json:"kk,omitempty"`
	SQ     *Localization `json:"sq,omitempty"`
	LO     *Localization `json:"lo,omitempty"`
	HY     *Localization `json:"hy,omitempty"`
	CA     *Localization `json:"ca,omitempty"`
	MS     *Localization `json:"ms,omitempty"`
	BS     *Localization `json:"bs,omitempty"`
	UK     *Localization `json:"uk,omitempty"`
	MY     *Localization `json:"my,omitempty"`
	KA     *Localization `json:"ka,omitempty"`
	ET     *Localization `json:"et,omitempty"`
	EN     *Localization `json:"en,omitempty"`
	LV     *Localization `json:"lv,omitempty"`
	AZ     *Localization `json:"az,omitempty"`
	AM     *Localization `json:"am,omitempty"`
	DA     *Localization `json:"da,omitempty"`
	ZH_HK  *Localization `json:"zh-HK,omitempty"`
	GL     *Localization `json:"gl,omitempty"`
	DE     *Localization `json:"de,omitempty"`
	VI     *Localization `json:"vi,omitempty"`
	UZ     *Localization `json:"uz,omitempty"`
	FR_CA  *Localization `json:"fr-CA,omitempty"`
	ZH_CN  *Localization `json:"zh-CN,omitempty"`
	FA     *Localization `json:"fa,omitempty"`
	TE     *Localization `json:"te,omitempty"`
	CS     *Localization `json:"cs,omitempty"`
	GU     *Localization `json:"gu,omitempty"`
	EU     *Localization `json:"eu,omitempty"`
	NO     *Localization `json:"no,omitempty"`
	AF     *Localization `json:"af,omitempty"`
	KM     *Localization `json:"km,omitempty"`
	FR     *Localization `json:"fr,omitempty"`
	TR     *Localization `json:"tr,omitempty"`
	BG     *Localization `json:"bg,omitempty"`
	HR     *Localization `json:"hr,omitempty"`
	IS     *Localization `json:"is,omitempty"`
	JA     *Localization `json:"ja,omitempty"`
	MN     *Localization `json:"mn,omitempty"`
	KN     *Localization `json:"kn,omitempty"`
	EN_GB  *Localization `json:"en-GB,omitempty"`
	SI     *Localization `json:"si,omitempty"`
	IT     *Localization `json:"it,omitempty"`
}

type Player struct {
	EmbedHtml string `json:"embedHtml"`
}

type TopicDetails struct {
	TopicIds         []string `json:"topicIds,omitempty"`
	RelevantTopicIds []string `json:"relevantTopicIds,omitempty"`
}