package utube

const (
	k_VIDEO_API = "/v3/videos"
)

// GetVideos https://developers.google.cn/youtube/v3/docs/videos/list
func (this *Youtube) GetVideos(param GetVideosParam) (results *Videos, err error) {
	var api = this.BuildAPI(k_VIDEO_API)
	err = this.doRequest("GET", api, param, &results)
	return results, err
}