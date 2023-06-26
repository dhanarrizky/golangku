package entities

import "time"

type Category struct {
	Id       uint
	Name     string
	CreateAt time.Time
	UpdateAt time.Time
}