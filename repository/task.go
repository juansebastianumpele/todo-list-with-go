package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	taksData := []entity.Task{}
	err := r.db.WithContext(ctx).Model(&entity.Task{}).Where("user_id = ?", id).Scan(&taksData)
	if err.Error != nil {
		return []entity.Task{}, err.Error
	}
	return taksData, nil // TODO: replace this
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	res := r.db.WithContext(ctx).Create(&task)
	if res.Error != nil {
		return 0, res.Error
	}
	return task.ID, nil // TODO: replace this
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	taskData := entity.Task{}
	res := r.db.WithContext(ctx).Model(&entity.Task{}).Where("id", id).Find(&taskData)
	if res.Error != nil {
		return entity.Task{}, res.Error
	}
	return taskData, nil // TODO: replace this
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	taskData := []entity.Task{}
	res := r.db.WithContext(ctx).Model(&entity.Task{}).Where("category_id = ?", catId).Scan(&taskData)
	if res.Error != nil {
		return []entity.Task{}, res.Error
	}
	return taskData, nil // TODO: replace this
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	res := r.db.WithContext(ctx).Model(&entity.Task{}).Where("id = ?", task.ID).Updates(&task)
	if res.Error != nil {
		return res.Error
	}
	return nil // TODO: replace this
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	res := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Task{})
	if res.Error != nil {
		return res.Error
	}
	return nil // TODO: replace this
}
