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

func (app *Config) createProjectService(newProject *data.ProjectRequest) (*data.ProjectResponse, error) {
	createProject := data.ProjectResponse{}
	jsonData := data.Project{
		Name:        newProject.Name,
		Description: newProject.Description,
		Status:      newProject.Status,
	}

	postBody, _ := json.Marshal(jsonData)
	bodyData := bytes.NewBuffer(postBody)

	//api request
	fullURL := fmt.Sprintf("%s:main/tables/Project/data", baseURL)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", fullURL, bodyData)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", xataAPIKey))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &createProject)
	if err != nil {
		return nil, err
	}
	return &createProject, nil
}

func (app *Config) getProjectService(projectId string) (*data.Project, error) {
	projectDetails := data.Project{}

	//api request
	fullURL := fmt.Sprintf("%s:main/tables/Project/data/%s", baseURL, projectId)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fullURL, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", xataAPIKey))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &projectDetails)
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

	//api request
	fullURL := fmt.Sprintf("%s:main/tables/Project/data/%s", baseURL, projectId)
	client := &http.Client{}
	req, _ := http.NewRequest("PUT", fullURL, bodyData)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", xataAPIKey))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &updateProject)
	if err != nil {
		return nil, err
	}
	return &updateProject, nil
}

func (app *Config) deleteProjectService(projectId string) (string, error) {

	//api request
	fullURL := fmt.Sprintf("%s:main/tables/Project/data/%s", baseURL, projectId)
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", fullURL, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", xataAPIKey))

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return projectId, nil
}