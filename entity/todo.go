package entity

// Todo is entity
type Todo struct {
	ID      int
	Title   string
	Content string
	Status  string
}

// Todos is slice of entity.Todo
type Todos []Todo
