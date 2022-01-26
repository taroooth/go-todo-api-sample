package service

import (
	"fmt"
	"todo-api-sample/model"
)

type TodoService struct{}

func (TodoService) SetTodo(todo *model.Todo) error {
	_, err := db.Exec("INSERT INTO todo (title, content) VALUES (?, ?)", todo.Title, todo.Content)
	if err != nil {
		return err
	}
	return nil
}

func (TodoService) GetTodoList() []model.Todo {
	var todos []model.Todo

	rows, err := db.Query("SELECT * FROM todo")
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.Id, &todo.Title, &todo.Content); err != nil {
			return nil
		}
		fmt.Println(todo)
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		return nil
	}

	return todos
}

func (TodoService) UpdateTodo(id int, newTodo *model.Todo) error {
	_, err := db.Exec("UPDATE todo SET title = ?, content = ? WHERE id = ?", newTodo.Title, newTodo.Content, id)
	if err != nil {
		return err
	}
	return nil
}

func (TodoService) DeleteTodo(id int) error {
	_, err := db.Exec("DELETE FROM todo WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
