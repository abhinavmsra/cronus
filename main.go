package main

import (
	cronusConfig "github.com/abhinavmsra/cronus/config"
	"github.com/abhinavmsra/cronus/harvest"
	"github.com/abhinavmsra/cronus/jira"
	"github.com/abhinavmsra/cronus/vcs"
)

func main() {
	config := cronusConfig.Load()
	taskID, err := vcs.CurrentBranch(config)

	// Only log time when `taskID` is present
	if err == nil {
		jiraIssue := jira.FetchIssue(taskID, config)
		harvest.Log(jiraIssue, config)
	}
}
