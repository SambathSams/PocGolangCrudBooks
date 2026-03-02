package services

import (
	"errors"
	"go-crud-backend/models"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

var projects = []models.Project{
	{ID: "1", Name: "Project", Task: "Denoiser", TaskType: "Audio"},
}

func GetAllProjects() []models.Project {
	return projects
}

func GetProjectByID(id string) (models.Project, error) {
	for _, b := range projects {
		if b.ID == id {
			return b, nil
		}
	}
	return models.Project{}, errors.New("project not found")
}

func CreateProject(newProject models.Project) (models.Project, error) {
	if err := validate.Struct(newProject); err != nil {
		return models.Project{}, err
	}

	for _, b := range projects {
		if b.ID == newProject.ID {
			return models.Project{}, errors.New("a project with this ID already exists")
		}
	}

	projects = append(projects, newProject)
	return newProject, nil
}

func UpdateProject(id string, updatedProject models.UpdateProjectInput) (models.Project, error) {
	for i, b := range projects {
		if b.ID == id {
			if updatedProject.Name != nil {
				projects[i].Name = *updatedProject.Name
			}

			return projects[i], nil
		}
	}
	return models.Project{}, errors.New("project not found")
}

func DeleteProject(id string) error {
	for i, b := range projects {
		if b.ID == id {
			projects = append(projects[:i], projects[i+1:]...)
			return nil
		}
	}
	return errors.New("project not found")
}
