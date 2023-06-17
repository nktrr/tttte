package server

import (
	"context"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"nktrr/pfizer/api_gateway_service/config"
	"nktrr/pfizer/api_gateway_service/internal/client"
	"nktrr/pfizer/api_gateway_service/internal/metrics"
	"nktrr/pfizer/api_gateway_service/internal/middlewares"
	v1 "nktrr/pfizer/api_gateway_service/internal/users/delivery/http/v1"
	"nktrr/pfizer/api_gateway_service/internal/users/service"
	"nktrr/pfizer/pkg/interceptors"
	"nktrr/pfizer/pkg/kafka"
	"nktrr/pfizer/pkg/logger"
	"nktrr/pfizer/pkg/tracing"
	readerService "nktrr/pfizer/reader_service/proto/product_reader"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	log  logger.Logger
	cfg  *config.Config
	v    *validator.Validate
	mw   middlewares.MiddlewareManager
	im   interceptors.InterceptorManager
	echo *echo.Echo
	us   *service.UserService
	m    *metrics.ApiGatewayMetrics
}

func NewServer(log logger.Logger, cfg *config.Config) *server {
	return &server{log: log, cfg: cfg, echo: echo.New(), v: validator.New()}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	s.mw = middlewares.NewMiddlewareManager(s.log, s.cfg)
	s.im = interceptors.NewInterceptorManager(s.log)
	s.m = metrics.NewApiGatewayMetrics(s.cfg)

	readerServiceConn, err := client.NewReaderServiceConn(ctx, s.cfg, s.im)
	if err != nil {
		return err
	}
	defer readerServiceConn.Close() // nolint: errcheck
	rsClient := readerService.NewReaderServiceClient(readerServiceConn)

	kafkaProducer := kafka.NewProducer(s.log, s.cfg.Kafka.Brokers)
	defer kafkaProducer.Close() // nolint: errcheck

	s.us = service.NewUserService(s.log, s.cfg, kafkaProducer, rsClient)

	usersHandlers := v1.NewUsersHandlers(s.echo.Group(s.cfg.Http.UsersPath), s.log, s.mw, s.cfg, s.us, s.v, s.m)
	usersHandlers.MapRoutes()

	go func() {
		if err := s.runHttpServer(); err != nil {
			s.log.Errorf(" s.runHttpServer: %v", err)
			cancel()
		}
	}()
	s.log.Infof("API Gateway is listening on PORT: %s", s.cfg.Http.Port)

	s.runMetrics(cancel)
	s.runHealthCheck(ctx)

	if s.cfg.Jaeger.Enable {
		tracer, closer, err := tracing.NewJaegerTracer(s.cfg.Jaeger)
		if err != nil {
			return err
		}
		defer closer.Close() // nolint: errcheck
		opentracing.SetGlobalTracer(tracer)
	}

	<-ctx.Done()
	if err := s.echo.Server.Shutdown(ctx); err != nil {
		s.log.WarnMsg("echo.Server.Shutdown", err)
	}

	return nil
}
