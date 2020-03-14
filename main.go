package main

import (
	"fmt"

	"github.com/naoto67/cal-app-go/src/database"
)

func main() {
	// debug code
	database.New()
	schedules, _ := database.DB.FindAllSchedules()
	fmt.Println(schedules)
}
