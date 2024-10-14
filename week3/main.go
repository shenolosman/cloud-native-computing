package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name          string `json:"name" gorm:"not null"` // mandatory
	Age           int    `json:"age" gorm:"not null"`  // mandatory
	StreetAddress string `json:"streetAddress"`        // optional
	PostalCode    int    `json:"postalCode"`           // optional
	City          string `json:"city"`                 // optional
}

var db *gorm.DB

// Initialize the database connection
func initDatabase() {
	var err error
	db, err = gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	if err != nil {
		panic("Error connecting/creating the sqlite db")
	}

	// Migrate the schema
	db.AutoMigrate(&Employee{})

	// Prepopulate the database with initial employees
	employees := []Employee{
		{Name: "John Doe", Age: 30, StreetAddress: "123 Elm St", PostalCode: 12345, City: "New York"},
		{Name: "Jane Smith", Age: 28, StreetAddress: "456 Oak St", PostalCode: 67890, City: "Los Angeles"},
		{Name: "Alice Johnson", Age: 35, StreetAddress: "789 Pine St", PostalCode: 54321, City: "Chicago"},
	}

	for _, emp := range employees {
		CreateEmployee(emp) // Create each employee in the database
	}
}

// CreateEmployee creates a new employee and adds it to the database
func CreateEmployee(emp Employee) Employee {
	db.Create(&emp) // Create a new employee record in the database
	return emp
}

// GetAllEmployees returns a list of all employees from the database
func GetAllEmployees() []Employee {
	var employees []Employee
	db.Find(&employees) // Retrieve all employees
	return employees
}

// GetEmployeeById retrieves an employee by ID from the database
func GetEmployeeById(id int) (Employee, bool) {
	var emp Employee
	result := db.First(&emp, id) // Find employee by ID
	return emp, result.Error == nil
}

// UpdateEmployee updates an employee by ID in the database
func UpdateEmployee(id int, updated Employee) (Employee, bool) {
	var emp Employee
	result := db.First(&emp, id) // Find employee by ID
	if result.Error != nil {
		return Employee{}, false // Employee not found
	}
	updated.ID = emp.ID // Ensure the ID remains the same
	db.Save(&updated)   // Save the updated employee record
	return updated, true
}

// DeleteEmployee deletes an employee by ID from the database
func DeleteEmployee(id int) bool {
	result := db.Delete(&Employee{}, id) // Delete employee by ID
	return result.RowsAffected > 0
}

// REST API Endpoints

// GetEmployees is the GET handler for retrieving all employees
func GetEmployees(c *gin.Context) {
	emps := GetAllEmployees()
	c.IndentedJSON(http.StatusOK, emps)
}

// GetEmployee is the GET handler for retrieving a specific employee by ID
func GetEmployee(c *gin.Context) {
	id := c.Param("id")
	empId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid employee ID"})
		return
	}
	emp, found := GetEmployeeById(empId)
	if found {
		c.JSON(http.StatusOK, emp)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
	}
}

// PostEmployee is the POST handler for creating a new employee
func PostEmployee(c *gin.Context) {
	var newEmployee Employee

	// Bind JSON to newEmployee struct
	if err := c.BindJSON(&newEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Add new employee to the database
	createdEmployee := CreateEmployee(newEmployee)
	c.JSON(http.StatusCreated, createdEmployee)
}

// PutEmployee is the PUT handler for updating an existing employee
func PutEmployee(c *gin.Context) {
	id := c.Param("id")
	empId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid employee ID"})
		return
	}

	var updatedEmployee Employee

	// Bind JSON to updatedEmployee struct
	if err := c.BindJSON(&updatedEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Update employee if found
	if emp, updated := UpdateEmployee(empId, updatedEmployee); updated {
		c.JSON(http.StatusOK, emp)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
	}
}

// DeleteEmployeeHandler is the DELETE handler for removing an employee by ID
func DeleteEmployeeHandler(c *gin.Context) {
	id := c.Param("id")
	empId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid employee ID"})
		return
	}

	// Delete employee if found
	if deleted := DeleteEmployee(empId); deleted {
		c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
	}
}

func main() {
	initDatabase() // Initialize the database connection

	r := gin.Default()

	// Simple ping endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// CRUD routes
	r.GET("/api/employees", GetEmployees)                 // GET all employees
	r.GET("/api/employees/:id", GetEmployee)              // GET a single employee by ID
	r.POST("/api/employees", PostEmployee)                // POST a new employee
	r.PUT("/api/employees/:id", PutEmployee)              // PUT (update) an employee by ID
	r.DELETE("/api/employees/:id", DeleteEmployeeHandler) // DELETE an employee by ID

	// Run the server
	r.Run() // listen and serve on 0.0.0.0:8080
}
