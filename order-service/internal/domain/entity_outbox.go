package domain

type OutboxEntity struct {
	AggregateID string `json:"aggregateId" db:"AGGREGATE_ID"`
	CommandID   string `json:"commandId" db:"COMMAND_ID"`
	CommandType string `json:"commandType" db:"COMMAND_TYPE"`
	Payloay     string `json:"payloay" db:"PAYLOAD"`
	ReplyTo     string `json:"replyTo" db:"REPLY_TO"`
}
