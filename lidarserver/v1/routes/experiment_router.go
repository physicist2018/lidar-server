package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"lidarserver.sqlc/app/lidarserver/v1/controller"
)

func GetExperimentById(c *gin.Context) {
	expId := c.Param("id")
	expid, _ := strconv.Atoi(expId)
	exp := controller.GetExperimentById(expid)
	c.JSON(200, gin.H{"data": exp})
}

func GetAllExperiments(c *gin.Context) {
	exps := controller.GetAllExperients()
	c.JSON(200, gin.H{"data": exps})
}
