package model

import "time"

type Animal struct {
	Id              string
	CategoryId      string
	Name            string
	ShelterId       string
	Age             int
	Species         string
	Gender          string
	Weight          float64
	ReservationDate time.Time
	IsReserved      bool
	Location        string
}
