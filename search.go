package utube

const (
	k_SEARCH_API = "/v3/search"
)

// Search https://developers.google.cn/youtube/v3/docs/search/list
func (this *Youtube) Search(param SearchParam) (results *SearchResults, err error) {
	var api = this.BuildAPI(k_SEARCH_API)
	err = this.doRequest("GET", api, param, &results)
	return results, err
}

func (this *Youtube) SearchVideo(param SearchParam) (results *SearchResults, err error) {
	param.Type = "video"

	var api = this.BuildAPI(k_SEARCH_API)
	err = this.doRequest("GET", api, param, &results)
	return results, err
}

func (this *Youtube) SearchChannel(param SearchParam) (results *SearchResults, err error) {
	param.Type = "channel"

	var api = this.BuildAPI(k_SEARCH_API)
	err = this.doRequest("GET", api, param, &results)
	return results, err
}

func (this *Youtube) SearchPlaylist(param SearchParam) (results *SearchResults, err error) {
	param.Type = "playlist"

	var api = this.BuildAPI(k_SEARCH_API)
	err = this.doRequest("GET", api, param, &results)
	return results, err
}