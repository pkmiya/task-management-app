package usecase

import (
	"errors"

	"task-management-app/backend/internal/domain"
)

// TaskService は Phase 01 のタスク関連ユースケースをまとめる。
type TaskService struct {
	Repo  TaskRepository
	Clock Clock
}

func (s *TaskService) CreateTask(rawTitle string) (*domain.Task, error) {
	title, err := domain.ParseTitle(rawTitle)
	if err != nil {
		return nil, err
	}
	id := domain.NewTaskID()
	task := domain.NewTask(id, title, s.Clock.Now())
	if err := s.Repo.Save(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) ListTasks() ([]*domain.Task, error) {
	return s.Repo.FindAll()
}

func (s *TaskService) GetTask(id domain.TaskID) (*domain.Task, error) {
	task, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) UpdateTaskTitle(id domain.TaskID, rawTitle string) (*domain.Task, error) {
	task, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	title, err := domain.ParseTitle(rawTitle)
	if err != nil {
		return nil, err
	}
	task.Rename(title)
	if err := s.Repo.Save(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) DeleteTask(id domain.TaskID) error {
	err := s.Repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// IsNotFound はリポジトリ／ユースケースが返した NotFound を判定する。
func IsNotFound(err error) bool {
	return errors.Is(err, domain.ErrNotFound)
}

// IsValidation は入力検証エラーを判定する。
func IsValidation(err error) bool {
	return errors.Is(err, domain.ErrValidation)
}
