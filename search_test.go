package utube

import (
	"fmt"
	"testing"
)

func TestYoutube_Search(t *testing.T) {
	fmt.Println("=====Search=====")
	var c = New("AIzaSyAeDwd1bXWY7Z86YxEqBTSOBNkbBfkM5i4")
	var p = SearchParam{}
	p.ChannelId = "UCgZd5ygXFoQry9KDGlddSBg"
	var rs, err = c.SearchChannel(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, s := range rs.Items {
		fmt.Println(s.GetId(), s.Id.Kind, s.Snippet.Title)
	}
}
