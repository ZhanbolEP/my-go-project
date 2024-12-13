package models

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"` // Includes ID, created_at, and updated_at fields
	Email            string           `json:"email" bson:"email"`
	Name             string           `json:"name" bson:"name"`
	Password         string           `json:"-" bson:"password"` // Excluded from JSON output
	IsAdmin          bool             `json:"is_admin" bson:"is_admin"`
}
