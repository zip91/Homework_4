package model

type Status string

const (
	StatusNew       Status = "Новая"
	StatusInProcess Status = "В процессе"
	StatusDone      Status = "Завершена"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}
