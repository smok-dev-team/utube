package utube

import (
	"testing"
	"fmt"
)

func NewYoutube() *Youtube {
	return New("AIzaSyAeDwd1bXWY7Z86YxEqBTSOBNkbBfkM5i4", "")
}

func TestYoutube_GetChannels(t *testing.T) {
	fmt.Println("=====Channels=====")
	var c = NewYoutube()
	var p = GetChannelsParam{}
	p.ForUsername = "RiPTrippers"
	var rs, err = c.GetChannels(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, r := range rs.Items {
		fmt.Println(r.Id, r.Snippet.Title)
	}
}
