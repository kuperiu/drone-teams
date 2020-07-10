package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type (
	Repo struct {
		Owner string
		Name  string
	}

	Build struct {
		Tag      string
		Event    string
		Number   int
		Commit   string
		Ref      string
		Branch   string
		Author   Author
		Pull     string
		Message  Message
		DeployTo string
		Status   string
		Link     string
		Started  int64
		Created  int64
	}

	Author struct {
		Username string
		Name     string
		Email    string
		Avatar   string
	}

	Message struct {
		msg   string
		Title string
		Body  string
	}

	Job struct {
		Started int64
	}

	Teams struct {
		Summary         string            `json:"summary"`
		ThemeColor      string            `json:"themeColor"`
		Sections        []Section         `json:"sections"`
		PotentialAction []PotentialAction `json:"potentialAction"`
	}

	PotentialAction struct {
		Type    string   `json:"@type"`
		Name    string   `json:"name"`
		Targets []Target `json:"targets"`
	}

	Target struct {
		OS  string `json:"os"`
		URI string `json:"uri"`
	}

	Section struct {
		Facts         []Fact `json:"facts"`
		ActivityImage string `json:"activityImage"`
	}

	Fact struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
)

func (t Teams) exec(webhook string) error {

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(t)
	req, err := http.NewRequest("POST", webhook, b)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
