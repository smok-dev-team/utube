package utube

import (
	"testing"
	"fmt"
)

func TestGetVideoInfoWithVideoId(t *testing.T) {
	fmt.Println("=====Ext Videos=====")
	var info, err = GetVideoInfoWithVideoId("nLw-zd87wVk")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info.VideoId, info.Title, info.IURL)
}
