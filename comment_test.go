package utube

import (
	"testing"
	"fmt"
)

func TestYoutube_GetCommentThreads(t *testing.T) {
	fmt.Println("=====CommentThreads=====")
	var c = GetYoutube()
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

func TestYoutube_GetComments(t *testing.T) {
	fmt.Println("=====Comments=====")
	var c = GetYoutube()
	var p = GetCommentsParams{}
	p.Part.ShowSnippet()
	p.AppendId("z23libhysym5wndzr04t1aokgeo4euvmqqjrd42ynhe4bk0h00410")
	p.AppendId("z22fwzvzauj2vlweracdp433sp4uf4sddechmmbwvvlw03c010c")

	var cs, err = c.GetComments(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, c := range cs.Items {
		fmt.Println(c.Id, c.Snippet.TextDisplay)
	}
}
