package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/naoto67/cal-app-go/src/database"
)

type ScheduleHandler interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
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

func (handler scheduleHandler) Show(c *gin.Context) {
	// 1 fetch schedules from rdb
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	schedule, err := database.DB.FindScheduleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	// 2 write response
	c.JSON(http.StatusOK, schedule)
}
