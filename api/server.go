package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/feeds"
	"github.com/gorilla/mux"
	"github.com/hibare/go-rss-hub/hub"
	"github.com/hibare/go-rss-hub/util"
)

var (
	listenAddr string
	listenPort string
)

func init() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	listenAddr = config.ListenAddr
	listenPort = config.ListenPort
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Good to see you")
}

func dockerTags(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	urlParams := r.URL.Query()
	user := vars["user"]
	repository := vars["repository"]
	include := strings.Split(urlParams.Get("include"), ",")
	exclude := strings.Split(urlParams.Get("exclude"), ",")

	log.Printf("Processing: user %v, repository %v, include %v, exclude %s", user, repository, include, exclude)

	repo := hub.GetRepo(user, repository)
	tags := hub.GetDockerTags(user, repository)

	feed := &feeds.Feed{
		Title:       user + "/" + repository + " | Docker Hub Images",
		Link:        &feeds.Link{Href: "https://hub.docker.com/r/" + user + "/" + repository},
		Description: repo.Description,
		Author:      &feeds.Author{Name: user},
		Created:     time.Now(),
	}

	feed.Items = make([]*feeds.Item, 0)

	for _, tag := range tags {

		feed.Items = append(feed.Items, &feeds.Item{
			Title:       user + "/" + repository + ":" + tag.Name,
			Link:        &feeds.Link{Href: "https://hub.docker.com/r/" + user + "/" + repository + "/tags?name=" + tag.Name},
			Description: repo.Description,
			Content:     fmt.Sprint("Docker image ID: ", tag.Id, ", Status: ", tag.Status),
			Author:      &feeds.Author{Name: user},
			Created:     tag.LastUpdated,
			Id:          fmt.Sprint("tag:hub.docker.com,", tag.LastUpdated.Format("2006-01-02"), ":/r/", user, "/", repository, "/tags?name=", tag.Name),
		})
	}

	atom, err := feed.ToAtom()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, atom)
}

// HandleRequests : API request handler
func HandleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", http.HandlerFunc(home)).Methods("GET")
	r.HandleFunc("/ping/", ping).Methods("GET")
	r.HandleFunc("/docker/{user}/{repository}/tags/", dockerTags).Methods("GET")

	log.Printf("Listening for address %s on port %s\n", listenAddr, listenPort)
	log.Fatal(http.ListenAndServe(listenAddr+":"+listenPort, r))
}
