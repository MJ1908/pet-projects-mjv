package controller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"petgitlab/common"
	"petgitlab/model"

	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	config "petgitlab/configs"
)

func ProyectList(c *gin.Context) {
	filter := c.Query("filter")

	response, err := http.Get(config.GitlabApiPath)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("proyects", errors.New("Invalid param")))
	} else {
		responseData, _ := ioutil.ReadAll(response.Body)

		var data, results, resultFiltered []model.Project

		_ = json.Unmarshal(responseData, &data)
		if filter == "date" {
			desde := c.Query("desde")
			hasta := c.Query("hasta")
			initialDay, _ := time.Parse(time.RFC3339, desde)
			lastDay, _ := time.Parse(time.RFC3339, hasta)
			for _, project := range data {
				creationDay, _ := time.Parse(time.RFC3339, project.CreatedAt)
				if creationDay.Before(lastDay) && creationDay.After(initialDay) {
					results = append(results, project)
				}
			}
			c.JSON(http.StatusOK, results)
		} else if filter == "projectsDefined" {
			for _, project := range data {
				if project.ReadmeURL != nil {
					results = append(results, project)
				}
			}
			for _, project := range results {
				id := strconv.Itoa(project.ID)
				response, _ := http.Get(config.GitlabApiPath + id + "/pipelines")
				responseData, _ := ioutil.ReadAll(response.Body)
				var pipelines []model.Pipeline
				_ = json.Unmarshal(responseData, &pipelines)
				if len(pipelines) > 0 {
					resultFiltered = append(resultFiltered, project)
				}
			}
			c.JSON(http.StatusOK, resultFiltered)
		}
		c.JSON(http.StatusOK, data)
	}

}
