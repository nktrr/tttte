package v1

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"net/http"
	"nktrr/pfizer/api_gateway_service/config"
	"nktrr/pfizer/api_gateway_service/internal/dto"
	"nktrr/pfizer/api_gateway_service/internal/metrics"
	"nktrr/pfizer/api_gateway_service/internal/middlewares"
	"nktrr/pfizer/api_gateway_service/internal/users/commands"
	"nktrr/pfizer/api_gateway_service/internal/users/service"
	httpErrors "nktrr/pfizer/pkg/http_errors"
	"nktrr/pfizer/pkg/logger"
	"nktrr/pfizer/pkg/tracing"
)

type usersHandlers struct {
	group   *echo.Group
	log     logger.Logger
	mw      middlewares.MiddlewareManager
	cfg     *config.Config
	ps      *service.UserService
	v       *validator.Validate
	metrics *metrics.ApiGatewayMetrics
}

func NewUsersHandlers(
	group *echo.Group,
	log logger.Logger,
	mw middlewares.MiddlewareManager,
	cfg *config.Config,
	ps *service.UserService,
	v *validator.Validate,
	metrics *metrics.ApiGatewayMetrics,
) *usersHandlers {
	return &usersHandlers{group: group, log: log, mw: mw, cfg: cfg, ps: ps, v: v, metrics: metrics}
}

// CreateUser
// @Tags Users
// @Summary Create users
// @Description Create new user item
// @Accept json
// @Produce json
// @Success 201 {object} dto.CreateUserResponseDto
// @Router /users [post]
func (h *usersHandlers) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.metrics.CreateUserHttpRequests.Inc()

		ctx, span := tracing.StartHttpServerTracerSpan(c, "handlers.CreateUser")
		defer span.Finish()

		createDto := &dto.CreateUserDto{}
		if err := c.Bind(createDto); err != nil {
			h.log.WarnMsg("Bind", err)
			h.traceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		if err := h.v.StructCtx(ctx, createDto); err != nil {
			h.log.WarnMsg("validate", err)
			h.traceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		if err := h.ps.Commands.CreateUser.Handle(ctx, commands.NewCreateUserCommand(createDto)); err != nil {
			h.log.WarnMsg("CreateUser", err)
			h.metrics.ErrorHttpRequests.Inc()
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		h.metrics.SuccessHttpRequests.Inc()
		return c.JSON(http.StatusCreated, dto.CreateUserResponseDto{UserID: createDto.UserID})
	}
}

func (h *usersHandlers) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
	h.metrics.ErrorHttpRequests.Inc()
}
