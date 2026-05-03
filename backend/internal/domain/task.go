package domain

import "time"

// Task は Phase 01 におけるタスク集約ルート（ステータスなし）。
type Task struct {
	id        TaskID
	title     Title
	createdAt time.Time
}

func NewTask(id TaskID, title Title, createdAt time.Time) *Task {
	return &Task{
		id:        id,
		title:     title,
		createdAt: createdAt.UTC(),
	}
}

func (t *Task) ID() TaskID {
	return t.id
}

func (t *Task) Title() Title {
	return t.title
}

func (t *Task) CreatedAt() time.Time {
	return t.createdAt
}

// Rename は検証済みの Title でタイトルを置き換える（UC-04 のドメイン操作）。
func (t *Task) Rename(title Title) {
	t.title = title
}
