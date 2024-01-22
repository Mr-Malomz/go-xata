package api

import (
	"context"
	"fmt"
	"go-xata/data"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const appTimeout = time.Second * 10

func (app *Config) createProjectHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.ProjectRequest
		defer cancel()

		app.validateJsonBody(ctx, &payload)

		newProject := data.ProjectRequest{
			Name:        payload.Name,
			Description: payload.Description,
			Status:      payload.Status,
		}

		data, err := app.createProjectService(&newProject)
		if err != nil {
			app.errorJSON(ctx, err)
			return
		}

		app.writeJSON(ctx, http.StatusCreated, data)
	}
}

func (app *Config) getProjectHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		projectId := ctx.Param("projectId")
		defer cancel()

		data, err := app.getProjectService(projectId)

		if err != nil {
			app.errorJSON(ctx, err)
			return
		}

		app.writeJSON(ctx, http.StatusOK, data)
	}
}

func (app *Config) updateProjectHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		projectId := ctx.Param("projectId")
		var payload data.ProjectRequest
		defer cancel()

		app.validateJsonBody(ctx, &payload)

		newProject := data.ProjectRequest{
			Name:        payload.Name,
			Description: payload.Description,
			Status:      payload.Status,
		}

		data, err := app.updateProjectService(&newProject, projectId)

		if err != nil {
			app.errorJSON(ctx, err)
			return
		}

		app.writeJSON(ctx, http.StatusOK, data)
	}
}

func (app *Config) deleteProjectHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		projectId := ctx.Param("projectId")
		defer cancel()

		data, err := app.deleteProjectService(projectId)

		if err != nil {
			app.errorJSON(ctx, err)
			return
		}

		app.writeJSON(ctx, http.StatusAccepted, fmt.Sprintf("Project with ID: %s deleted successfully!!", data))
	}
}
