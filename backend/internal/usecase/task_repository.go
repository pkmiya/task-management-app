package usecase

import "task-management-app/backend/internal/domain"

// TaskRepository はユースケースから見た永続化の抽象（要件 5.6）。
// インタフェースは利用側である usecase パッケージに置く。
type TaskRepository interface {
	Save(task *domain.Task) error
	FindByID(id domain.TaskID) (*domain.Task, error)
	FindAll() ([]*domain.Task, error)
	Delete(id domain.TaskID) error
}
