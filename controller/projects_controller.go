package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"petgitlab/common"
	services "petgitlab/services"
)

var data, err = services.GetAllProjects()

func ProyectList(c *gin.Context) {
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("proyects", errors.New("Invalid param")))
	} else {
		c.JSON(http.StatusOK, data)
	}
}
