package utube

import (
	"fmt"
	"testing"
)

func TestYoutube_Playlists(t *testing.T) {
	fmt.Println("=====Playlists=====")
	var c = GetYoutube()
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
	var c = GetYoutube()
	var p = GetPlaylistItemsParam{}
	p.PlaylistId = "UUgZd5ygXFoQry9KDGlddSBg"
	var ps, err = c.GetPlaylistItems(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, p := range ps.Items {
		fmt.Println(p.Id, p.Snippet.ResourceId.GetId())
	}

}