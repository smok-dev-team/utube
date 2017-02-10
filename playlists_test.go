package utube

import (
	"fmt"
	"testing"
)

func TestYoutube_Playlists(t *testing.T) {
	fmt.Println("=====Playlists=====")
	var c = New("AIzaSyAeDwd1bXWY7Z86YxEqBTSOBNkbBfkM5i4")
	var p = GetPlaylistsParam{}
	p.ChannelId = "UCv2LIk0KDlM413MIt8lt9fg"
	var ps, err = c.GetPlaylists(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ps)

	for _, p := range ps.Items {
		fmt.Println(p.Id, p.Snippet.Title)
	}
}
