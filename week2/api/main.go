package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`          // mandatory
	Age           int    `json:"age"`           // mandatory
	StreetAddress string `json:"streetAddress"` // optional
	PostalCode    int    `json:"postalCode"`    // optional
	City          string `json:"city"`          // optional
}

// In-memory list of employees
var employees = []Employee{}
var nextID = 1 // Variable to automatically increment the ID

// CreateEmployee creates a new employee and adds it to the list
func CreateEmployee(name string, age int, streetAddress string, postalCode int, city string) Employee {
	emp := Employee{Id: nextID, Name: name, Age: age, StreetAddress: streetAddress, PostalCode: postalCode, City: city}
	nextID++ // Increment the ID for the next employee
	employees = append(employees, emp)
	return emp
}

// GetAllEmployees returns a list of all employees
func GetAllEmployees() []Employee {
	return employees
}

// GetEmployeeById retrieves an employee by ID
func GetEmployeeById(id int) (Employee, bool) {
	for _, emp := range employees {
		if emp.Id == id {
			return emp, true
		}
	}
	return Employee{}, false
}

// UpdateEmployee updates an employee by ID
func UpdateEmployee(id int, updated Employee) (Employee, bool) {
	for i, emp := range employees {
		if emp.Id == id {
			updated.Id = id // Ensure the ID remains the same
			employees[i] = updated
			return updated, true
		}
	}
	return Employee{}, false
}

// DeleteEmployee deletes an employee by ID
func DeleteEmployee(id int) bool {
	for i, emp := range employees {
		if emp.Id == id {
			employees = append(employees[:i], employees[i+1:]...)
			return true
		}
	}
	return false
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

	// Add new employee to the list
	createdEmployee := CreateEmployee(newEmployee.Name, newEmployee.Age, newEmployee.StreetAddress, newEmployee.PostalCode, newEmployee.City)
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

// DeleteEmployee is the DELETE handler for removing an employee by ID
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
	// Prepopulate employees for testing
	CreateEmployee("John Doe", 30, "123 Elm St", 12345, "New York")
	CreateEmployee("Jane Smith", 28, "456 Oak St", 67890, "Los Angeles")

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
