package repository

import (
	"database/sql"
	"fmt"

	"github.com/leolucena22/todo-list-API/model"
)

type TodoListRepository struct {
	connection *sql.DB
}

func NewTodolistRepository(connection *sql.DB) TodoListRepository {
	return TodoListRepository{
		connection: connection,
	}
}

func (pr *TodoListRepository) CreateTask(task model.Task) (int, error) {
	var id int
	query, err := pr.connection.Prepare(`
		INSERT INTO task (title, description) VALUES ($1, $2 RETURNING id)
`)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(task.Title, task.Description).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return id, nil
}

func (pr *TodoListRepository) GetTasksByStatus(status string) ([]model.Task, error) {
	query := `SELECT id, title, description, status, created_at 
              FROM tasks 
              WHERE status = $1 
              ORDER BY created_at DESC`

	rows, err := pr.connection.Query(query, status)
	if err != nil {
		fmt.Println(err.Error())
		return []model.Task{}, err
	}

	var taskList []model.Task
	var taskObj model.Task

	for rows.Next() {
		err = rows.Scan(
			&taskObj.ID,
			&taskObj.Title,
			&taskObj.Description)

		if err != nil {
			fmt.Println(err.Error())
			return []model.Task{}, err
		}

		taskList = append(taskList, taskObj)
	}

	rows.Close()

	return taskList, nil
}

func (pr *TodoListRepository) GetAllTasks() ([]model.Task, error) {
	query := `SELECT * FROM tasks;`

	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return []model.Task{}, err
	}

	var taskList []model.Task
	var taskObj model.Task

	for rows.Next() {
		err = rows.Scan(
			&taskObj.ID,
			&taskObj.Title,
			&taskObj.Description)

		if err != nil {
			fmt.Println(err.Error())
			return []model.Task{}, err
		}

		taskList = append(taskList, taskObj)
	}

	rows.Close()

	return taskList, nil
}

func (pr *TodoListRepository) GetTaskById(id int) (*model.Task, error) {
	query, err := pr.connection.Prepare("SELECT * FROM tasks WHERE id = $1")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var task model.Task

	err = query.QueryRow(id).Scan(&task.ID, &task.Title, &task.Description)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &task, nil
}

func (pr *TodoListRepository) UpdateTask(id int, newTitle string, newDescription string) (*model.Task, error) {
	query, err := pr.connection.Prepare(`UPDATE tasks SET title = $1, description = $2 WHERE id = $3
		RETURNING id, title, description`)
	if err != nil {
		return nil, fmt.Errorf("erro ao preparar query: %w", err)
	}
	defer query.Close()

	var task model.Task
	err = query.QueryRow(id).Scan(&task.ID, &task.Title, &task.Description)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("erro ao atualizar tarefa: %w", err)
	}

	return &task, nil
}

func (pr *TodoListRepository) DeleteTask(id int) (*model.Task, error) {
	query, err := pr.connection.Prepare(`DELETE FROM tasks WHERE id = $1 RETURNING id, title, description`)

	if err != nil {
		return nil, fmt.Errorf("Erro ao preparar query: %w", err)
	}
	defer query.Close()

	var task model.Task
	err = query.QueryRow(id).Scan(&task.ID, &task.Title, &task.Description)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("erro ao deletar tarefa: %w", err)
	}
	return &task, nil
}
