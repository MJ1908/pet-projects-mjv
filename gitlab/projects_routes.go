package gitlab

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"petgitlab/controller"
)

var r = gin.Default()

var v1 = r.Group("/api")

func Init() {
	fmt.Println("Accediendo a la API de Gitlab...")
	r.Run()
}

func GitlabRegister() {
	v1.Group("/projects").GET("/", controller.ProyectList)
}
