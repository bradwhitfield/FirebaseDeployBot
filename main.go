package main

import (
	"encoding/json"
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

func F(w http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		log.Panic(err)
		return
	}

	if err := json.NewEncoder(w).Encode(response{Message: string(d)}); err != nil {
		w.WriteHeader(500)
		return
	}
}

func main() {
	id := "nimble-card-172800"
	build(id)
}

func build(id string) {
	// Google Default Auth
	ctx := context.TODO()
	client, err := google.DefaultClient(ctx, cloudbuild.CloudPlatformScope)
	if err != nil {
		log.Print(err)
		log.Panic("Error retrieving default credentials.")
	}
	log.Print(client)
	svc, err := cloudbuild.New(client)
	if err != nil {
		log.Print(err)
		log.Panic("Unable to initialize GCP Container client.")
	}

	steps := []*cloudbuild.BuildStep{
		&cloudbuild.BuildStep{
			Name: "hello-world",
		},
	}

	build := &cloudbuild.Build{
		Steps: steps,
	}

	create := svc.Projects.Builds.Create(id, build)
	results, err := create.Do()
	if err != nil {
		log.Panic(err)
	}
	log.Print(results.HTTPStatusCode)
	rsp, _ := results.MarshalJSON()
	log.Print(string(rsp))
}
