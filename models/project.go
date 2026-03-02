package models

type Project struct {
	ID       string `json:"id" validate:"required,gt=0" example:"2"`
	Name     string `json:"name" validate:"required,min=2" example:"Project2"`
	Task     string `json:"task" validate:"required,min=2" example:"Keyword Spotting"`
	TaskType string `json:"taskType" validate:"required,min=2" example:"Audio"`
}

type UpdateProjectInput struct {
	Name *string `json:"name" validate:"omitempty,min=2"`
}
