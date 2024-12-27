package domain

import "context"

type GetHistoryRequest struct {
}

type GetHistoryResponse struct {
}

type GetHistoryHandler interface {
	Handle(ctx context.Context, req *GetHistoryRequest) (res *GetHistoryResponse, err error)
}
