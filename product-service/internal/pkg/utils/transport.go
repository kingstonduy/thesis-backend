package utils

import (
	"context"
	"net/http"
	"time"

	"github.com/kingstonduy/go-core/errorx"
	"github.com/kingstonduy/go-core/transport"
)

// response trace là saga cầm cái trace trả về thằng gọi nó (vd: napas-fast-fund-247)
//
// original trace là cái trace đầu tiên saga nhận được từ service napas-fast-fund-247
//
// vì original trace đã được lưu ở bảng transaction nên chỉ cần query lấy nó ra và tạo response trace
func GenResponseTrace(originalTrace transport.Trace) transport.Trace {
	now := time.Now()

	trace := transport.Trace{
		From:               originalTrace.To,
		To:                 originalTrace.From,
		Cid:                originalTrace.Cid,
		Sid:                originalTrace.Sid,
		Cts:                originalTrace.Cts,
		Sts:                now.UnixMilli(),
		Dur:                now.UnixMilli() - originalTrace.Cts,
		Username:           originalTrace.Username,
		ReplyTo:            originalTrace.ReplyTo,
		TransactionTimeout: originalTrace.TransactionTimeout,
	}

	return trace
}

// request trace nghĩa là saga cầm cái trace này đi command các serivce khác (vd: fund-transfer, curent-account, ...)
//
// original trace là cái trace đầu tiên saga nhận được từ service napas-fast-fund-247
//
// vì original trace đã được lưu ở bảng transaction nên chỉ cần query lấy nó ra và tạo request trace
func GenRequestTrace(originalTrace transport.Trace, to string, replyTo string) transport.Trace {
	now := time.Now()

	trace := transport.Trace{
		From:               "asset-mgmt",
		To:                 to,
		Cid:                originalTrace.Cid,
		Sid:                originalTrace.Sid,
		Cts:                now.UnixMilli(),
		Sts:                0,
		Dur:                0,
		Username:           originalTrace.Username,
		ReplyTo:            replyTo,
		TransactionTimeout: originalTrace.TransactionTimeout,
	}

	return trace
}

// cung cấp 1 errorx trả về 1 result tương ứng vs errx đó. nếu errorx là nil thì trả về defaulf success
func GenResultFromErrorx(ctx context.Context, errx *errorx.Error) (result *transport.Result) {
	if errx != nil {
		result = &transport.Result{
			StatusCode: errx.Status,
			Code:       errx.Code,
			Message:    errx.Message,
			Details:    errx.Details,
		}
	} else {
		result = &transport.Result{
			StatusCode: http.StatusOK,
			Code:       "00",
			Message:    "success",
		}
	}

	return result
}
