package utube

import (
	"fmt"
	"testing"
)

// 搜索指定 Channel 下的视频
func TestYoutube_SearchVideoWithChannel(t *testing.T) {
	fmt.Println("=====Search=====")
	var c = GetYoutube()
	var p = SearchParam{}
	p.ChannelId = "UCBR8-60-B28hp2BmDPdntcQ "
	var rs, err = c.SearchVideo(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, s := range rs.Items {
		fmt.Println(s.GetId(), s.Id.Kind, s.Snippet.Title)
	}
}
