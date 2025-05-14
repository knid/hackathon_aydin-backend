package responses

type Task struct {
	Title string
	Description string
	Hour string
}

type TaskList struct {
	Tasks []Task
}
