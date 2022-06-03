package controllers

import (
	"net/http"
	"time"

	"github.com/marcellof23/ocbc-practice-day3/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateEmployeeInput struct {
	Name      string `json:"name"`
	BirthDate string `json:"birthDate"`
	Address   string `json:"address"`
	Job       string `json:"job"`
	JoinDate  string `json:"joinDate"`
}

type UpdateEmployeeInput struct {
	Name      string `json:"name"`
	BirthDate string `json:"birthDate"`
	Address   string `json:"address"`
	Job       string `json:"job"`
	JoinDate  string `json:"joinDate"`
}

// GET /Employees
// Get all Employees
func FindEmployees(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var Employees []models.Employee
	db.Find(&Employees)

	c.JSON(http.StatusOK, gin.H{"data": Employees})
}

// POST /Employees
// Create new Employee
func CreateEmployee(c *gin.Context) {
	// Validate input
	var input CreateEmployeeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	date := "2006-01-02"
	joinDate, _ := time.Parse(date, input.JoinDate)
	birthDate, _ := time.Parse(date, input.BirthDate)

	// Create Employee
	Employee := models.Employee{Name: input.Name, BirthDate: birthDate, Address: input.Address, Job: input.Job, JoinDate: joinDate}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&Employee)

	c.JSON(http.StatusOK, gin.H{"data": Employee})
}

// GET /Employees/:id
// Find a Employee
func FindEmployee(c *gin.Context) { // Get model if exist
	var Employee models.Employee

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&Employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Employee})
}

// PATCH /Employees/:id
// Update a Employee
func UpdateEmployee(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var Employee models.Employee
	if err := db.Where("id = ?", c.Param("id")).First(&Employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateEmployeeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	date := "2006-01-02"
	joinDate, _ := time.Parse(date, input.JoinDate)
	birthDate, _ := time.Parse(date, input.BirthDate)

	var updatedInput models.Employee
	updatedInput.Name = input.Name
	updatedInput.BirthDate = birthDate
	updatedInput.Address = input.Address
	updatedInput.Job = input.Job
	updatedInput.JoinDate = joinDate

	db.Model(&Employee).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": Employee})
}

// DELETE /Employees/:id
// Delete a Employee
func DeleteEmployee(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var employee models.Employee
	if err := db.Where("id = ?", c.Param("id")).First(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Employee not found!"})
		return
	}

	db.Delete(&employee)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
