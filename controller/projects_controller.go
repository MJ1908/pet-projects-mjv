package controller

import (
	"errors"
	"net/http"
	"petgitlab/common"
	"petgitlab/model"
	services "petgitlab/services"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

var data, err = services.GetAllProjects()
var layoutISO = "0001-01-01"

func ProyectList(c *gin.Context) {
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("proyects", errors.New("Invalid param")))
	} else {
		c.JSON(http.StatusOK, data)
	}
}

func WrongNameProjectsList(c *gin.Context) {
	pattern := c.Param("pattern")
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("proyects", errors.New("Invalid param")))
	}

	var wrongNameProjects []model.Project

	for _, project := range data {
		match, _ := regexp.MatchString(pattern, project.Name)
		if !match {
			wrongNameProjects = append(wrongNameProjects, project)
		}
	}

	c.JSON(200, gin.H{
		"projects": wrongNameProjects,
	})
}

func TimeWindowProjectsList(c *gin.Context) {
	startDate, _ := time.Parse(layoutISO, c.Param("startDate"))
	endDate, _ := time.Parse(layoutISO, c.Param("endDate"))

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("proyects", errors.New("Invalid param")))
	}

	var timeWindowProjects []model.Project

	for _, project := range data {

		projectCreationDate, _ := time.Parse(layoutISO, project.CreatedAt[0:10])
		if services.IsInTimeWindow(projectCreationDate, startDate, endDate) {
			timeWindowProjects = append(timeWindowProjects, project)
		}

	}

	c.JSON(200, gin.H{
		"projects": timeWindowProjects,
	})
}
