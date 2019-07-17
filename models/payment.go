package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Payment struct {
	ID          int       `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Amount      int       `json:"amount" db:"amount"`
	Email       string    `json:"email" db:"email"`
	Mobile      string    `json:"mobile" db:"mobile"`
	Description string    `json:"description" db:"description"`
	RefID       string    `json:"ref_id" db:"ref_id"` //maybe should be bigint
	Message     string    `json:"message" db:"message"`
	Athority    string    `json:"athority" db:"athority"`
	Status      int       `json:"status" db:"status"`
	UserID      int       `json:"user_id" db:"user_id"`
}

// String is not required by pop and may be deleted
func (p Payment) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Payments is not required by pop and may be deleted
type Payments []Payment

// String is not required by pop and may be deleted
func (p Payments) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Payment) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: p.Amount, Name: "Amount"},
		&validators.StringIsPresent{Field: p.Email, Name: "Email"},
		&validators.StringIsPresent{Field: p.Mobile, Name: "Mobile"},
		&validators.StringIsPresent{Field: p.Description, Name: "Description"},
		&validators.StringIsPresent{Field: p.RefID, Name: "RefID"}, //IntIsPresent{
		&validators.StringIsPresent{Field: p.Message, Name: "Message"},
		&validators.StringIsPresent{Field: p.Athority, Name: "Athority"},
		&validators.IntIsPresent{Field: p.Status, Name: "Status"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Payment) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Payment) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
