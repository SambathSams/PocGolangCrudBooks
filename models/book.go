package models

type Book struct {
	ID     string  `json:"id" validate:"required,gt=0" example:"2"`
	Title  string  `json:"title" validate:"required,min=2" example:"The Power of your subconscious mind"`
	Author string  `json:"author" validate:"required,min=2" example:"Joseph Murphy"`
	Price  float64 `json:"price" validate:"required,gt=0" example:"11.99"`
}

type UpdateBookInput struct {
	Title  *string  `json:"title" validate:"omitempty,min=2"`
	Author *string  `json:"author" validate:"omitempty,min=2"`
	Price  *float64 `json:"price" validate:"omitempty,gt=0"`
}
