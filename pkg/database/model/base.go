package model

type QueryResult[T any] struct {
	Result []T     `json:"result"`
	Status *string `json:"status"`
	Time   *string `json:"time"`
}
