package dto

type Simple struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
