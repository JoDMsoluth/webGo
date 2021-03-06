package model

import (
	"time"
)


type Todo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type DBHandler interface {
	GetTodo(sessionId string) []*Todo
	AddTodo(name string, sessionId string) *Todo 
	RemoveTodo(id int) bool
	CompleteTodo(id int, complete bool) bool
	Close() 
}

type memoryHandler struct {
	todoMap map[int]*Todo
}


// var handler DBHandler

// func init() {
// 	// handler = newMemoryHandler()
// 	handler = newSqliteHandler()
	
// }

func NewDBHandler(filepath string) DBHandler {
	return newSqliteHandler(filepath)
}
