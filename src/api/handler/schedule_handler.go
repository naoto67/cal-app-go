package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naoto67/cal-app-go/src/database"
)

type ScheduleHandler interface {
	Index(c *gin.Context)
}

type scheduleHandler struct{}

func NewScheduleHandler() ScheduleHandler {
	var handler ScheduleHandler
	handler = scheduleHandler{}
	return handler
}

func (handler scheduleHandler) Index(c *gin.Context) {
	// 1 fetch schedules from rdb
	schedules, _ := database.DB.FindAllSchedules()
	// 2 write response
	c.JSON(http.StatusOK, schedules)
}
