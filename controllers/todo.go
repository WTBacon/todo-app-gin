package controllers

import (
	"net/http"

	"github.com/WTBacon/todo-app-gin/models"
	"github.com/gin-gonic/gin"
)

func FetchAllTodo(c *gin.Context) {
	var todo []models.Todo
	if err := models.FetchAllTodo(&todo); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func FetchSingleTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := models.FetchSingleTodo(&todo, id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.BindJSON(&todo); err != nil {
		return
	}
	if err := models.CreateTodo(&todo); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	if err := models.FetchSingleTodo(&todo, id); err != nil {
		c.JSON(http.StatusNotFound, todo)
	}
	if err := c.BindJSON(&todo); err != nil {
		return
	}
	if err := models.UpdateTodo(&todo); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	if err := models.DeleteTodo(&todo, id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
