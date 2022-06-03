package main

import (
	"github.com/marcellof23/ocbc-practice-day3/models"
	"github.com/marcellof23/ocbc-practice-day3/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Employee{})

	defer db.Close()
	r := routes.SetupRoutes(db)
	r.Run()
}
