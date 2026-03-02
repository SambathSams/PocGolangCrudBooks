package controllers

import (
	"go-crud-backend/models"
	"go-crud-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProjects godoc
// @Summary      Get all projects
// @Description  Responds with the list of all projects as JSON.
// @Tags         projects
// @Produce      json
// @Success      200  {array}  models.Project
// @Router       /projects [get]
func GetProjects(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetAllProjects())
}

// GetProjectByID godoc
// @Summary      Get a project by ID
// @Description  Returns a single project based on the ID parameter.
// @Tags         projects
// @Produce      json
// @Param        id   path      string  true  "Project ID"
// @Success      200  {object}  models.Project
// @Failure      404  {object}  nil
// @Router       /projects/{id} [get]
func GetProjectByID(c *gin.Context) {
	id := c.Param("id")
	project, err := services.GetProjectByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, project)
}

// CreateProject godoc
// @Summary      Add a new project
// @Description  Takes a JSON input and adds it to the project collection.
// @Tags         projects
// @Accept       json
// @Produce      json
// @Param        project  body      models.Project  true  "Add project"
// @Success      201   {object}  models.Project
// @Router       /projects [post]
func CreateProject(c *gin.Context) {
	var newProject models.Project

	// 1. Bind JSON to the struct
	if err := c.ShouldBindJSON(&newProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// 2. Call the service (which now returns an error for validation/uniqueness)
	result, err := services.CreateProject(newProject)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

// UpdateProject godoc
// @Summary      Update an existing project
// @Description  Update the details of a project by its ID
// @Tags         projects
// @Accept       json
// @Produce      json
// @Param        id    path      string       true  "Project ID"
// @Param        project  body      models.UpdateProjectInput  true  "Updated project object"
// @Success      200   {object}  models.Project
// @Failure      400   {object}  nil "Invalid input"
// @Failure      404   {object}  nil "Project not found"
// @Router       /projects/{id} [put]
func UpdateProject(c *gin.Context) {
	id := c.Param("id")
	var updatedProject models.UpdateProjectInput

	if err := c.ShouldBindJSON(&updatedProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	project, err := services.UpdateProject(id, updatedProject)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}

// DeleteProject godoc
// @Summary      Delete a project
// @Description  Remove a project from the store by its ID
// @Tags         projects
// @Param        id   path      string  true  "Project ID"
// @Success      204  {object}  nil     "No Content"
// @Failure      404  {object}  nil "Project not found"
// @Router       /projects/{id} [delete]
func DeleteProject(c *gin.Context) {
	id := c.Param("id")

	err := services.DeleteProject(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
