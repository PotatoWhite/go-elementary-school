package dto

type Simple struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}
