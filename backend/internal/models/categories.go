package models

type Category struct {
	CategoryID int    `json:"category_id"`
	UserID     int    `json:"user_id"`
	Name       string `json:"name"`
}
