package model

import "time"

type Movie struct {
	ID          int64
	Name        string
	Url         string
	Description string
	KpRating    float64
	ImdbRating  float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Chat struct {
	ID       int64
	Name     string
	MemberID int64
}

//type Movie struct {
//	Name string
//	Rating string
//	Url string
//	ChatID int64
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	ViewedAt time.Time
//}
