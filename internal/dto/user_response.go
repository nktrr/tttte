package dto

import userService "nktrr/pfizer/user_service/proto/user_reader"

type UserResponse struct {
	UserID       string `json:"userID" bson:"id"`
	Email        string `json:"email" bson:"email"`
	Name         string `json:"name" bson:"name"`
	Surname      string `json:"surname" bson:"surname"`
	Patronymic   string `json:"patronymic" bson:"patronymic"`
	Role         string `json:"role" bson:"role"`
	Phone        string `json:"phone" bson:"phone"`
	Country      string `json:"country" bson:"country"`
	Address      string `json:"address" bson:"address"`
	Region       string `json:"region" bson:"region"`
	JobPosition  string `json:"jobPosition" bson:"jobPosition"`
	About        string `json:"about" bson:"about"`
	City         string `json:"city" bson:"city"`
	DepartmentID string `json:"departmentID" bson:"departmentID"`
	ChiefID      string `json:"chiefID" bson:"chiefID"`
}

func UserResponseFromGrpc(user *userService.User) *UserResponse {
	return &UserResponse{
		UserID:       user.GetUserID(),
		Email:        user.GetEmail(),
		Name:         user.GetName(),
		Surname:      user.GetSurname(),
		Patronymic:   user.GetPatronymic(),
		Role:         user.GetRole(),
		Phone:        user.GetPhone(),
		Country:      user.GetCountry(),
		Address:      user.GetAddress(),
		Region:       user.GetRegion(),
		JobPosition:  user.GetJobPosition(),
		About:        user.GetAbout(),
		City:         user.GetCity(),
		DepartmentID: user.GetDepartmentID(),
		ChiefID:      user.GetChiefID(),
	}
}
