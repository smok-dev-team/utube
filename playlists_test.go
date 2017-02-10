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

	for _, p := range ps.Items {
		fmt.Println(p.Id, p.Snippet.Title)
	}
}


func TestYoutube_GetPlaylistItems(t *testing.T) {
	fmt.Println("=====PlaylistItems=====")
	var c = New("AIzaSyAeDwd1bXWY7Z86YxEqBTSOBNkbBfkM5i4")
	var p = GetPlaylistItemsParam{}
	p.PlaylistId = "PLtCPCoZDB0vqC-EsAX3GVNW6N3bHLlaS8"
	var ps, err = c.GetPlaylistItems(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, p := range ps.Items {
		fmt.Println(p.Id, p.Snippet.ResourceId.GetId())
	}

}