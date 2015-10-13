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
	"testing"
	"fmt"
)

func TestUpload(t *testing.T) {
	var v, err = UploadVideo("http://7xk4hl.com2.z0.glb.qiniucdn.com/images/haha.mp4", "haha", "this is my haha", "", "SMOK", "", true)
	if err == nil {
		fmt.Println("video id:", v.Id)
	} else {
		fmt.Println(err)
	}

	fmt.Println("")
}

func TestGetList(t *testing.T) {
	var v, err = VideoList([]string{"spENjwU0TU8", "y4o2aySRDJc"}, "statistics")
	if err == nil {

		for _, item := range v.Items {
			fmt.Println(item.Id , "被赞了", item.Statistics.LikeCount)
		}
	}
}