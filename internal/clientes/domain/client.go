package domain

type Client struct {
	ID    int    `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}
