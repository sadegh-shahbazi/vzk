package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"github.com/sadegh-shahbazi/vzk/mailers"
	"github.com/sadegh-shahbazi/vzk/models"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

func UsersNew(c buffalo.Context) error {
	u := models.User{}
	c.Set("user", u)
	return c.Render(200, r.HTML("users/new.html", "index/application.html"))
}

// UsersCreate registers a new user with the application.
func UsersCreate(c buffalo.Context) error {
	//u := &models.User{}
	//if err := c.Bind(u); err != nil {
	//	c.Flash().Add("danger", "1111")
	//	return errors.WithStack(err)
	//}
	//
	//tx := c.Value("tx").(*pop.Connection)
	//verrs, err := u.Create(tx)
	//if err != nil {
	//	c.Flash().Add("danger", "222")
	//	return errors.WithStack(err)
	//}
	//
	//if verrs.HasAny() {
	//	c.Set("user", u)
	//	c.Set("errors", verrs)
	//	c.Flash().Add("danger", "این ایمیل در دیتابیس موجود است. لطفا ایمیل دیگری انتخاب نمایید و یا از طریق لینک ورود، وارد سایت شوید.")
	//	return c.Render(200, r.HTML("users/new.html", "index/application.html"))
	//}
	////////////////////////////start edit

	name := c.Param("Name")
	email := c.Param("Email")
	password := c.Param("Password")
	passwordConfirmation := c.Param("PasswordConfirmation")

	if passwordConfirmation != password {
		c.Flash().Add("danger", "رمز عبور و تکرار رمز عبور یکسان نیست.")
		return c.Redirect(302, "/")
	}

	tx := c.Value("tx").(*pop.Connection)
	var u models.User

	err := tx.Where("email = ?", email).First(&u)
	if err != nil {

		u.Name = name
		u.Email = email
		ph, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		u.PasswordHash = string(ph)
		u.Image = "/uploads/AnonymousMale.jpg"
		u.ImageOriginal = "/uploads/AnonymousMale.jpg"
		u.RoleID = 1
		u.IsActive = true
		u.RememberToken = RandStringRunes(141)
		u.VipEndTime = time.Now()
		u.Disliked = 0
		u.Liked = 0
		u.Bio = "داستان زندگی شما"
		u.Withdraw = 0
		u.Balance = 0
		u.LastWithdrawDate = u.VipEndTime

		err = tx.Create(&u)
		if err != nil {
			c.Flash().Add("danger", "111")
			return c.Redirect(302, "/register")
		}

		c.Session().Set("current_user_id", u.ID)
		c.Flash().Add("success", "ثبت نام شما با موفیت انجام شد.")

	} else {
		c.Flash().Add("danger", "این ایمیل در دیتابیس موجود است. لطفا ایمیل دیگری انتخاب نمایید و یا از طریق لینک ورود، وارد سایت شوید.")
		return c.Redirect(302, "/register")
	}

	return c.Redirect(302, "/")
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// SetCurrentUser attempts to find a user based on the current_user_id
// in the session. If one is found it is set on the context.
func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			err := tx.Find(u, uid)
			if err != nil {
				return errors.WithStack(err)
			}
			c.Set("current_user", u)
		}
		return next(c)
	}
}

// Authorize require a user be logged in before accessing a route
func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid == nil {
			c.Session().Set("redirectURL", c.Request().URL.String())

			err := c.Session().Save()
			if err != nil {
				return errors.WithStack(err)
			}

			c.Flash().Add("danger", "شما مجوز دسترسی ندارید.")
			return c.Redirect(302, "/")
		}
		return next(c)
	}
}

// UserRecovery returns the html form for requesting a
// recovery code
func UserRecovery(c buffalo.Context) error { //ok
	return c.Render(200, r.HTML("users/recovery.html", "index/application.html"))
}

// UserRecover returns the html form for using the recovery code
func UserRecover(c buffalo.Context) error {
	recoveryEmail := c.Session().Get("recoveryEmail")
	c.Set("recoveryEmail", recoveryEmail)

	return c.Render(200, r.HTML("users/recover.html", "index/application.html"))
}

// UserRequestRecovery handles requesting recovery for an account
// when user provides matching email
func UserRequestRecovery(c buffalo.Context) error { //ok
	tx := c.Value("tx").(*pop.Connection)
	email := c.Param("Email")
	var user models.User
	err := tx.Where("email = ?", email).First(&user)
	if err != nil {
		c.Flash().Add("danger", "این ایمیل تا کنون در وبسایت ثبت نام نکرده است.")
		return c.Redirect(302, c.Request().Header.Get("Referer"))
	}
	randomNumber := rand.Intn(1000000)
	user.RecoveryCode = nulls.NewString(strconv.Itoa(randomNumber))
	user.RecoveryExp = nulls.NewTime(time.Now().Add(time.Second * 1200)) //20 minutes for expire time of recover password
	err = tx.Update(&user)
	if err != nil {
		return error(err)
	}

	err = mailers.SendRecoveryEmail(user.Email, user.RecoveryCode.String)
	if err != nil {
		return error(err)
	}

	c.Set("Email", user.Email)
	c.Flash().Add("success", "یک ایمیل حاوی کد فعال سازی برای "+user.Email+" ارسال گردید. لطفا پوشه اصلی و پوشه اسپم خود را چک نمایید.")
	c.Flash().Add("success", "کد فعال سازی تا ۲۰ دقیقه دیگر اعتبار دارد.")
	c.Session().Set("recoveryEmail", user.Email)

	return c.Redirect(302, "recoverPath()")
	//return c.Render(200, r.HTML("users/recover.html", "index/application.html"))
}

// UserRequestRecover is used to allow users to reset their passwords
// using the recovery code we've sent them
func UserRequestRecover(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	var user models.User
	email := c.Param("Email")
	code := c.Param("ActivationCode")
	password := c.Param("Password")
	PasswordConfirmation := c.Param("PasswordConfirmation")

	if password != PasswordConfirmation {
		c.Flash().Add("danger", "رمز عبور و تکرار رمز عبور با هم یکسان نیستند.")
		return c.Redirect(302, c.Request().Header.Get("Referer"))
	}

	err := tx.Where("email = ?", email).First(&user)
	if err != nil {
		c.Flash().Add("danger", "ایمیل شما در دیتابیس پیدا نشد.")
		return c.Redirect(302, c.Request().Header.Get("Referer"))
	}

	if nulls.NewString(code) != nulls.String(user.RecoveryCode) {
		c.Flash().Add("danger", "کد فعال سازی شما صحیح نمیباشد.")
		return c.Redirect(302, c.Request().Header.Get("Referer"))
	}

	var timeRecoveryExpTemp time.Time
	timeRecoveryExpTemp = user.RecoveryExp.Time

	timeTemp := timeRecoveryExpTemp.Sub(time.Now())
	if timeTemp < 0 {
		c.Flash().Add("danger", "تاریخ انقضای کد فعال سازی شما گذشته است. برای وارد کردن کد فعال سازی فقط ۲۰ دقیقه مهلت دارید.")
		return c.Redirect(302, c.Request().Header.Get("Referer"))
	}

	if code == user.RecoveryCode.String {

		ph, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return error(err)
		}
		user.PasswordHash = string(ph)
		err = tx.Update(&user)
		if err != nil {
			return error(err)
		}
		c.Flash().Add("success", "پسورد شما برای ایمیل "+user.Email+" با موفقیت تغییر کرد. هم اکنون میتوانید با رمز عبور جدید خود وارد شوید.")

		return c.Redirect(302, "/login")
	}

	return c.Redirect(302, c.Request().Header.Get("Referer"))
}
