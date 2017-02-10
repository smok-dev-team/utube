package utube

import (
	"fmt"
	"net/http"
)

type ResponseError struct {
	Response *http.Response `json:"-"`
	RepError struct {
		Errors  []*ErrorInfo `json:"errors"`
		Code    int          `json:"code"`
		Message string       `json:"message"`
	} `json:"error"`
}

func (this *ResponseError) Error() string {
	return fmt.Sprintf("[%s]%d %s [Message]%s", this.Response.Request.Method, this.Response.StatusCode, this.Response.Request.URL, this.RepError.Message)
}

type ErrorInfo struct {
	Domain       string `json:"domain"`
	Reason       string `json:"reason"`
	Message      string `json:"message"`
	LocationType string `json:"locationType"`
	Location     string `json:"location"`
}
