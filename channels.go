package utube

const (
	k_CHANNELS_API = "/v3/channels"
)

// GetChannels https://developers.google.com/youtube/v3/docs/channels/list
func (this *Youtube) GetChannels(param GetChannelsParam) (results *Channels, err error) {
	var api = this.BuildAPI(k_CHANNELS_API)
	err = this.doRequest("GET", api, param, &results)
	return results, err
}