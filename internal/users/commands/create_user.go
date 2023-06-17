package commands

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"nktrr/pfizer/api_gateway_service/config"
	kafkaClient "nktrr/pfizer/pkg/kafka"
	"nktrr/pfizer/pkg/logger"
	"nktrr/pfizer/pkg/tracing"
	kafkaMessages "nktrr/pfizer/proto/kafka"
	"time"
)

type CreateUserCmdHandler interface {
	Handle(ctx context.Context, command *CreateUserCommand) error
}

type createUserHandler struct {
	log           logger.Logger
	cfg           *config.Config
	kafkaProducer kafkaClient.Producer
}

func NewCreateUserHandler(log logger.Logger, cfg *config.Config, kafkaProducer kafkaClient.Producer) *createUserHandler {
	return &createUserHandler{
		log:           log,
		cfg:           cfg,
		kafkaProducer: kafkaProducer,
	}
}

func (c *createUserHandler) Handle(ctx context.Context, command *CreateUserCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "createUserHandler.Handle")
	defer span.Finish()
	createDto := &kafkaMessages.CreateUserReq{
		UserID:       command.CreateDto.UserID.String(),
		Email:        command.CreateDto.Email,
		Name:         command.CreateDto.Name,
		Surname:      command.CreateDto.Surname,
		Patronymic:   command.CreateDto.Patronymic,
		Role:         command.CreateDto.Role,
		Phone:        command.CreateDto.Phone,
		Country:      command.CreateDto.Country,
		Address:      command.CreateDto.Address,
		Region:       command.CreateDto.Region,
		JobPosition:  command.CreateDto.JobPosition,
		About:        command.CreateDto.About,
		City:         command.CreateDto.City,
		DepartmentID: command.CreateDto.DepartmentID.String(),
		ChiefID:      command.CreateDto.ChiefID.String(),
	}

	dtoBytes, err := proto.Marshal(createDto)
	if err != nil {
		return err
	}
	return c.kafkaProducer.PublishMessage(ctx, kafka.Message{
		Topic:   c.cfg.KafkaTopics.UserCreate.TopicName,
		Value:   dtoBytes,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	})
}
