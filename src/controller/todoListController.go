package controller

import (
	usecase "github.com/leolucena22/todo-list-API/UseCase"
	"github.com/leolucena22/todo-list-API/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskUseCase usecase.TaskUseCase
}

func NewTaskController(taskUseCase usecase.TaskUseCase) TaskController {
	return TaskController{
		taskUseCase: taskUseCase,
	}
}

func (p *TaskController) CreateTask(ctx *gin.Context) {
	var task model.Task
	err := ctx.BindJSON(&task)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedTask, err := p.taskUseCase.CreateTask(task)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedTask)
}

func (p *TaskController) GetTasks(ctx *gin.Context) {
	tasks, err := p.taskUseCase.GetTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (p *TaskController) GetTasksByStatus(ctx *gin.Context) {
	status := ctx.Query("status")
	tasks, err := p.taskUseCase.GetTasksByStatus(status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (p *TaskController) GetTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response := model.Response{
			Message: "ID do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)

		return
	}

	taskId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "ID do produto precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)

		return
	}

	task, err := p.taskUseCase.GetTaskByID(taskId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if task == nil {
		response := model.Response{
			Message: "Produto não encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, task)

}

func (p *TaskController) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{
			"ID é obrigatório"})
		return
	}

	taskId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "" +
			"ID inválido"})
		return
	}

	var request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "" +
			"Dados inválidos"})
		return
	}

	task, err := p.taskUseCase.UpdateTask(taskId, request.Title, request.Description)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "" +
			"Erro ao atualizar tarefa"})
		return
	}

	if task == nil {
		ctx.JSON(http.StatusNotFound, model.Response{Message: "" +
			"Tarefa não encontrada"})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (p *TaskController) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{
			"ID é obrigatório"})
		return
	}

	taskId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "" +
			"ID inválido"})
		return
	}

	task, err := p.taskUseCase.DeleteTask(taskId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "" +
			"Erro ao deletar tarefa"})
	}

	if task == nil {
		ctx.JSON(http.StatusNotFound, model.Response{Message: "" +
			"Tarefa não encontrada"})
	}

	ctx.JSON(http.StatusOK, task)

}
