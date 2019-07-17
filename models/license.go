package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
	"time"
)

type License struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	LicenseText string    `json:"license_text" db:"license_text"`
	Email       string    `json:"email" db:"email"`
	IsSold      bool      `json:"is_sold" db:"is_sold"`
	PeriodDay   int       `json:"period_day" db:"period_day"`
	AntiID      int       `json:"anti_id" db:"anti_id"`
}

// String is not required by pop and may be deleted
func (l License) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Licenses is not required by pop and may be deleted
type Licenses []License

// String is not required by pop and may be deleted
func (l Licenses) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (l *License) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: l.LicenseText, Name: "LicenseText"},
		&validators.StringIsPresent{Field: l.Email, Name: "Email"},
		&validators.IntIsPresent{Field: l.PeriodDay, Name: "PeriodDay"},
		&validators.IntIsPresent{Field: l.AntiID, Name: "AntiID"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (l *License) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (l *License) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
