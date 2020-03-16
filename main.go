package main

import (
	"github.com/naoto67/cal-app-go/src/api/router"
	"github.com/naoto67/cal-app-go/src/database"
)

func main() {
	database.New()

	r := router.New()
	r.Run()
}
