package model

import (
	"github.com/google/uuid"
	"time"
)

type Animal struct {
	Id              uuid.UUID  `json:"id"`
	CategoryId      string     `json:"categoryId"`
	Name            *string    `json:"name"`
	Age             *int       `json:"age"`
	Species         *string    `json:"species"`
	Gender          *string    `json:"gender"`
	Weight          *float64   `json:"weight"`
	ReservationDate *time.Time `json:"reservationDate"`
	IsReserved      *bool      `json:"isReserved"`
	IsAdopted       *bool      `json:"isAdopted"`
	Image           *string    `json:"image"`
	Location        *string    `json:"location"`
}
