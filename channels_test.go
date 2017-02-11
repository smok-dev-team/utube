package utube

import (
	"testing"
	"fmt"
)

func NewYoutube() *Youtube {
	return New("AIzaSyAeDwd1bXWY7Z86YxEqBTSOBNkbBfkM5i4", "ya29.GlvvAwaPc8nlY6OkiOlkXevTBHZ2HpqM_9ZFTN7sMT2TFGtER4Nlgzf90_bg2yK4vIHZZ3ylGMHyoNOmnVbYhAWaMlvyGUEJW6XrNJLZyz1K1EIujUEsD_UcKepC")
}

func TestYoutube_GetChannels(t *testing.T) {
	fmt.Println("=====Channels=====")
	var c = NewYoutube()
	var p = GetChannelsParam{}
	//p.ForUsername = "RiPTrippers"
	p.Mine = 1
	var rs, err = c.GetChannels(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, r := range rs.Items {
		fmt.Println(r.Id, r.Snippet.Title)
	}
}
