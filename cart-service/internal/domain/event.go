package domain

type Event[T any] struct {
	Payload T `json:"payload"`
}
