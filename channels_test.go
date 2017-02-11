package utube

import (
	"testing"
	"fmt"
)

func TestYoutube_GetChannels(t *testing.T) {
	fmt.Println("=====Channels=====")
	var c = New("AIzaSyAeDwd1bXWY7Z86YxEqBTSOBNkbBfkM5i4")
	var p = GetChannelsParam{}
	//p.ForUsername = "RiPTrippers"
	p.Mine = 1
	var rs, err = c.GetChannels("ya29.GlvvAwaPc8nlY6OkiOlkXevTBHZ2HpqM_9ZFTN7sMT2TFGtER4Nlgzf90_bg2yK4vIHZZ3ylGMHyoNOmnVbYhAWaMlvyGUEJW6XrNJLZyz1K1EIujUEsD_UcKepC", p)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, r := range rs.Items {
		fmt.Println(r.Id, r.Snippet.Title)
	}
}
