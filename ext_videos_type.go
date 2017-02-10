package utube

import (
	"github.com/smartwalle/form"
	"net/url"
	"strings"
)

type VideoInfo struct {
	VideoId             string             `form:"video_id"`
	EventId             string             `form:"eventid"`
	Keywords            string             `form:"keywords"`
	Author              string             `form:"author"`
	Title               string             `form:"title"`
	ViewCount           string             `form:"view_count"`
	IURL                string             `form:"iurl"`          // 视频预览图片
	IURLMQ              string             `form:"iurlmq"`        // 视频预览图片 - 高等质量
	IURLHQ              string             `form:"iurlhq"`        // 视频预览图片 - 中等质量
	ThumbnailURL        string             `form:"thumbnail_url"` // 视频预览图片 - 缩略图
	LengthSeconds       string             `form:"length_seconds"`
	VideoStreamInfoList []*VideoStreamInfo `form:"url_encoded_fmt_stream_map"`

	HL                          string `form:"hl"`
	CaptionAudioTracks          string `form:"caption_audio_tracks"`
	CsiPageType                 string `form:"csi_page_type"`
	VM                          string `form:"vm"`
	AvgRating                   string `form:"avg_rating"`
	Fexp                        string `form:"fexp"`
	Enablecsi                   string `form:"enablecsi"`
	Token                       string `form:"token"`
	DefaultAudioTrackIndex      string `form:"default_audio_track_index"`
	CCFontsURL                  string `form:"cc_fonts_url"`
	CCModule                    string `form:"cc_module"`
	CaptionTracks               string `form:"caption_tracks"`
	Oid                         string `form:"oid"`
	Pltype                      string `form:"pltype"`
	HasCC                       string `form:"has_cc"`
	CaptionTranslationLanguages string `form:"caption_translation_languages"`
	Ptk                         string `form:"ptk"`
	Timestamp                   string `form:"timestamp"`
	FmtList                     string `form:"fmt_list"`
	Cver                        string `form:"cver"`
	Plid                        string `form:"plid"`
	Ptchn                       string `form:"ptchn"`
	UseCipherSignature          string `form:"use_cipher_signature"`
	Dashmpd                     string `form:"dashmpd"`
	StoryboardSpec              string `form:"storyboard_spec"`
	Ldpj                        string `form:"ldpj"`
	AdaptiveFmts                string `form:"adaptive_fmts"`
	AccountPlaybackToken        string `form:"account_playback_token"`
	VideoVerticals              string `form:"video_verticals"`
	Loudness                    string `form:"loudness"`
	NoGetVideoLog               string `form:"no_get_video_log"`
	AllowedAds                  string `form:"allowed_ads"`
	Tmi                         string `form:"tmi"`
	Ucid                        string `form:"ucid"`
	Of                          string `form:"of"`
	AllowEmbed                  string `form:"allow_embed"`
	CCAsr                       string `form:"cc_asr"`
	YpcAdIndicator              string `form:"ypc_ad_indicator"`
	Idpj                        string `form:"idpj"`
	CC3Module                   string `form:"cc3_module"`
	AllowRatings                string `form:"allow_ratings"`
	IsListed                    string `form:"is_listed"`
	Cos                         string `form:"cos"`
	TTSURL                      string `form:"ttsurl"`
	C                           string `form:"c"`
	Watermark                   string `form:"watermark"`
	CL                          string `form:"cl"`
	Status                      string `form:"status"`
	ProbeURL                    string `form:"probe_url"`
	Muted                       string `form:"muted"`
	CCFont                      string `form:"cc_font"`
}

func (this VideoInfo) CleanedVideoStreamInfoList(p string) (infoList []*VideoStreamInfo) {
	var ps = strings.Split(p, ",")

	infoList = make([]*VideoStreamInfo, 0, 0)

	for i := 0; i < len(ps); i++ {
		param, err := url.ParseQuery(ps[i])
		if err != nil {
			continue
		}

		var info *VideoStreamInfo
		err = form.Bind(param, &info)
		if err != nil {
			continue
		}
		infoList = append(infoList, info)
	}

	return infoList
}

type VideoStreamInfo struct {
	Type         string `form:"type"`
	Quality      string `form:"quality"`
	URL          string `form:"url"`
	ITag         string `form:"itag"`
	FallbackHost string `form:"fallback_host"`
}
