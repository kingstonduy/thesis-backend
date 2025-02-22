package domain

import "context"

type WriteOutboxEntity struct {
	AggregateID string `json:"AGGREGATE_ID" db:"AGGREGATE_ID"`
	CommandID   string `json:"COMMAND_ID" db:"COMMAND_ID"`
	CommandType string `json:"COMMAND_TYPE" db:"COMMAND_TYPE"`
	Payload     string `json:"PAYLOAD" db:"PAYLOAD"`
	Trace       string `json:"TRACE" db:"TRACE"`
	ReplyTo     string `json:"REPLY_TO" db:"REPLY_TO"`
	TraceParent string `json:"TRACE_PARENT" db:"TRACE_PARENT"`
}

type IOutboxRepo interface {
	Insert(ctx context.Context, outbox WriteOutboxEntity) error
}
