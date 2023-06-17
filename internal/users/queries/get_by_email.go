package queries

import (
	"context"
	"nktrr/pfizer/api_gateway_service/internal/dto"
)

type GetUserByEmailHandler interface {
	Handle(ctx context.Context, query *GetUserByEmailQuery) (*dto.UserResponse, error)
}
