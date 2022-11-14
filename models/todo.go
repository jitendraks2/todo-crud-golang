package models

type Todos struct {
	Id          string `json:"id"`
	Todo        string `json:"todo"`
	IsCompleted bool   `json:"completed"`
}
