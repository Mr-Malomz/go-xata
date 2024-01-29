package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-xata/data"
	"io/ioutil"
	"net/http"
)

var xataAPIKey = GetEnvVariable("XATA_API_KEY")
var baseURL = GetEnvVariable("XATA_DATABASE_URL")

func createRequest(method, url string, bodyData *bytes.Buffer) (*http.Request, error) {
	var req *http.Request
	var err error

	if method == "GET" || method == "DELETE" {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bodyData)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", xataAPIKey))

	return req, nil
}

func makeRequest(req *http.Request, target interface{}) error {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if target != nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(body, target)
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *Config) createProjectService(newProject *data.ProjectRequest) (*data.ProjectResponse, error) {
	createProject := data.ProjectResponse{}
	jsonData := data.Project{
		Name:        newProject.Name,
		Description: newProject.Description,
		Status:      newProject.Status,
	}

	postBody, _ := json.Marshal(jsonData)
	bodyData := bytes.NewBuffer(postBody)

	fullURL := fmt.Sprintf("%s:main/tables/Project/data", baseURL)
	req, err := createRequest("POST", fullURL, bodyData)
	if err != nil {
		return nil, err
	}

	err = makeRequest(req, &createProject)
	if err != nil {
		return nil, err
	}

	return &createProject, nil
}

func (app *Config) getProjectService(projectId string) (*data.Project, error) {
	projectDetails := data.Project{}

	fullURL := fmt.Sprintf("%s:main/tables/Project/data/%s", baseURL, projectId)
	req, err := createRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	err = makeRequest(req, &projectDetails)
	if err != nil {
		return nil, err
	}

	return &projectDetails, nil
}

func (app *Config) updateProjectService(updatedProject *data.ProjectRequest, projectId string) (*data.ProjectResponse, error) {
	updateProject := data.ProjectResponse{}
	jsonData := data.Project{
		Name:        updatedProject.Name,
		Description: updatedProject.Description,
		Status:      updatedProject.Status,
	}

	postBody, _ := json.Marshal(jsonData)
	bodyData := bytes.NewBuffer(postBody)

	fullURL := fmt.Sprintf("%s:main/tables/Project/data/%s", baseURL, projectId)
	req, err := createRequest("PUT", fullURL, bodyData)
	if err != nil {
		return nil, err
	}

	err = makeRequest(req, &updateProject)
	if err != nil {
		return nil, err
	}

	return &updateProject, nil
}

func (app *Config) deleteProjectService(projectId string) (string, error) {
	fullURL := fmt.Sprintf("%s:main/tables/Project/data/%s", baseURL, projectId)
	req, err := createRequest("DELETE", fullURL, nil)
	if err != nil {
		return "", err
	}

	err = makeRequest(req, nil)
	if err != nil {
		return "", err
	}

	return projectId, nil
}
