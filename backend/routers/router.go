package routers

import (
	"net/http"

	"todoui/middleware"
	"todoui/models"

	"github.com/gin-gonic/gin"
)

type work struct {
	Data string `json:"data"`
}

// func routers() {
// 	router := gin.Default()

// 	router.POST("/task", PostTasks)
// 	//router.DELETE("/task/:id", deleteTask)

// }
func RegisterRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	router.POST("/key", PostTasks)
	router.GET("/todo", GetTasks)
	router.DELETE("/todo/:id", DeleteTasks)
	router.PUT("/todo/:id", UpdateTasks)
	return router
}

func GetTasks(c *gin.Context) {
	var get []models.PostTasks
	models.DB.Find(&get)

	c.IndentedJSON(http.StatusOK, gin.H{"todolist": get})
}
func PostTasks(c *gin.Context) {
	var todo work
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	kw := models.PostTasks{Data: &todo.Data}
	models.DB.Create(&kw)
	//c.IndentedJSON(http.StatusOK, gin.H{"message": "ho gya"})
	GetTasks(c)
}

func DeleteTasks(c *gin.Context) {
	var del models.PostTasks
	if err := models.DB.Where("id = ?", c.Param("id")).First(&del).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&del)

	GetTasks(c)
}

func UpdateTasks(c *gin.Context) {

	var change models.PostTasks
	if err := models.DB.Where("id = ?", c.Param("id")).First(&change).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.PostTasks
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&change).Updates(input)

	GetTasks(c)
}
