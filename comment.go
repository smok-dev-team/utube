package utube

const (
	k_COMMENT_THREADS_API = "/v3/commentThreads"
)

func (this *Youtube) GetCommentThreads(param GetCommentThreadsParam) (results *CommentThreads, err error) {
	var api = this.BuildAPI(k_COMMENT_THREADS_API)
	err = this.doRequest("GET", api, param, &results)
	return results, err
}