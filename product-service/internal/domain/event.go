package domain

import "time"

type Event[T any] struct {
	Payload T `json:"payload"`
}

type EventPayload[T any] struct {
	Before T         `json:"before"`
	After  T         `json:"after"`
	Op     string    `json:"op"`
	Ts     time.Time `json:"ts"`
}
