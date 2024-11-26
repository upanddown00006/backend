package main

import (
	EmployeeController "backend/api/controller/employee"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//Employee API Method
	router.GET("/employee", EmployeeController.GET)
	router.GET("/employee/:id", EmployeeController.GetEmployeeByID)
	router.POST("/employee", EmployeeController.POST)
	router.POST("/employee/:id", EmployeeController.POST)
	router.PUT("/employee", EmployeeController.PUT)
	router.PUT("/employee/:id", EmployeeController.PUT)
	router.DELETE("/employee", EmployeeController.DELETE)
	router.DELETE("/employee/:id", EmployeeController.DELETE)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
