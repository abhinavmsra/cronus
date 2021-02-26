package jira

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	cronusConfig "github.com/abhinavmsra/cronus/config"
)

type JiraIssue struct {
	TaskID           string
	OrganizationName string
	Fields           struct {
		Summary string `json:"summary"`
	} `json:"fields"`
}

func (issue *JiraIssue) TaskURL() string {
	return fmt.Sprintf("https://%s.atlassian.net/browse/%s", issue.OrganizationName, issue.TaskID)
}

func FetchIssue(issueID string, config cronusConfig.Config) JiraIssue {
	client := &http.Client{}
	requestURL := fmt.Sprintf(
		"https://%s.atlassian.net/rest/api/latest/issue/%s",
		config.Jira.Organization,
		issueID,
	)
	req, err := http.NewRequest("GET", requestURL, nil)
	req.SetBasicAuth(config.Jira.Email, config.Jira.AccessToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	var issue = JiraIssue{
		TaskID:           issueID,
		OrganizationName: config.Jira.Organization,
	}
	err = json.Unmarshal(respBody, &issue)
	if err != nil {
		log.Fatal(err)
	}

	return issue
}
