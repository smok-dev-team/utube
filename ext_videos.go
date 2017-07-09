package utube

import (
	"github.com/smartwalle/form"
	"github.com/smartwalle/going/request"
	"net/url"
)

const (
	k_YOUTUBE_GET_VIDEO_INFO_URL = "https://www.youtube.com/get_video_info"
)

func GetVideoInfoWithVideoId(videoId string) (videoInfo *VideoInfo, err error) {
	var req = request.NewRequest("GET", k_YOUTUBE_GET_VIDEO_INFO_URL)
	req.SetParam("video_id", videoId)

	//result, err := request.Request("GET", k_YOUTUBE_GET_VIDEO_INFO_URL, url.Values{"video_id": []string{videoId}})
	//if err != nil {
	//	return nil, err
	//}
	var rep = req.Exec()

	var queryStr = rep.MustString()
	videoInfo, err = VideoInfoWithQueryString(queryStr)
	return videoInfo, err
}

func VideoInfoWithQueryString(query string) (videoInfo *VideoInfo, err error) {
	param, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}

	err = form.Bind(param, &videoInfo)
	if err != nil {
		return nil, err
	}
	return videoInfo, nil
}
