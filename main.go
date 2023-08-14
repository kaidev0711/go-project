package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO: checkHealth
func routeHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// TODO: Student service
type Student struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Age      int    `json:"age"`
}

var Students = []Student{
	{ID: 1, FullName: "Tuan", Age: 22},
	{ID: 2, FullName: "Thin", Age: 23},
}

func routeGetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, Students)
}

func routePostStudent(c *gin.Context) {
	var student Student

	err := c.Bind(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Error",
		})
		return
	}
	student.ID = Students[len(Students)-1].ID + 1
	Students = append(Students, student)
	c.JSON(http.StatusCreated, student)
}

func routePutStudent(c *gin.Context) {
	var studentPayload Student
	var studentLocal Student
	var newStudent []Student

	err := c.BindJSON(&studentPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Error",
		})
		return
	}

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Error",
		})
		return
	}

	for _, studentElemt := range Students {
		if studentElemt.ID == id {
			studentLocal = studentElemt
		}
	}
	if studentLocal.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Error",
		})
		return
	}
	studentLocal.FullName = studentPayload.FullName
	studentLocal.Age = studentPayload.Age

	for _, studentElement := range Students {
		if id == studentElement.ID {
			newStudent = append(newStudent, studentLocal)
		} else {
			newStudent = append(newStudent, studentElement)
		}
	}

	Students = newStudent
	c.JSON(http.StatusOK, studentLocal)
}

func routeDeleteStudent(c *gin.Context) {
	var newStudent []Student
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Error",
		})
		return
	}
	for _, studentElement := range Students {
		if studentElement.ID == id {
			newStudent = append(newStudent, studentElement)
		}
	}
	Students = newStudent
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete student success",
	})
}

func routeGetStudent(c *gin.Context) {
	var student Student

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Error",
		})
		return
	}
	for _, studentElement := range Students {
		if studentElement.ID == id {
			student = studentElement
		}
	}
	if student.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message_error": "Error",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}

func main() {
	service := gin.Default()

	getRoutes(service)

	service.Run()
}

func getRoutes(c *gin.Engine) *gin.Engine {
	c.GET("/health", routeHealth)

	groupStudents := c.Group("/students")
	groupStudents.GET("/", routeGetStudents)
	groupStudents.POST("/", routePostStudent)
	groupStudents.PUT("/:id", routePutStudent)
	groupStudents.DELETE("/:id", routeDeleteStudent)
	groupStudents.GET("/:id", routeGetStudent)
	return c
}
