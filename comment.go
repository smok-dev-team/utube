package utube

const (
	k_COMMENT_THREADS_API = "/v3/commentThreads"
	k_COMMENTS_API        = "/v3/comments"
)

func (this *Client) GetCommentThreads(param GetCommentThreadsParam) (results *CommentThreads, err error) {
	var api = this.BuildAPI(k_COMMENT_THREADS_API)
	err = this.doRequest("GET", api, param, &results)
	return results, err
}

func (this *Client) GetComments(param GetCommentsParams) (results *Comments, err error) {
	var api = this.BuildAPI(k_COMMENTS_API)
	err = this.doRequest("GET", api, param, &results)
	return results, err
}
