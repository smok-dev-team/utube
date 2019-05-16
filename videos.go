package utube

const (
	k_VIDEO_API        = "/v3/videos"
	k_VIDEO_CATEGORIES = "/v3/videoCategories"
)

// GetVideos https://developers.google.cn/youtube/v3/docs/videos/list
func (this *Client) GetVideos(param GetVideosParam) (results *Videos, err error) {
	var api = this.BuildAPI(k_VIDEO_API)
	err = this.doRequest("GET", api, param, &results)
	return results, err
}

// GetVideoCategories https://developers.google.cn/youtube/v3/docs/videoCategories/list
func (this *Client) GetVideoCategories(param GetVideoCategoriesParam) (results *VideoCategories, err error) {
	var api = this.BuildAPI(k_VIDEO_CATEGORIES)
	err = this.doRequest("GET", api, param, &results)
	return results, err
}
