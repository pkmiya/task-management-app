package memory

import (
	"sort"

	"task-management-app/backend/internal/domain"
)

// TaskRepository は単一ワーカー前提のインメモリ実装（要件 6・Phase 01）。
type TaskRepository struct {
	tasks map[string]*domain.Task
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[string]*domain.Task),
	}
}

func (r *TaskRepository) Save(task *domain.Task) error {
	cp := domain.NewTask(task.ID(), task.Title(), task.CreatedAt())
	r.tasks[task.ID().String()] = cp
	return nil
}

func (r *TaskRepository) FindByID(id domain.TaskID) (*domain.Task, error) {
	t, ok := r.tasks[id.String()]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return cloneTask(t), nil
}

func (r *TaskRepository) FindAll() ([]*domain.Task, error) {
	out := make([]*domain.Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		out = append(out, cloneTask(t))
	}
	sort.SliceStable(out, func(i, j int) bool {
		ci, cj := out[i].CreatedAt(), out[j].CreatedAt()
		switch {
		case ci.Before(cj):
			return true
		case ci.After(cj):
			return false
		default:
			return out[i].ID().String() < out[j].ID().String()
		}
	})
	return out, nil
}

func (r *TaskRepository) Delete(id domain.TaskID) error {
	key := id.String()
	if _, ok := r.tasks[key]; !ok {
		return domain.ErrNotFound
	}
	delete(r.tasks, key)
	return nil
}

func cloneTask(t *domain.Task) *domain.Task {
	return domain.NewTask(t.ID(), t.Title(), t.CreatedAt())
}
