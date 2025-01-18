package domain

type Event[T any] struct {
	Payload EventPayload[T] `json:"payload"`
}

type EventPayload[T any] struct {
	Before T      `json:"before"`
	After  T      `json:"after"`
	Op     string `json:"op"`
	Ts     int64  `json:"ts_ms"`
}
