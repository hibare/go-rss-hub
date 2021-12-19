package hub

import (
	"encoding/json"
	"fmt"
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

type RepoReponse struct {
	User            string    `json:"user"`
	Name            string    `json:"name"`
	Namespace       string    `json:"namespace"`
	RepositoryType  string    `json:"repository_type"`
	Status          int       `json:"status"`
	Description     string    `json:"description"`
	StartCount      int       `json:"star_count"`
	PullCount       int       `json:"pull_count"`
	LastUpdated     time.Time `json:"last_updated"`
	HubUser         string    `json:"hub_user"`
	FullDescription string    `json:"full_description"`
}

// GetRepo : Fetch repository details
func GetRepo(user string, repository string) RepoReponse {

	url := fmt.Sprint(dockerHubBaseUrl, "/v2/repositories/", user, "/", repository, "/")

	repoResponse := RepoReponse{}

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

	if err := json.Unmarshal(body, &repoResponse); err != nil {
		log.Fatal(err)
	}

	return repoResponse
}

// GetDockerTags : Fetch docker image tags
func GetDockerTags(user string, repository string) []Tag {

	tags := []Tag{}

	url := fmt.Sprint(dockerHubBaseUrl, "/v2/repositories/", user, "/", repository, "/tags")

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
