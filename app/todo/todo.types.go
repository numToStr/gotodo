package todo

// Todo defines the todo item
type Todo struct {
	Title string `json:"title"`
}

// Todos defines the todos list
type Todos struct {
	Todos []Todo `json:"todos"`
}
