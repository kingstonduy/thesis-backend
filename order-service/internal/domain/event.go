package domain

type Command[T any] struct {
	AggregateID string `json:"aggregateId"`
	CommandID   string `json:"commandId"`
	CommandType string `json:"commandType"`
	Payloay     T      `json:"payloay"`
	ReplyTo     string `json:"replyTo"`
}

const (
	EXECUTE_TRANSACTION_COMMAND = "EXECUTE_TRANSACTION_COMMAND"
)
