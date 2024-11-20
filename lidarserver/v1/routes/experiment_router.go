package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"lidarserver.sqlc/app/lidarserver/v1/controller"
)

func GetAllExperimentById(c *gin.Context) {
	expId := c.Param("id")
	expid, _ := strconv.Atoi(expId)
	exp := controller.GetAllExperiments(expid)
	c.JSON(200, gin.H{"data": exp})
}
