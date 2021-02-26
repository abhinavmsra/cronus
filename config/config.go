package config

import (
	"encoding/json"
	"io/ioutil"
	"os/exec"
	"strings"
)

type JiraConfig struct {
	ProjectCode  string `json:"code"`
	Organization string `json:"organization"`
	Email        string `json:"email"`
	AccessToken  string `json:"accessToken"`
}

type HarvestConfig struct {
	AccessToken string `json:"access_token"`
	AccountID   string `json:"account_id"`
	ProjectID   string `json:"project_id"`
	TaskID      string `json:"task_id"`
}

type Config struct {
	Jira    JiraConfig    `json:"jira"`
	Harvest HarvestConfig `json:"harvest"`
}

func Load() Config {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		panic("Couldnt find git root")
	}

	rootPath := strings.TrimSuffix(string(out), "\n") + "/cronus.json"
	jsonFile, err := ioutil.ReadFile(rootPath)
	if err != nil {
		panic("Couldnt find config file")
	}

	var config Config
	json.Unmarshal(jsonFile, &config)
	return config
}
