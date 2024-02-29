package entity

import "time"

type Customer struct {
	Id, Balance          int32
	Name, Email, Address string
	Rating               float64
	Married              bool
	BirthDate            time.Time
}
