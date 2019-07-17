package actions

import (
	"strconv"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/pop"
	"github.com/sadegh-shahbazi/vzk/models"
	ptime "github.com/yaa110/go-persian-calendar"
)

var r *render.Engine
var assetsBox = packr.New("app:assets", "../public")

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.html",

		// Box containing all of the templates:
		TemplatesBox: packr.New("app:templates", "../templates"),
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{
			// uncomment for non-Bootstrap form helpers:
			// "form":     plush.FormHelper,
			// "form_for": plush.FormForHelper,
			"licenseDateColor": func(lastUpdate time.Time) string {
				t1 := lastUpdate
				t2 := time.Now()
				days := t2.Sub(t1).Hours() / 24
				if days <= 30 {
					return "green"
				}
				if days <= 90 {
					return "orange"
				}
				return "red"
			},

			//used in home page
			"antivirusNewVersionDate": func() int {
				year, month, _ := time.Now().Date()
				if int(month) < 8 {
					return year
				} else {
					return year + 1
				}
			},

			//
			"now": func() string {
				t := time.Now()
				pt := ptime.New(t).Format("dd-MMM-yyyy E")
				return pt
			},

			//used in home page
			"persianLicenseUpdate": func(date time.Time) string {
				if date.IsZero() {
					return "لایسنس آپدیت نشده"
				} else {
					pt := ptime.New(date).Format("E dd-MMM-yyyy")
					return pt
				}
			},

			//used in show date of comments and posts and article and questions
			"persianDateClock": func(date time.Time) string {
				if date.IsZero() {
					return "لایسنس آپدیت نشده"
				} else {
					pt := ptime.New(date).Format("E dd-MMM-yyyy ساعت: HH:MM")
					return pt
				}
			},

			//usd in license.html addSpaceToNullString
			"vipUserCanSee": func(VipTime time.Time, text string) string {
				if VipTime.Sub(time.Now()) > 0 {
					return text
				}
				return string("اکانت ویژه شما پایان یافته است. برای مشاهده محتوا باید از طریق منوی بالای سایت اکانت ویژه تهیه نمایید.")
			},

			//used in comments in all posts
			"showUserRole": func(id int) string {
				switch id {
				case 1:
					return "یوزر معمولی"
				case 2:
					return "یوزر ویژه"
				case 3:
					return "نویسنده"
				case 4:
					return "مدیر"
				default:
					return "ادمین"
				}
			},

			//use in every edit buttons for owner - editor - admin
			"ownerCan": func(currentUser interface{}, ownerID int) bool {
				//(For example the assertion: `c.Value("current_user").(*models.User)`)
				//You can try `currentUser.(bool)` which will panic with an error message saying what type it actually is (since it is not a `bool`)
				//I should understand this code:
				u := currentUser.(*models.User)

				if u.RoleID == 4 || u.RoleID == 5 || u.ID == ownerID {
					return true
				}

				return false
			},

			//use top menu
			"softwareTitle": func(i int, s string) string {
				//tempString:= strings.Replace(s, " ", "%20", -1)
				newString := strconv.Itoa(i)
				newString = newString + "-"
				newString = newString + s

				return newString
			},

			//use in every edit buttons for editor - admin
			"editorCan": func(currentUser interface{}) bool {
				//(For example the assertion: `c.Value("current_user").(*models.User)`)
				//You can try `currentUser.(bool)` which will panic with an error message saying what type it actually is (since it is not a `bool`)
				//I should understand this code:
				u := currentUser.(*models.User)

				if u.RoleID == 4 || u.RoleID == 5 {
					return true
				}

				return false
			},

			//not used yet just test
			"aclTest": func(c buffalo.Context, tx *pop.Connection) int {

				count, _ := tx.Count(models.Setting{})
				return count
			},
		},
	})
}
