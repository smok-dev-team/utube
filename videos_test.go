package utube

import (
	"fmt"
	"testing"
)

func TestYoutube_GetVideos(t *testing.T) {
	fmt.Println("=====Videos=====")
	var c = GetYoutube()
	var p = GetVideosParam{}
	p.Chart = "mostPopular"
	var vs, err = c.GetVideos(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range vs.Items {
		fmt.Println(v.Id, v.Snippet.Title, v.Snippet.Thumbnails.GetThumbnailURL())
	}
}

func TestYoutube_GetVideoCategories(t *testing.T) {
	fmt.Println("=====Video Categories=====")
	var c = GetYoutube()
	var p = GetVideoCategoriesParam{}
	p.RegionCode = "US"
	p.HL = "zh_CN"
	var vs, err = c.GetVideoCategories(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range vs.Items {
		fmt.Println(v.Snippet.ChannelId, v.Snippet.Title)
	}
}
