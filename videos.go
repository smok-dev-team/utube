package utube

const (
	k_VIDEO_API        = "/v3/videos"
	k_VIDEO_CATEGORIES = "/v3/videoCategories"
)

// GetVideos https://developers.google.cn/youtube/v3/docs/videos/list
func (this *Youtube) GetVideos(accessToken string, param GetVideosParam) (results *Videos, err error) {
	var api = this.BuildAPI(k_VIDEO_API)
	err = this.doRequest("GET", api, accessToken, param, &results)
	return results, err
}

// GetVideoCategories https://developers.google.cn/youtube/v3/docs/videoCategories/list
func (this *Youtube) GetVideoCategories(accessToken string, param GetVideoCategoriesParam) (results *VideoCategories, err error) {
	var api = this.BuildAPI(k_VIDEO_CATEGORIES)
	err = this.doRequest("GET", api, accessToken, param, &results)
	return results, err
}
