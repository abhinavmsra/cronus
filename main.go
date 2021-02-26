package main

import (
	cronusConfig "github.com/abhinavmsra/cronus/config"
	"github.com/abhinavmsra/cronus/harvest"
	"github.com/abhinavmsra/cronus/jira"
	"github.com/abhinavmsra/cronus/vcs"
)

func main() {
	config := cronusConfig.Load()
	taskID := vcs.CurrentBranch(config)
	jiraIssue := jira.FetchIssue(taskID, config)
	harvest.Log(jiraIssue, config)
}
