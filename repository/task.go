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
	var tasks []entity.Task
	err := r.db.WithContext(ctx).Model(&entity.Task{}).Where("user_id = ?", id).Find(&tasks).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []entity.Task{}, nil
		}
		return nil, ctx.Err()
	}
	return tasks, nil // TODO: replace this
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	err = r.db.WithContext(ctx).Model(&entity.Task{}).Create(&task).Error
	if err != nil {
		return 0, err
	}
	taskId = task.ID
	return taskId, nil // TODO: replace this
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	var task entity.Task
	err := r.db.WithContext(ctx).Model(&entity.Task{}).Where("id = ?", id).Find(&task).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return entity.Task{}, nil
		}
		return entity.Task{}, err
	}
	return task, nil // TODO: replace this
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	var tasks []entity.Task
	err := r.db.WithContext(ctx).Model(&entity.Task{}).Where("category_id = ?", catId).Find(&tasks).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return []entity.Task{}, nil
		}
		return nil, ctx.Err()
	}
	return tasks, nil // TODO: replace this
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	err := r.db.WithContext(ctx).Model(&entity.Task{}).Where("id = ?", task.ID).Updates(&task).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	err := r.db.WithContext(ctx).Model(&entity.Task{}).Delete(&entity.Task{}, id).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}
