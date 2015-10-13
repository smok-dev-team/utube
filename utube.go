/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
	http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utube

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"google.golang.org/api/youtube/v3"
	"errors"
)

type customReader struct {
	ReadCloser io.ReadCloser
	Progress func(Bps int, total int)

	totalBytes int
	lapTime    time.Time
	startTime  time.Time
}

func GetService() (*youtube.Service, error) {
	client, err := BuildOAuthHTTPClient([]string{youtube.YoutubeScope, youtube.YoutubeForceSslScope, youtube.YoutubeReadonlyScope, youtube.YoutubeUploadScope, youtube.YoutubepartnerScope, youtube.YoutubepartnerChannelAuditScope})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error building OAuth client: %v", err))
	}

	service, err := youtube.New(client)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error creating YouTube client: %v", err))
	}
	return service, err
}

func UploadVideo(filename string, title string, description string, category string, keywords string, privacy string, showProgress bool) (*youtube.Video, error) {
	if len(filename) == 0 {
		return nil, errors.New("You must provide a filename of a video file to upload")
	}

	service, err := GetService()
	if err != nil {
		return nil, err
	}

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       title,
			Description: description,
			CategoryId:  category,
		},
		Status: &youtube.VideoStatus{PrivacyStatus: privacy},
	}

	if strings.Trim(keywords, "") != "" {
		upload.Snippet.Tags = strings.Split(keywords, ",")
	}

	call := service.Videos.Insert("snippet,status", upload)

	reader := &customReader{}
	if showProgress {
		reader.Progress = progress
	}

	if strings.HasPrefix(filename, "http") {
		resp, err := http.Get(filename)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error opening %v: %v", filename, err))
		}
		reader.ReadCloser = resp.Body
	} else {
		file, err := os.Open(filename)
		defer file.Close()
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error opening %v: %v", filename, err))
		}
		reader.ReadCloser = file
	}

	response, err := call.Media(reader).Do()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error making YouTube API call: %v", err))
	}

	return response, nil
}

func VideoList(videoIdList []string, part string) (*youtube.VideoListResponse, error) {
	service, err := GetService()
	if err != nil {
		return nil, err
	}

	call := service.Videos.List(part)

	call.Id(strings.Join(videoIdList, ","))

	response, err := call.Do()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error making YouTube API call: %v", err))
	}
	return response, err
}

func progress(Bps int, total int) {
	fmt.Printf("\rTransfer rate %.2f Mbps, total %d", float32(Bps*8)/(1000*1000), total)
}

func (r *customReader) Read(p []byte) (n int, err error) {
	if r.startTime.IsZero() {
		r.startTime = time.Now()
	}
	if r.lapTime.IsZero() {
		r.lapTime = time.Now()
	}
	if len(p) == 0 {
		return 0, nil
	}
	n, err = r.ReadCloser.Read(p)
	r.totalBytes += n

	if time.Since(r.lapTime) >= time.Second || err == io.EOF {
		if r.Progress != nil {
			r.Progress(r.totalBytes/int(time.Since(r.startTime).Seconds()), r.totalBytes)
		}
		r.lapTime = time.Now()
	}

	return n, err
}