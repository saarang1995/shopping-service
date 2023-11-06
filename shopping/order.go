package shopping

import "time"

type Order struct {
	Id *uint
	Name string
	UserId *uint
	Amount uint64
	Quantity uint
	CreatedAt time.Time
}