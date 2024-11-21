package routes

import (
	"github.com/gin-gonic/gin"
)

func MakeRoutes(m *gin.Engine) {
	m.GET("/experiments", GetAllExperiments)
	m.GET("/experiments/:id", GetExperimentById)
}
