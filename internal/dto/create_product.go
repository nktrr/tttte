package dto

import uuid "github.com/satori/go.uuid"

type CreateProductDto struct {
	ProductID   uuid.UUID `json:"productId" `
	Name        string    `json:"name"`
	Description string    `json:"description" `
	Price       float64   `json:"price"`
}

type CreateProductResponseDto struct {
	ProductID uuid.UUID `json:"productId" validate:"required"`
}
