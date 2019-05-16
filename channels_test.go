package utube

import (
	"fmt"
	"os"
	"testing"
)

var client *Client

func TestMain(m *testing.M) {
	client = New("AIzaSyAeDwd1bXWY7Z86YxEqBTSOBNkbBfkM5i4", "")
	var exitCode = m.Run()
	os.Exit(exitCode)
}

func GetYoutube() *Client {
	return client
}

func TestYoutube_GetChannels(t *testing.T) {
	fmt.Println("=====Channels=====")
	var c = GetYoutube()
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
