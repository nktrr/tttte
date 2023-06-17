package dto

import uuid "github.com/satori/go.uuid"

type UpdateUserDto struct {
	UserID       uuid.UUID `json:"userID" bson:"id"`
	Email        string    `json:"email" bson:"email"`
	Name         string    `json:"name" bson:"name"`
	Surname      string    `json:"surname" bson:"surname"`
	Patronymic   string    `json:"patronymic" bson:"patronymic"`
	Role         string    `json:"role" bson:"role"`
	Phone        string    `json:"phone" bson:"phone"`
	Country      string    `json:"country" bson:"country"`
	Address      string    `json:"address" bson:"address"`
	Region       string    `json:"region" bson:"region"`
	JobPosition  string    `json:"jobPosition" bson:"jobPosition"`
	About        string    `json:"about" bson:"about"`
	City         string    `json:"city" bson:"city"`
	DepartmentID uuid.UUID `json:"departmentID" bson:"departmentID"`
	ChiefID      uuid.UUID `json:"chiefID" bson:"chiefID"`
}

type UpdateUserResponseDto struct {
	UserID uuid.UUID `json:"userID"`
}
