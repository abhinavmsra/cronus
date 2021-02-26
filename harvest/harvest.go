package harvest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	cronusConfig "github.com/abhinavmsra/cronus/config"
	"github.com/abhinavmsra/cronus/jira"
)

type externalReference struct {
	ID        string `json:"id"`
	Permalink string `json:"permalink"`
}

type timeEntry struct {
	ProjectID         string            `json:"project_id"`
	TaskID            string            `json:"task_id"`
	SpentDate         string            `json:"spent_date"`
	Notes             string            `json:"notes"`
	ExternalReference externalReference `json:"external_reference"`
}

func Log(issue jira.JiraIssue, config cronusConfig.Config) {
	todayTime := time.Now().Format(time.RFC3339)
	reference := externalReference{issue.TaskID, issue.TaskURL()}
	entry := timeEntry{
		config.Harvest.ProjectID,
		config.Harvest.TaskID,
		todayTime,
		issue.Fields.Summary,
		reference,
	}

	client := &http.Client{}
	requestURL := "https://api.harvestapp.com/v2/time_entries"
	requestPayload, err := json.Marshal(entry)

	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(requestPayload))
	req.Header.Add("User-Agent", "Harvest CLI (c01mo283u@relay.firefox.com)")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+config.Harvest.AccessToken)
	req.Header.Add("Harvest-Account-Id", config.Harvest.AccountID)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
