package utube

import (
	"testing"
	"fmt"
)

func TestYoutube_GetVideos(t *testing.T) {
	fmt.Println("=====Videos=====")
	var c = New("AIzaSyAeDwd1bXWY7Z86YxEqBTSOBNkbBfkM5i4")
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
