package utube

import (
	"github.com/smartwalle/form"
	"github.com/smartwalle/going/request"
	"net/url"
)

const (
	k_YOUTUBE_GET_VIDEO_INFO_URL = "https://www.youtube.com/get_video_info"
)

func GetVideoInfoWithVideoId(videoId string) (metadata *VideoInfo, err error) {
	result, err := request.Request("GET", k_YOUTUBE_GET_VIDEO_INFO_URL, url.Values{"video_id": []string{videoId}})

	if err != nil {
		return nil, err
	}

	var queryStr = string(result)
	return VideoInfoWithQueryString(queryStr)
}

func VideoInfoWithQueryString(query string) (metadata *VideoInfo, err error) {
	param, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}

	err = form.Bind(param, &metadata)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}
