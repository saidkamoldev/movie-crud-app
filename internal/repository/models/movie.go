package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title    string `json:"title"`
	Director string `json:"director"`
	Year     int    `json:"year"`
	Plot     string `json:"plot"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
