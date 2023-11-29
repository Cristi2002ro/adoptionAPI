package model

import "github.com/google/uuid"

type Favorite struct {
	Id       uuid.UUID `json:"id"`
	UserId   *string   `json:"userId"`
	AnimalId *string   `json:"animalId"`
}
