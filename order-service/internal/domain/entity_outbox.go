package domain

import "context"

type OutboxEntity struct {
	AggregateID string `json:"AGGREGATE_ID" db:"AGGREGATE_ID"`
	CommandID   string `json:"COMMAND_ID" db:"COMMAND_ID"`
	CommandType string `json:"COMMAND_TYPE" db:"COMMAND_TYPE"`
	Payloay     string `json:"PAYLOAD" db:"PAYLOAD"`
	ReplyTo     string `json:"REPLY_TO" db:"REPLY_TO"`
}

type IOutboxRepo interface {
	Insert(ctx context.Context, outbox OutboxEntity) error
}
