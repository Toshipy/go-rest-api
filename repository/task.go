package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
)

// Interfaceを作成する
type ITaskRepository interface {
	GetTasks(tasks *[]model.Task, userId uint) error
	GetTaskById(task *model.Task, userId uint, taskId uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

// Repositoryを作成する
type taskRepository struct {
	db *gorm.DB
}

// Factory関数（オブジェクト（構造体やクラス）のインスタンスを生成するための関数）
// 1.gorm.DB型のポインタを受け取る
// 2.その引数を使ってtaskRepository構造体のインスタンスを返す
// 3.そのインスタンスをITaskRepository型にキャストして返す
func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db: db}
}

// GetTasksを作成する
func (tr *taskRepository) GetTasks(tasks *[]model.Task, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).Order("created_at DESC").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

// GetTaskByIdを作成する
func (tr *taskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).Where("id = ?", taskId).First(task).Error; err != nil {
		return err
	}
	return nil
}

// CreateTaskを作成する
func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

// UpdateTaskを作成する
func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	result := tr.db.Joins("User").Where("user_id = ?", userId).Where("id = ?", taskId).Updates(task)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("task not found")
	}
	return nil
}

// DeleteTaskを作成する
func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	result := tr.db.Where("user_id = ?", userId).Where("id = ?", taskId).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("task not found")
	}
	return nil
}
