package routes

import (
	"github.com/gin-gonic/gin"
)

func MakeRoutes(m *gin.Engine) {
	m.GET("/experiments/:id", GetAllExperimentById)
}
