package commands

import (
	uuid "github.com/satori/go.uuid"
	"nktrr/pfizer/api_gateway_service/internal/dto"
)

type UserCommands struct {
	CreateUser CreateUserCmdHandler
	UpdateUser UpdateUserCmdHandler
}

func NewUserCommands(createUser CreateUserCmdHandler, updateUser UpdateUserCmdHandler) *UserCommands {
	return &UserCommands{CreateUser: createUser, UpdateUser: updateUser}
}

func NewCreateUserCommand(createDto *dto.CreateUserDto) *CreateUserCommand {
	return &CreateUserCommand{CreateDto: createDto}
}

type CreateUserCommand struct {
	CreateDto *dto.CreateUserDto
}

type UpdateUserCommand struct {
	UpdateDto *dto.UpdateUserDto
}

type DeleteUserCommand struct {
	UserID uuid.UUID `json:"userID"`
}

type CreateUserCOmmand struct {
	CreateDto *dto.CreateUserDto
}
