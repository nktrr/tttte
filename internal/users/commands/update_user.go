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

type UpdateUserCmdHandler interface {
	Handle(ctx context.Context, command *UpdateUserCommand) error
}

type updateUserCmdHandler struct {
	log           logger.Logger
	cfg           *config.Config
	kafkaProducer kafkaClient.Producer
}

func NewUpdateProductHandler(log logger.Logger, cfg *config.Config, kafkaProducer kafkaClient.Producer) *updateUserCmdHandler {
	return &updateUserCmdHandler{log: log, cfg: cfg, kafkaProducer: kafkaProducer}
}

func (c *updateUserCmdHandler) Handle(ctx context.Context, command *UpdateUserCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "updateProductCmdHandler.Handle")
	defer span.Finish()

	updateDto := &kafkaMessages.UserUpdate{
		UserID:       command.UpdateDto.UserID.String(),
		Email:        command.UpdateDto.Email,
		Name:         command.UpdateDto.Name,
		Surname:      command.UpdateDto.Surname,
		Patronymic:   command.UpdateDto.Patronymic,
		Role:         command.UpdateDto.Role,
		Phone:        command.UpdateDto.Phone,
		Country:      command.UpdateDto.Country,
		Address:      command.UpdateDto.Address,
		Region:       command.UpdateDto.Region,
		JobPosition:  command.UpdateDto.JobPosition,
		About:        command.UpdateDto.About,
		City:         command.UpdateDto.City,
		DepartmentID: command.UpdateDto.DepartmentID.String(),
		ChiefID:      command.UpdateDto.ChiefID.String(),
	}

	dtoBytes, err := proto.Marshal(updateDto)
	if err != nil {
		return err
	}

	return c.kafkaProducer.PublishMessage(ctx, kafka.Message{
		Topic:   c.cfg.KafkaTopics.ProductUpdate.TopicName,
		Value:   dtoBytes,
		Time:    time.Now().UTC(),
		Headers: tracing.GetKafkaTracingHeadersFromSpanCtx(span.Context()),
	})
}
