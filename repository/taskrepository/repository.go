package taskrepository

import (
	"hacktiv8-msib-final-project-3/entity"
	"hacktiv8-msib-final-project-3/pkg/errs"
)

type TaskRepository interface {
	CreateTask(user *entity.User, task *entity.Task) (*entity.Task, errs.MessageErr)
	GetAllTasks() ([]entity.Task, errs.MessageErr)
	GetAllTasksByCategoryID(categoryID uint) ([]entity.Task, errs.MessageErr)
	GetTaskByID(id uint) (*entity.Task, errs.MessageErr)
	UpdateTask(oldTask *entity.Task, newTask *entity.Task) (*entity.Task, errs.MessageErr)
	UpdateTaskStatus(id uint, newStatus bool) (*entity.Task, errs.MessageErr)
	UpdateTaskCategory(id uint, newCategoryID uint) (*entity.Task, errs.MessageErr)
	DeleteTask(id uint) errs.MessageErr
}
