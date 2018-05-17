package main

import (
	"fmt"
	"github.com/smartwalle/utube"
	"github.com/smartwalle/dbs"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"time"
)

var client *utube.Youtube
var db *sql.DB

func init() {
	client = utube.New("AIzaSyAeDwd1bXWY7Z86YxEqBTSOBNkbBfkM5i4", "")

	db, _ = sql.Open("mysql", "xxx")
}

func main() {
	var keywords = []string{"vape", "vaping", "vapor"}
	for _, keyword := range keywords {
		SearchChannel(keyword, "")
	}
}

func getVideos(channelId string) {
	playListId, err := getPlayListIdWithChannel(channelId)
	if err  != nil {
		return
	}

	getPlayListItemList(playListId, "")
}

func getPlayListIdWithChannel(channelId string) (playListId string, err error) {
	var p = utube.GetChannelsParam{}
	p.Part.ShowSnippet().ShowContentDetails()
	p.Id = channelId

	rs, err := client.GetChannels(p)
	if err != nil {
		fmt.Println("GetChannels", err)
		return
	}

	if len(rs.Items) > 0 {
		playListId = rs.Items[0].ContentDetails.RelatedPlayLists.Uploads
	}

	return playListId, err
}

func getPlayListItemList(playListId, nextPageToken string) (err error) {
	var p = utube.GetPlaylistItemsParam{}
	p.Part.ShowSnippet().ShowContentDetails()
	p.PlaylistId = playListId
	p.MaxResults = 50
	if nextPageToken != "" {
		p.PageToken = nextPageToken
	}

	rs, err := client.GetPlaylistItems(p)
	if err != nil {
		fmt.Println("GetPlaylistItems", err)
		return err
	}

	if len(rs.Items) > 0 {
		if err = insertVideo(rs.Items); err != nil {
			return err
		}
	} else {
		return
	}

	if rs.NextPageToken != "" {
		getPlayListItemList(playListId, rs.NextPageToken)
	}

	return err
}


func SearchChannel(keyword, nextPageToken string) {
	var p = utube.SearchParam{}
	p.Part.ShowSnippet()
	p.Q = keyword
	p.MaxResults = 50
	if nextPageToken != "" {
		p.PageToken = nextPageToken
	}

	var rs, err = client.SearchChannel(p)
	if err != nil {
		fmt.Println("SearchChannel", err)
		return
	}

	if len(rs.Items) > 0 {
		if err := insertChannel(rs.Items); err != nil {
			return
		}

		for _, item := range rs.Items {
			getVideos(item.GetId())
		}
	} else {
		return
	}

	if rs.NextPageToken != "" {
		SearchChannel(keyword, rs.NextPageToken)
	}
}

func insertChannel(items []*utube.SearchResult) (err error) {
	var ib = dbs.NewInsertBuilder()
	ib.Table("youtube_channel")
	ib.Columns("id", "title", "description", "thumbnails", "published_at")
	for _, item := range items {
		var pa, _ = time.Parse(time.RFC3339Nano, item.Snippet.PublishedAt)
		ib.Values(item.GetId(), item.Snippet.ChannelTitle, item.Snippet.Description, item.Snippet.Thumbnails.GetThumbnailURL(), pa)
		fmt.Println(item.GetId(), item.Snippet.ChannelTitle, item.Snippet.Description, item.Snippet.Thumbnails.GetThumbnailURL(), pa)
	}
	ib.Suffix("ON DUPLICATE KEY UPDATE title=VALUES(title), description=VALUES(description), thumbnails=VALUES(thumbnails), published_at=VALUES(published_at)")

	_, err = ib.Exec(db)
	return err
}

func insertVideo(items []*utube.PlaylistItem) (err error) {
	var ib = dbs.NewInsertBuilder()
	ib.Table("youtube_video")
	ib.Columns("id", "channel_id", "playlist_id", "title", "description", "thumbnails", "published_at")
	for _, item := range items {
		var pa, _ = time.Parse(time.RFC3339Nano, item.Snippet.PublishedAt)
		ib.Values(item.Snippet.ResourceId.VideoId, item.Snippet.ChannelId, item.Snippet.PlaylistId, item.Snippet.Title, item.Snippet.Description, item.Snippet.Thumbnails.GetThumbnailURL(), pa)
		fmt.Println(item.Snippet.ResourceId.VideoId, item.Snippet.ChannelId, item.Snippet.PlaylistId, item.Snippet.Title, item.Snippet.Description, item.Snippet.Thumbnails.GetThumbnailURL(), pa)
	}
	ib.Suffix("ON DUPLICATE KEY UPDATE channel_id=VALUES(channel_id), playlist_id=VALUES(playlist_id), title=VALUES(title), description=VALUES(description), thumbnails=VALUES(thumbnails), published_at=VALUES(published_at)")

	_, err = ib.Exec(db)
	return err
}
