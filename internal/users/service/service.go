package service

import (
	"nktrr/pfizer/api_gateway_service/config"
	"nktrr/pfizer/api_gateway_service/internal/users/commands"
	"nktrr/pfizer/api_gateway_service/internal/users/queries"
	kafkaClient "nktrr/pfizer/pkg/kafka"
	"nktrr/pfizer/pkg/logger"
	readerService "nktrr/pfizer/reader_service/proto/product_reader"
)

type UserService struct {
	Commands *commands.UserCommands
	Queries  *queries.UserQueries
}

func NewUserService(log logger.Logger, cfg *config.Config, kafkaProducer kafkaClient.Producer, rsClient readerService.ReaderServiceClient) *UserService {
	createUserHandler := commands.NewCreateUserHandler(log, cfg, kafkaProducer)
	updateUserHandler := commands.NewUpdateProductHandler(log, cfg, kafkaProducer)

	userCommands := commands.NewUserCommands(createUserHandler, updateUserHandler)

	return &UserService{
		Commands: userCommands,
		Queries:  nil,
	}
}
