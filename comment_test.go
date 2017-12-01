package utube

import (
	"testing"
	"fmt"
)

func TestYoutube_GetCommentThreads(t *testing.T) {
	fmt.Println("=====GetCommentThreads=====")
	var c = NewYoutube()
	var p = GetCommentThreadsParam{}
	p.Part.ShowSnippet()
	p.Part.ShowReplies()
	p.VideoId = "S5eFdckiA-c"
	var cts, err = c.GetCommentThreads(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, c := range cts.Items {
		fmt.Println(c.Id, c.Kind)
	}
}
