package domain

import (
	"github.com/google/uuid"
)

// TaskID はタスクを一意に識別する値（要件 5.1: UUID）。
type TaskID uuid.UUID

func NewTaskID() TaskID {
	return TaskID(uuid.New())
}

func TaskIDFromUUID(u uuid.UUID) TaskID {
	return TaskID(u)
}

func (id TaskID) UUID() uuid.UUID {
	return uuid.UUID(id)
}

func (id TaskID) String() string {
	return uuid.UUID(id).String()
}
