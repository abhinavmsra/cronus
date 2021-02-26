package vcs

import (
	"errors"
	"log"
	"os/exec"
	"regexp"
	"strings"

	cronusConfig "github.com/abhinavmsra/cronus/config"
)

func CurrentBranch(config cronusConfig.Config) (string, error) {
	out, err := exec.Command("git", "branch", "--show-current").Output()
	if err != nil {
		log.Fatal(err)
	}

	branchName := strings.Replace(string(out), "\n", "", -1)
	projectCode := strings.ToUpper(config.Jira.ProjectCode)
	regex := regexp.MustCompile(projectCode + `-[0-9]{1,}`)
	regexMatches := regex.FindAllString(branchName, -1)

	if len(regexMatches) > 0 {
		return regexMatches[0], nil
	} else {
		return "", errors.New("Couldnt find issue id from branch name")
	}
}
