package usecase

import (
	"fmt"
	"github.com/leolucena22/todo-list-API/model"
	"github.com/leolucena22/todo-list-API/repository"

	_ "github.com/lib/pq"
)

type TaskUseCase struct {
	repository repository.TodoListRepository
}

func NewTaskUsecase(repo repository.TodoListRepository) TaskUseCase {
	return TaskUseCase{
		repository: repo,
	}
}

func (pu *TaskUseCase) GetTasksByStatus(status string) ([]model.Task, error) {
	if status != "" && status != "pending" && status != "completed" {
		return nil, fmt.Errorf("status inválido: deve ser 'pending', 'completed' ou vazio")
	}

	// Chamada ao repositório
	tasks, err := pu.repository.GetTasksByStatus(status)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar tarefas: %w", err)
	}

	// Verifica se há resultados
	if len(tasks) == 0 {
		if status != "" {
			return nil, fmt.Errorf("nenhuma tarefa encontrada com status '%s'", status)
		}
		return nil, fmt.Errorf("nenhuma tarefa cadastrada")
	}

	return tasks, nil
}

func (pu *TaskUseCase) GetTasks() ([]model.Task, error) {
	return pu.GetTasks()
}

func (pu *TaskUseCase) GetTaskByID(id int) (*model.Task, error) {
	task, err := pu.repository.GetTaskById(id)

	if err != nil {
		return nil, fmt.Errorf("erro ao buscar tarefas: %w", err)
	}

	return task, nil
}

func (pu *TaskUseCase) CreateTask(task model.Task) (model.Task, error) {
	id, err := pu.repository.CreateTask(task)

	if err != nil {
		return model.Task{}, err
	}

	task.ID = id

	return task, nil
}

func (pu *TaskUseCase) UpdateTask(id int, title string, description string) (*model.Task, error) {
	return pu.repository.UpdateTask(id, title, description)
}

func (pu *TaskUseCase) DeleteTask(id int) (*model.Task, error) {
	return pu.repository.DeleteTask(id)
}
