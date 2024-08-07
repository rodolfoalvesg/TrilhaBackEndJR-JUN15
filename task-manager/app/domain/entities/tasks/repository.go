package tasks

import (
	"context"

	"github.com/google/uuid"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task Task) error
	DeleteTask(ctx context.Context, taskID uuid.UUID) error
	GetTask(ctx context.Context, taskID uuid.UUID) (Task, error)
	GetTasks(ctx context.Context) ([]Task, error)
	UpdateTask(ctx context.Context, task Task) error
}
