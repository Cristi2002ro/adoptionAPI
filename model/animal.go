package model

import (
	"github.com/google/uuid"
)

type Animal struct {
	Id          uuid.UUID `json:"id"`
	CategoryId  string    `json:"categoryId"`
	BreedId     *string   `json:"breedId"`
	Name        *string   `json:"name"`
	Image       *string   `json:"image"`
	Age         *int      `json:"age"`
	Weight      *float64  `json:"weight"`
	Gender      *string   `json:"gender"`
	Adopted     *bool     `json:"adopted"`
	Reserved    *bool     `json:"reserved"`
	UserId      *string   `json:"userId"`
	Description *string   `json:"description"`
	Location    *string   `json:"location"`
}
