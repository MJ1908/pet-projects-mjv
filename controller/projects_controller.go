package controller

import (
	"errors"
	"net/http"
	"petgitlab/common"
	"petgitlab/model"
	services "petgitlab/services"
	"regexp"

	"github.com/gin-gonic/gin"
)

var data, err = services.GetAllProjects()

func ProyectList(c *gin.Context) {
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("proyects", errors.New("Invalid param")))
	} else {
		c.JSON(http.StatusOK, data)
	}
}

func ProjectListWithWrongName(c *gin.Context) {
	pattern := c.Param("pattern")
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("proyects", errors.New("Invalid param")))
	}

	var wrongNameProjects []model.Project

	for _, project := range data {
		match, _ := regexp.MatchString(pattern, project.Name)
		if match {
			wrongNameProjects = append(wrongNameProjects, project)
		}
	}

	c.JSON(200, gin.H{
		"projects": wrongNameProjects,
	})
}
