package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type GistRequest struct {
	Files       map[string]File `json:"files"`
	Description string          `json:"description"`
	Public      bool            `json:"public"`
}

type File struct {
	Content string `json:"content"`
}

type CreateGistParams struct {
	Files       map[string]File
	Description string
	Public      bool
}
type EditGistParams struct {
	ID          string
	Files       map[string]File
	Description string
	Public      bool
}

type DeleteGistParams struct {
	ID string
}

func CreateGist(c *APIClient, p *CreateGistParams) {
	r := GistRequest{
		p.Files,
		p.Description,
		p.Public,
	}

	j, err := json.Marshal(r)
	if err != nil {
		fmt.Printf("couldn't marshal request struct: %v", err)
		return
	}

	req, err := http.NewRequest(
		http.MethodPost,
		"https://api.github.com/gists",
		bytes.NewBuffer([]byte(j)),
	)
	if err != nil {
		fmt.Printf("couldn't create request: %v", err)
		return
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("TOKEn")))

	res, err := c.client.Do(req)
	if err != nil {
		fmt.Printf("couldn't send request: %v", err)
		return
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("couln't read request body: %v", err)
		return
	}

	fmt.Printf("body: %v", data)
	fmt.Printf("status code: %v", res.StatusCode)
}

func EditGist(c *APIClient, p *EditGistParams) {
	gr := GistRequest{
		p.Files,
		p.Description,
		p.Public,
	}

	json, err := json.Marshal(gr)
	if err != nil {
		fmt.Printf("couldn't marshal request struct: %v", err)
		return
	}

	req, err := http.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("https://api.github.com/gists/%s", p.ID),
		bytes.NewBuffer(json),
	)
	if err != nil {
		fmt.Printf("couldn't create request object: %v", err)
		return
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("TOKEN")))

	res, err := c.client.Do(req)
	if err != nil {
		fmt.Printf("client couldn't send request: %v", err)
		return
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error reading response body: %v", err)
	}
	fmt.Printf("body: %v", data)
	fmt.Printf("status: %v", res.StatusCode)
}

func DeleteGist(c *APIClient, p *DeleteGistParams) {
	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("https://api.github.com/gists/%s", p.ID),
		nil,
	)
	if err != nil {
		fmt.Printf("couldn't create request object: %v", err)
		return
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("TOKEN")))

	res, err := c.client.Do(req)
	if err != nil {
		fmt.Printf("client couldn't send request: %v", err)
		return
	}

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error reading response body; %v", err)
		return
	}
	fmt.Printf("status: %v", res.StatusCode)

}
