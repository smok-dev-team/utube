package utube

import (
	"encoding/json"
	"github.com/smartwalle/going/request"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	k_YOUTUBE_API_URL = "https://www.googleapis.com/youtube"
)

type Youtube struct {
	key         string
	accessToken string
	apiDomain   string
}

func New(key, accessToken string) (client *Youtube) {
	if len(key) == 0 {
		return nil
	}
	client = &Youtube{}
	client.key = key
	client.accessToken = accessToken
	client.apiDomain = k_YOUTUBE_API_URL
	return client
}

func (this *Youtube) BuildAPI(paths ...string) string {
	var path = this.apiDomain
	for _, p := range paths {
		p = strings.TrimSpace(p)
		if len(p) > 0 {
			if strings.HasSuffix(path, "/") {
				path = path + p
			} else {
				if strings.HasPrefix(p, "/") {
					path = path + p
				} else {
					path = path + "/" + p
				}
			}
		}
	}
	return path
}

func (this *Youtube) doRequest(method, url string, param YoutubeParam, result interface{}) (err error) {
	var (
		req  *http.Request
		rep  *http.Response
		data []byte
	)

	var v = param.Params()
	v.Add("key", this.key)

	req, err = request.NewRequest(method, url, v)
	if err != nil {
		return err
	}
	if this.accessToken != "" {
		req.Header.Add("Authorization", "Bearer " + this.accessToken)
	}

	rep, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer rep.Body.Close()

	data, err = ioutil.ReadAll(rep.Body)
	if err != nil {
		return err
	}

	switch rep.StatusCode {
	case http.StatusOK:
		if result != nil {
			if err = json.Unmarshal(data, result); err != nil {
				return err
			}
		}
	default:
		var e = &ResponseError{}
		e.Response = rep
		if len(data) > 0 {
			if err = json.Unmarshal(data, e); err != nil {
				return err
			}
		}
		return e
	}
	return err
}
