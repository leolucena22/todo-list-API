package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	usecase "github.com/leolucena22/todo-list-API/UseCase"
	"github.com/leolucena22/todo-list-API/controller"
	"github.com/leolucena22/todo-list-API/db"
	"github.com/leolucena22/todo-list-API/repository"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	TaskRepository := repository.NewTodolistRepository(dbConnection)

	TaskUseCase := usecase.NewTaskUsecase(TaskRepository)

	TaskController := controller.NewTaskController(TaskUseCase)

	server.Use(cors.Default())

	/*
		POST /tasks - Criar nova tarefa
		GET /tasks - Listar todas tarefas (com filtro por status)
		GET /tasks/:id - Buscar tarefa por ID
		PUT /tasks/:id - Atualizar tarefa
		DELETE /tasks/:id - Remover tarefa
	*/

	server.POST("/tasks", TaskController.CreateTask)
	server.GET("/tasks", TaskController.GetTasksByStatus)
	server.GET("/tasksAll", TaskController.GetTasks)
	server.GET("/tasks/:id", TaskController.GetTaskById)
	server.PUT("/tasks/:id", TaskController.UpdateTask)
	server.DELETE("/tasks/:id", TaskController.DeleteTask)

	server.Run(":8000")
}
