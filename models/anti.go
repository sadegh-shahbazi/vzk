package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
	"time"
)

type Anti struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Name        string    `json:"name" db:"name"`
	NameFa      string    `json:"name_fa" db:"name_fa"`
	TypeName    string    `json:"type_name" db:"type_name"`
	TypeNameFa  string    `json:"type_name_fa" db:"type_name_fa"`
	Description string    `json:"description" db:"description"`
	Image       string    `json:"image" db:"image"`
	AntiOrder   int       `json:"anti_order" db:"anti_order"`
	Price       int       `json:"price" db:"price"`
	PeriodDay   int       `json:"period_day" db:"period_day"`
}

// String is not required by pop and may be deleted
func (a Anti) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Antis is not required by pop and may be deleted
type Antis []Anti

// String is not required by pop and may be deleted
func (a Antis) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *Anti) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: a.Name, Name: "Name"},
		&validators.StringIsPresent{Field: a.NameFa, Name: "NameFa"},
		&validators.StringIsPresent{Field: a.TypeName, Name: "TypeName"},
		&validators.StringIsPresent{Field: a.TypeNameFa, Name: "TypeNameFa"},
		&validators.StringIsPresent{Field: a.Description, Name: "Description"},
		&validators.StringIsPresent{Field: a.Image, Name: "Image"},
		&validators.IntIsPresent{Field: a.AntiOrder, Name: "AntiOrder"},
		&validators.IntIsPresent{Field: a.Price, Name: "Price"},
		&validators.IntIsPresent{Field: a.PeriodDay, Name: "PeriodDay"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *Anti) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *Anti) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
