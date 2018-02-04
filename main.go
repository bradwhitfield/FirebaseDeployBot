package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudbuild/v1"
)

type response struct {
	Message string `json:"message"`
}

type BuildRequest struct {
	ProjectID string `json:"project-id"`
	Token     string `json:"token"`
	GitURL    string `json:"git-url"`
}

func F(w http.ResponseWriter, r *http.Request) {
	var br BuildRequest
	log.Println("Starting POST call.")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		log.Panic(err)
		return
	}
	err = json.Unmarshal(d, &br)
	if err != nil {
		w.WriteHeader(500)
		log.Panic(err)
	}
	build(br)

	if err := json.NewEncoder(w).Encode(response{Message: "Build Started."}); err != nil {
		w.WriteHeader(500)
		return
	}
}

func main() {
	test, err := ioutil.ReadFile("main.json")
	if err != nil {
		log.Panic(err)
	}
	var br BuildRequest
	err = json.Unmarshal(test, &br)
	if err != nil {
		log.Panic(err)
	}
	build(br)
}

func build(br BuildRequest) {
	// Google Default Auth
	log.Println("Starting build function.")
	ctx := context.TODO()
	client, err := google.DefaultClient(ctx, cloudbuild.CloudPlatformScope)
	if err != nil {
		log.Print(err)
		log.Panic("Error retrieving default credentials.")
	}
	svc, err := cloudbuild.New(client)
	if err != nil {
		log.Print(err)
		log.Panic("Unable to initialize GCP Container client.")
	}

	token := fmt.Sprintf("%s%s", "FIREBASE_DEPLOY_TOKEN=", br.Token)
	project := fmt.Sprintf("%s%s", "FIREBASE_PROJECT=", br.ProjectID)
	git := fmt.Sprintf("%s%s", "GIT_REPO=", br.GitURL)

	log.Println("Creating build steps.")
	steps := []*cloudbuild.BuildStep{
		&cloudbuild.BuildStep{
			Name: "bradwhitfield/firebasedeploybot",
			Env: []string{
				token,
				project,
				git,
			},
		},
	}

	build := &cloudbuild.Build{
		Steps: steps,
	}

	log.Println("Creating build and starting it.")
	create := svc.Projects.Builds.Create(br.ProjectID, build)
	results, err := create.Do()
	if err != nil {
		log.Panic(err)
	}
	rsp, _ := results.MarshalJSON()
	log.Printf(
		"Build execution request triggered with HTTP status code of %d and response of %s.\n",
		results.HTTPStatusCode,
		string(rsp),
	)
}
