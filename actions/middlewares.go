package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"github.com/sadegh-shahbazi/vzk/models"
	"github.com/tomasen/realip"
	"time"
)

func WriterCanMiddleware(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {

		id := c.Session().Get("current_user_id")
		u := &models.User{}
		tx := c.Value("tx").(*pop.Connection)

		if id != nil {
			err := tx.Find(u, id)
			if err != nil {
				return errors.WithStack(err)
			}
		}
		if u.RoleID <= 2 { //just 3, 4, 5 (writer, editor, admin) can
			c.Flash().Add("danger", "You must be Writer or Editor or Admin to see that page")
			return c.Redirect(302, "/")
		}

		return next(c)
	}
}
func EditorCanMiddleware(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {

		id := c.Session().Get("current_user_id")
		u := &models.User{}
		tx := c.Value("tx").(*pop.Connection)

		if id != nil {
			err := tx.Find(u, id)
			if err != nil {
				return errors.WithStack(err)
			}
		}
		if u.RoleID <= 3 { //just 5 (admin) and 4 (editor) can
			c.Flash().Add("danger", "You must be Editor to see that page")
			return c.Redirect(302, "/")
		}

		return next(c)
	}
}

func AdminCanMiddleware(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {

		id := c.Session().Get("current_user_id")
		u := &models.User{}
		tx := c.Value("tx").(*pop.Connection)

		if id != nil {
			err := tx.Find(u, id)
			if err != nil {
				return errors.WithStack(err)
			}
		}
		if u.RoleID <= 4 { //just 5 (admin) can
			c.Flash().Add("danger", "You must be Admin to see that page")
			return c.Redirect(302, "/")
		}

		return next(c)
	}
}

func ThrottleMiddleware(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {

		clientIp := realip.FromRequest(c.Request())
		tx := c.Value("tx").(*pop.Connection)

		//start delete old ips
		var allIps models.Throttles
		err := tx.All(&allIps)

		for _, all := range allIps {
			diff := time.Now().Sub(all.CreatedAt)
			//20 minutes
			if diff > 20*time.Second*60 {
				err = tx.Destroy(&all)
			}
		}
		//end delete old ips

		err = tx.Create(&models.Throttle{
			Ip: clientIp,
		})

		var sameIps models.Throttles
		err = tx.Where("ip = ?", clientIp).All(&sameIps)
		countSameIps := len(sameIps)

		if countSameIps > 40 {
			c.Flash().Add("danger", "تعداد درخواست های شما بیشتر از حد مجاز است. لطفا ۲۰ دقیقه دیگر امتحان بفرمایید.")
			return c.Redirect(403, "/login")
		}

		if err != nil {
			return error(err)
		}

		return next(c)
	}
}
