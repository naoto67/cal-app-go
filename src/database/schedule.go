package database

import (
	"github.com/naoto67/cal-app-go/src/model"
)

// errorを返すかは要検討
func (db *database) FindAllSchedules() ([]model.Schedule, error) {
	var schedules []model.Schedule
	db.orm.Find(&schedules)

	return schedules, nil
}

func (db *database) CreateSchedule(schedule model.Schedule) error {
	db.orm.Create(&schedule)
	return nil
}
