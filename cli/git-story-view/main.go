package main

import (
	"fmt"
	"os"

	"github.com/vmware-tanzu-labs/git-story/adapters"
	"github.com/vmware-tanzu-labs/git-story/usecases"
	"gopkg.in/salsita/go-pivotaltracker.v2/v5/pivotal"
)

func newTracker() usecases.Tracker {
	apiToken := os.Getenv("TRACKER_API_TOKEN")
	client := pivotal.NewClient(apiToken)
	return adapters.NewPivotalTracker(client.Stories)
}

func main() {
	tracker := newTracker()
	story, error := usecases.GetStory(adapters.NewRepository(), tracker)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Printf("State: %v\n%v\n", story.State, story.Description)
	}
}
