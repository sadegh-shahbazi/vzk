package models

import (
	"encoding/json"
	"github.com/gobuffalo/nulls"
	"strings"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

//User is a generated model from buffalo-auth, it serves as the base for username/password authentication.
type User struct {
	ID                   int          `json:"id" db:"id"`
	CreatedAt            time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time    `json:"updated_at" db:"updated_at"`
	Email                string       `json:"email" db:"email"`
	PasswordHash         string       `json:"password_hash" db:"password_hash"`
	Password             string       `json:"-" db:"-"`
	PasswordConfirmation string       `json:"-" db:"-"`
	UID                  uuid.UUID    `json:"uid" db:"uid"`
	RecoveryCode         nulls.String `json:"-" db:"recovery_code"`
	RecoveryExp          nulls.Time   `json:"-" db:"recovery_expiration"`
	Name                 string       `json:"name" db:"name"`
	RememberToken        string       `json:"remember_token" db:"remember_token"`
	IsActive             bool         `json:"is_active" db:"is_active"`
	Image                string       `json:"image" db:"image"`
	ImageOriginal        string       `json:"image_original" db:"image_original"`
	Bio                  string       `json:"bio" db:"bio"`
	Liked                int          `json:"liked" db:"liked"`
	Disliked             int          `json:"dislike" db:"disliked"`
	VipEndTime           time.Time    `json:"vip_end_time" db:"vip_end_time"`
	RoleID               int          `json:"role_id" db:"role_id"`
	Balance              int          `json:"balance" db:"balance"`
	Withdraw             int          `json:"withdraw" db:"withdraw"`
	LastWithdrawDate     time.Time    `json:"last_withdraw_date" db:"last_withdraw_date"`
	//Relations:

}

// Create wraps up the pattern of encrypting the password and
// running validations. Useful when writing tests.
//func (u *User) Create(tx *pop.Connection) (*validate.Errors, error) {
//	u.Email = strings.ToLower(u.Email)
//	ph, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
//	if err != nil {
//		return validate.NewErrors(), errors.WithStack(err)
//	}
//	u.PasswordHash = string(ph)
//	return tx.ValidateAndCreate(u)
//}
//
//// String is not required by pop and may be deleted
//func (u User) String() string {
//	ju, _ := json.Marshal(u)
//	return string(ju)
//}
//
//// Users is not required by pop and may be deleted
//type Users []User
//
//// String is not required by pop and may be deleted
//func (u Users) String() string {
//	ju, _ := json.Marshal(u)
//	return string(ju)
//}
//
//// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
//// This method is not required and may be deleted.
//func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
//	var err error
//	return validate.Validate(
//		&validators.StringIsPresent{Field: u.Email, Name: "Email"},
//		&validators.StringIsPresent{Field: u.PasswordHash, Name: "PasswordHash"},
//		// check to see if the email address is already taken:
//		&validators.FuncValidator{
//			Field:   u.Email,
//			Name:    "Email",
//			Message: "%s is already taken",
//			Fn: func() bool {
//				var b bool
//				q := tx.Where("email = ?", u.Email)
//				if u.UID != uuid.Nil {
//					q = q.Where("id != ?", u.ID)
//				}
//				b, err = q.Exists(u)
//				if err != nil {
//					return false
//				}
//				return !b
//			},
//		},
//	), err
//}
//
//// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
//// This method is not required and may be deleted.
//func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
//	var err error
//	return validate.Validate(
//		&validators.StringIsPresent{Field: u.Password, Name: "Password"},
//		&validators.StringsMatch{Name: "Password", Field: u.Password, Field2: u.PasswordConfirmation, Message: "Password does not match confirmation"},
//	), err
//}
//
//// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
//// This method is not required and may be deleted.
//func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
//	return validate.NewErrors(), nil
//}

const minPasswordLength int = 1

// Create wraps up the pattern of encrypting the password and
// running validations. Useful when writing tests.
func (u *User) Create(tx *pop.Connection) (*validate.Errors, error) {
	verrs := u.validatePassword()
	if verrs.HasAny() {
		return verrs, nil
	}

	ph, err := encryptPassword(u.Password)
	if err != nil {
		return validate.NewErrors(), errors.WithStack(err)
	}
	u.PasswordHash = string(ph)
	u = u.sanitizeFields()
	return tx.ValidateAndCreate(u)
}

// Update handles the extra work possibly needed during user update,
// hashing password and making email lowercase for consistency
func (u *User) Update(tx *pop.Connection) (*validate.Errors, error) {
	if len(u.Password) != 0 {
		verrs := u.validatePassword()
		if verrs.HasAny() {
			return verrs, nil
		}
		ph, err := encryptPassword(u.Password)
		if err != nil {
			return validate.NewErrors(), errors.WithStack(err)
		}
		u.PasswordHash = ph
	}
	u = u.sanitizeFields()
	return tx.ValidateAndUpdate(u)
}

func (u *User) sanitizeFields() *User {
	// force email to lowercase for better matching
	u.Email = strings.ToLower(u.Email)

	// wipe out password field after it's been hashed
	u.Password = ""
	u.PasswordConfirmation = ""
	return u
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Email, Name: "Email"},
		&validators.StringIsPresent{Field: u.PasswordHash, Name: "PasswordHash"},
		// check to see if the email address is already taken:
		&validators.FuncValidator{
			Field:   u.Email,
			Name:    "Email",
			Message: "%s is already taken",
			Fn: func() bool {
				var b bool
				q := tx.Where("email = ?", u.Email)
				if u.UID != uuid.Nil {
					q = q.Where("id != ?", u.ID)
				}
				b, err = q.Exists(u)
				if err != nil {
					return false
				}
				return !b
			},
		},
	), err
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (u *User) validatePassword() *validate.Errors {
	passwordLengthValidator := validators.StringLengthInRange{
		Field:   u.Password,
		Name:    "Password",
		Message: "Password is not long enough. Minimum of 8 characters required.",
		Min:     minPasswordLength,
		Max:     0,
	}
	passwordConfirmValidator := validators.StringsMatch{
		Field:   u.Password,
		Name:    "PasswordConfirmation",
		Message: "Password does not match confirmation",
		Field2:  u.PasswordConfirmation,
	}
	verrs := validate.NewErrors()
	passwordLengthValidator.IsValid(verrs)
	passwordConfirmValidator.IsValid(verrs)
	return verrs
}

func encryptPassword(password string) (string, error) {
	ph, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(ph), err
}
