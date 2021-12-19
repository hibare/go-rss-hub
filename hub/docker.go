package hub

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	dockerHubBaseUrl = "https://hub.docker.com"
)

type Tag struct {
	Id          int       `json:"id"`
	LastUpdated time.Time `json:"last_updated"`
	Name        string    `json:"name"`
	Status      string    `json:"tag_status"`
}

type TagResponse struct {
	Count   int    `json:"count"`
	Next    string `json:"next"`
	Results []Tag
}

// GetDockerTags : Fetch docker image tags
func GetDockerTags(user string, repository string) []Tag {

	tags := []Tag{}

	url := dockerHubBaseUrl + "/v2/repositories/" + user + "/" + repository + "/tags"

	for {
		var tagResponse TagResponse
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		if res.Body != nil {
			defer res.Body.Close()
		}

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal(body, &tagResponse); err != nil { // Parse []byte to the go struct pointer
			log.Fatal(err)
		}

		tags = append(tags, tagResponse.Results...)

		if tagResponse.Next == "" {
			return tags
		}

		url = tagResponse.Next
	}

}
