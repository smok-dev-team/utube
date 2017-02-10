package utube

const (
	k_PLAYLISTS_API = "/v3/playlists"
)

// GetPlaylists https://developers.google.cn/youtube/v3/docs/playlists/list
func (this *Youtube) GetPlaylists(param GetPlaylistsParam) (results *Playlists, err error) {
	var api = this.BuildAPI(k_PLAYLISTS_API)
	err = this.doRequest("GET", api, param, &results)
	return results, err
}