package router

import (
	"github.com/gin-gonic/gin"
	"github.com/naoto67/cal-app-go/src/api/handler"
)

func New() *gin.Engine {
	router := gin.Default()

	scheduleHandler := handler.NewScheduleHandler()
	scheduleGroup := router.Group("/schedules")
	{
		scheduleGroup.GET("", scheduleHandler.Index)
		scheduleGroup.GET("/:id", scheduleHandler.Show)
	}

	return router
}
