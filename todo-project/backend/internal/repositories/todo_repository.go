package repositories

import (
	"database/sql"

	"todo-project-backend/internal/api/models"
	"todo-project-backend/internal/database"

	_ "github.com/lib/pq"
)

type TodoRepository struct {
}

func NewTodoRepository() (*TodoRepository, error) {
	return &TodoRepository{}, nil
}

func (tr *TodoRepository) GetAll() ([]models.Todo, error) {
	todos := []models.Todo{}

	rows, err := database.Database.Query("SELECT id, title, completed FROM todos")
	if err == sql.ErrNoRows {
		return todos, nil
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Completed)
		if err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}
	err = rows.Err()
	if err != nil {
		return todos, err
	}

	return todos, nil
}

func (tr *TodoRepository) Create(newTodo models.NewTodo) (models.Todo, error) {
	var addedTodo models.Todo
	err := database.Database.QueryRow("INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id, title, completed",
		newTodo.Title, false).Scan(&addedTodo.Id, &addedTodo.Title, &addedTodo.Completed)

	return addedTodo, err
}
