package handlers

import (
	"errors"
	"github.com/cherrycutter/todo_app/internal/models"
	"github.com/cherrycutter/todo_app/internal/repos"
	"github.com/cherrycutter/todo_app/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TodoHandler struct {
	service services.TodoService
}

func NewTodoHandler(service services.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/todos", h.GetTodos)
	router.GET("/todo/:id", h.GetTodo)
	router.POST("/todo", h.PostTodo)
	router.PATCH("/todo/:id", h.UpdateTodo)
	router.DELETE("/todo/:id", h.DeleteTodo)
}

// GetTodos godoc
// @Summary Get all todos
// @Description Returns a list of all todos
// @Tags todos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.TodoModel
// @Failure 500 {object} errorResponse
// @Router /todos [get]
func (h *TodoHandler) GetTodos(ctx *gin.Context) {
	todos, err := h.service.GetTodos(ctx.Request.Context())
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, "failed to get todos")
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

// GetTodo godoc
// @Summary Get todo by ID
// @Description Returns one todo by id
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} models.TodoModel
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /todos/{id} [get]
func (h *TodoHandler) GetTodo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}
	todo, err := h.service.GetTodo(ctx.Request.Context(), id)
	if err != nil {
		if errors.Is(err, repos.ErrTodoNotFound) {
			newErrorResponse(ctx, http.StatusNotFound, "todo not found")
		} else {
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

// PostTodo godoc
// @Summary Create a new todo
// @Description Creates one new todo
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.TodoModel true "Todo Model"
// @Success 201 {object} models.TodoModel
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /todos [post]
func (h *TodoHandler) PostTodo(ctx *gin.Context) {
	var todo models.TodoModel
	if err := ctx.BindJSON(&todo); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	createdTodo, err := h.service.CreateTodo(ctx.Request.Context(), todo)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, createdTodo)
}

// UpdateTodo godoc
// @Summary Update an existing todo
// @Description Updates an existing todo by id
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body models.TodoModel true "Todo Model"
// @Success 200 {object} models.TodoModel
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /todos/{id} [patch]
func (h *TodoHandler) UpdateTodo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}
	var todo models.TodoModel
	if err = ctx.BindJSON(&todo); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	updatedTodo, err := h.service.UpdateTodo(ctx.Request.Context(), id, todo)
	if err != nil {
		if errors.Is(err, repos.ErrTodoNotFound) {
			newErrorResponse(ctx, http.StatusNotFound, "todo not found")
		} else {
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}
	ctx.JSON(http.StatusOK, updatedTodo)
}

// DeleteTodo godoc
// @Summary Delete todo by ID
// @Description Deletes an existing todo by id
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} errorResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}
	err = h.service.DeleteTodo(ctx.Request.Context(), id)
	if err != nil {
		if errors.Is(err, repos.ErrTodoNotFound) {
			newErrorResponse(ctx, http.StatusNotFound, "todo not found")
		} else {
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "todo deleted successfully"})
}
