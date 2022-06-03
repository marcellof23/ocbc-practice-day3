package routes

import (
	"github.com/marcellof23/ocbc-practice-day3/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/employees", controllers.FindEmployees)
	r.POST("/employees", controllers.CreateEmployee)
	r.GET("/employees/:id", controllers.FindEmployee)
	r.PATCH("/employees/:id", controllers.UpdateEmployee)
	r.DELETE("employees/:id", controllers.DeleteEmployee)
	return r
}
