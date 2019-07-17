package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/sadegh-shahbazi/vzk/models"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// AdminAdmin default implementation.
func AdminHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("admin/admin.html"))
}

func UploadHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("admin/upload.html"))
}

func UploadPostHandler(c buffalo.Context) error {
	uid := c.Session().Get("current_user_id")
	if uid == nil {
		return c.Redirect(302, "/")
	}
	tx := c.Value("tx").(*pop.Connection)
	var user models.User
	err := tx.Find(&user, uid)
	if err != nil {
		return error(err)
	}
	if user.RoleID != 4 && user.RoleID != 5 {
		return c.Redirect(302, "/")
	}

	image1, err := c.File("image1")
	if err != nil {
		return error(err)
	}

	if image1.Size < (3 * 1024 * 1024) { // convert to MB

	}

	purePath := "./"
	uploadPathAdd := "uploads/admin/randomName"

	uploadPath := purePath + uploadPathAdd

	fileBytes, err := ioutil.ReadAll(image1.File)
	if err != nil {
		log.Fatal("INVALID_FILE")
	}

	withTime := strconv.FormatInt(int64(time.Now().UnixNano()), 10)
	fileFullNameToSave := withTime + image1.Filename

	newPath := filepath.Join(uploadPath, fileFullNameToSave)

	//start upload image
	newFile, err := os.Create(newPath)
	if err != nil {
		return error(err)
	}
	_, err = newFile.Write(fileBytes)
	if err != nil {
		return error(err)
	}

	err = newFile.Close()
	if err != nil {
		return error(err)
	}
	//start send to front end
	image1Temp := "https://" + c.Request().Host + "/" + uploadPathAdd + "/" + fileFullNameToSave
	//image2Temp := "/" + uploadPathAdd + "/" + fileFullNameToSave
	image2Temp := "."
	//c.Set("image1", image1Temp)
	//end send to front end
	c.Flash().Add("success", "فایل "+
		image1.Filename+
		"آپلود گردید.")
	c.Flash().Add("success", image1Temp)
	c.Flash().Add("success", image2Temp)

	redirectBack := c.Request().Header.Get("Referer")
	return c.Redirect(302, redirectBack)
}

//آپلود با اسم مشخص
func UploadSpecificHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("admin/upload_specific.html"))
}

func UploadSpecificPostHandler(c buffalo.Context) error {
	uid := c.Session().Get("current_user_id")
	if uid == nil {
		return c.Redirect(302, "/")
	}
	tx := c.Value("tx").(*pop.Connection)
	var user models.User
	err := tx.Find(&user, uid)
	if err != nil {
		return error(err)
	}
	if user.RoleID != 4 && user.RoleID != 5 {
		return c.Redirect(302, "/")
	}

	image1, err := c.File("image1")
	if err != nil {
		return error(err)
	}

	if image1.Size < (3 * 1024 * 1024) { // convert to MB

	}

	purePath := "./"
	uploadPathAdd := "uploads/admin/specific"

	uploadPath := purePath + uploadPathAdd

	fileBytes, err := ioutil.ReadAll(image1.File)
	if err != nil {
		log.Fatal("INVALID_FILE")
	}

	fileFullNameToSave := image1.Filename

	newPath := filepath.Join(uploadPath, fileFullNameToSave)

	//start upload image
	newFile, err := os.Create(newPath)
	if err != nil {
		return error(err)
	}
	_, err = newFile.Write(fileBytes)
	if err != nil {
		return error(err)
	}

	err = newFile.Close()
	if err != nil {
		return error(err)
	}
	//start send to front end
	image1Temp := "https://" + c.Request().Host + "/" + uploadPathAdd + "/" + fileFullNameToSave
	//image2Temp := "/" + uploadPathAdd + "/" + fileFullNameToSave
	image2Temp := "."
	//end send to front end
	c.Flash().Add("success", "فایل "+
		image1.Filename+
		"آپلود گردید.")
	c.Flash().Add("success", image1Temp)
	c.Flash().Add("success", image2Temp)

	redirectBack := c.Request().Header.Get("Referer")
	return c.Redirect(302, redirectBack)
}
func UploadApkHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("admin/upload_apk.html"))
}

func UploadApkPostHandler(c buffalo.Context) error {
	uid := c.Session().Get("current_user_id")
	if uid == nil {
		return c.Redirect(302, "/")
	}
	tx := c.Value("tx").(*pop.Connection)
	var user models.User
	err := tx.Find(&user, uid)
	if err != nil {
		return error(err)
	}
	if user.RoleID != 4 && user.RoleID != 5 {
		return c.Redirect(302, "/")
	}

	image1, err := c.File("image1")
	if err != nil {
		return error(err)
	}

	if image1.Size < (3 * 1024 * 1024) { // convert to MB

	}

	purePath := "./"
	uploadPathAdd := "uploads/admin/apk"

	uploadPath := purePath + uploadPathAdd

	fileBytes, err := ioutil.ReadAll(image1.File)
	if err != nil {
		log.Fatal("INVALID_FILE")
	}

	fileFullNameToSave := image1.Filename

	newPath := filepath.Join(uploadPath, fileFullNameToSave)

	//start upload image
	newFile, err := os.Create(newPath)
	if err != nil {
		return error(err)
	}
	_, err = newFile.Write(fileBytes)
	if err != nil {
		return error(err)
	}

	err = newFile.Close()
	if err != nil {
		return error(err)
	}
	//start send to front end
	image1Temp := "https://" + c.Request().Host + "/" + uploadPathAdd + "/" + fileFullNameToSave
	//image2Temp := "/" + uploadPathAdd + "/" + fileFullNameToSave
	image2Temp := "."
	//end send to front end
	c.Flash().Add("success", "فایل "+
		image1.Filename+
		"آپلود گردید.")
	c.Flash().Add("success", image1Temp)
	c.Flash().Add("success", image2Temp)

	redirectBack := c.Request().Header.Get("Referer")
	return c.Redirect(302, redirectBack)
}

func AdminPostAddHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	var antivirushas []models.Antivirusha
	err := tx.All(&antivirushas)
	if err != nil {
		return error(err)
	}
	c.Set("antivirushas", antivirushas)

	var postTypes []models.PostType
	err = tx.All(&postTypes)
	if err != nil {
		return error(err)
	}
	c.Set("postTypes", postTypes)

	return c.Render(200, r.HTML("admin/post_add.html"))
}
func AdminPostAddPostHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	antivirushaID, err := strconv.Atoi(c.Param("antivirusha"))
	postTypeID, err := strconv.Atoi(c.Param("postType"))
	content := c.Param("content")
	title := c.Param("title")
	err = tx.Create(&models.Post{
		PostTypeID:    postTypeID,
		Content:       content,
		AntivirushaID: antivirushaID,
		Title:         title,
		CanComment:    true,
		IsActive:      true,
	})
	if err != nil {
		return error(err)
	}

	c.Flash().Add("success", "پست با موفقیت ساخته شد.")

	return c.Redirect(302, c.Request().Header.Get("Referer"))
}

func AdminLicenseAddHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	licenseID := c.Param("license_id")
	var license models.License
	err := tx.Find(&license, licenseID)
	if err != nil {
		return error(err)
	}
	c.Set("license", license)

	var anti models.Antivirusha
	err = tx.Find(&anti, license.AntivirushaID)
	if err != nil {
		return error(err)
	}

	title := "جدیدترین لایسنس رایگان" + " " + anti.Name + " " + anti.NameFa + " " + time.Now().Format("2006-01-02 3:04PM")

	c.Set("title", title)

	///////
	Referer := c.Request().Header.Get("Referer")
	c.Set("Referer", Referer)

	return c.Render(200, r.HTML("admin/edit_license.html"))
}
func AdminLicenseAddPostHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	licenseID := c.Param("license_id")
	var license models.License
	err := tx.Find(&license, licenseID)
	if err != nil {
		return error(err)
	}

	licenseOld := license

	license.ContentOne = c.Param("ContentOne")
	license.ContentOneVip = c.Param("ContentOneVip")
	license.ContentTwo = c.Param("ContentTwo")
	license.ContentTwoVip = c.Param("ContentTwoVip")
	license.ContentThree = c.Param("ContentThree")
	license.ContentThreeVip = c.Param("ContentThreeVip")
	license.ContentFour = c.Param("ContentFour")
	license.ContentFourVip = c.Param("ContentFourVip")
	license.ContentFive = c.Param("ContentFive")

	addPost := c.Param("add_post")

	if addPost == "add_post" {
		var anti models.Antivirusha
		err = tx.Find(&anti, license.AntivirushaID)
		err = tx.Create(&models.Post{
			Liked:      0,
			Disliked:   0,
			PostTypeID: 1,
			IsActive:   true,
			UserID:     c.Session().Get("current_user_id").(int),
			CanComment: false,
			Title:      c.Param("title"),
			Content: licenseOld.ContentOne + licenseOld.ContentOneVip +
				licenseOld.ContentTwo + licenseOld.ContentTwoVip +
				licenseOld.ContentThree + licenseOld.ContentThreeVip +
				licenseOld.ContentFour + licenseOld.ContentFourVip +
				licenseOld.ContentFive,
			AntivirushaID: license.AntivirushaID,
		})

		////////////////start create site map

		go SiteMapHandler(c)

		////////////////

	}

	err = tx.Update(&license)
	if err != nil {
		return error(err)
	}

	c.Flash().Add("success", "لایسنس با موفقیت آپدیت شد.")

	return c.Redirect(302, c.Param("Referer"))
	//return c.Redirect(302, c.Request().Header.Get("Referer"))
}

func AdminMataleb(c buffalo.Context) error {
	return c.Render(200, r.HTML("admin/mataleb.html"))
}

func AddVipAccount(c buffalo.Context) error {
	return c.Render(200, r.HTML("admin/add_vip_account.html"))
}

func AddVipAccountPost(c buffalo.Context) error {
	email := c.Param("email")
	month, err := strconv.Atoi(c.Param("month"))
	tx := c.Value("tx").(*pop.Connection)
	var user models.User
	err = tx.Where("email = ?", email).First(&user)
	if err != nil {
		c.Flash().Add("error", "ایمیل در دیتابیس پیدا نشد.")
		c.Flash().Add("error", "ایمیل های مشابه عبارت اند از:")
		var tempUsers models.Users
		err = tx.Where("email LIKE ?", "%"+email+"%").All(&tempUsers)
		if err != nil {
			c.Flash().Add("error", "ایمیل مشابه نیز پیدا نشد.")
			return c.Redirect(302, c.Request().Header.Get("Referer"))
		}
		for _, tempUser := range tempUsers {
			c.Flash().Add("error", tempUser.Email)
		}

		return c.Redirect(302, c.Request().Header.Get("Referer"))
	}

	days := month * 31
	user.VipEndTime = time.Now().AddDate(0, 0, days)

	err = tx.Update(&user)
	if err != nil {
		return error(err)
	}
	c.Flash().Add("success", "تعداد "+strconv.Itoa(month)+" ماه به اکانت ویژه یوزر شماره "+strconv.Itoa(user.ID)+" با ایمیل "+user.Email+" با موفقیت اضافه گردید. ")

	return c.Redirect(302, c.Request().Header.Get("Referer"))
}

func ChangeLicenseUpdateDateHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	var antivirushas models.Antivirushas
	err := tx.Eager("License").All(&antivirushas)
	if err != nil {
		return error(err)
	}
	c.Set("antivirushas", antivirushas)

	return c.Render(200, r.HTML("admin/change_license_updated_at.html"))
}

func ChangeLicenseUpdateDatePostHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	licenseID := c.Param("license_id")
	days, err := strconv.Atoi(c.Param("days"))
	var license models.License
	err = tx.Find(&license, licenseID)
	if err != nil {
		return error(err)
	}

	sub31Days := time.Now().AddDate(0, 0, -days)
	c.Set("s", sub31Days)

	err = tx.RawQuery("UPDATE licenses SET updated_at = ? WHERE licenses.id = ?;", sub31Days, license.ID).Exec()

	c.Flash().Add("success", "لایسنس یک ماه به عقب برگشت")

	return c.Redirect(302, c.Request().Header.Get("Referer"))
}

func AdminReportHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	var users []models.User
	err := tx.All(&users)
	c.Set("totalUsers", len(users))
	users = []models.User{}
	err = tx.Where("created_at >= ?", time.Now().AddDate(0, 0, -1)).All(&users)
	c.Set("todayTotalUsers", len(users))
	users = []models.User{}
	err = tx.Where("role_id = '1'").All(&users)
	c.Set("normalUsers", len(users))
	users = []models.User{}
	err = tx.Where("role_id = '2'").All(&users)
	c.Set("vipOldUsers", len(users))
	users = []models.User{}
	err = tx.Where("vip_end_time >= ?", time.Now()).All(&users)
	c.Set("vipUsers", len(users))
	users = []models.User{}
	err = tx.Where("vip_end_time >= ?", time.Now()).Where("created_at >= ?", time.Now().AddDate(0, 0, -1)).All(&users)
	c.Set("todayVipUsers", len(users))
	users = []models.User{}
	err = tx.Where("role_id = '3'").All(&users)
	c.Set("writerUsers", len(users))
	users = []models.User{}
	err = tx.Where("role_id = '4'").All(&users)
	c.Set("editorUsers", len(users))
	users = []models.User{}
	err = tx.Where("role_id = '5'").All(&users)
	c.Set("adminUsers", len(users))

	var posts []models.Post
	err = tx.All(&posts)
	c.Set("allPosts", len(posts))
	posts = []models.Post{}
	err = tx.Where("created_at >= ?", time.Now().AddDate(0, 0, -1)).All(&posts)
	c.Set("todayPosts", len(posts))
	posts = []models.Post{}
	err = tx.Where("created_at >= ?", time.Now().Add(time.Hour*6)).All(&posts)
	c.Set("past6HourPosts", len(posts))

	var comments []models.Comment
	err = tx.All(&comments)
	c.Set("allComments", len(comments))
	comments = []models.Comment{}
	err = tx.Where("created_at >= ?", time.Now().AddDate(0, 0, -1)).All(&comments)
	c.Set("todayComments", len(comments))
	comments = []models.Comment{}

	var antiviruses []models.Antivirusha
	err = tx.Where("type = 'pc'").All(&antiviruses)
	c.Set("pcAntiviruses", len(antiviruses))
	antiviruses = []models.Antivirusha{}
	err = tx.Where("type = 'mofid'").All(&antiviruses)
	c.Set("mofidAntiviruses", len(antiviruses))
	antiviruses = []models.Antivirusha{}
	err = tx.Where("type = 'mobile'").All(&antiviruses)
	c.Set("mobileAntiviruses", len(antiviruses))
	antiviruses = []models.Antivirusha{}
	err = tx.Where("type = 'windows'").All(&antiviruses)
	c.Set("windowsAntiviruses", len(antiviruses))
	antiviruses = []models.Antivirusha{}
	err = tx.Where("type = 'software'").All(&antiviruses)
	c.Set("softwareAntiviruses", len(antiviruses))
	antiviruses = []models.Antivirusha{}

	totalDeposit := 0
	var payments []models.Payment
	err = tx.Where("message = 'ok'").All(&payments)
	for _, payment := range payments {
		totalDeposit += payment.Amount
	}
	c.Set("totalDeposit", totalDeposit)

	totalDeposit = 0
	payments = []models.Payment{}
	err = tx.Where("message = 'ok'").Where("created_at >= ?", time.Now().AddDate(0, 0, -30)).All(&payments)
	for _, payment := range payments {
		totalDeposit += payment.Amount
	}
	c.Set("totalMonthDeposit", totalDeposit)

	totalDeposit = 0
	payments = []models.Payment{}
	err = tx.Where("message = 'ok'").Where("created_at >= ?", time.Now().AddDate(0, 0, -1)).All(&payments)
	for _, payment := range payments {
		totalDeposit += payment.Amount
	}
	c.Set("totalDayDeposit", totalDeposit)

	if err != nil {
		return error(err)
	}

	return c.Render(200, r.HTML("admin/report.html"))
}

func AdminAllTextsSettingsHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	var settings []models.Setting
	err := tx.All(&settings)
	if err != nil {
		return error(err)
	}
	c.Set("settings", settings)

	return c.Render(200, r.HTML("admin/settingsTextAll.html"))
}
func AdminTextSettingsHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	var setting models.Setting
	settingID := c.Param("setting_id")
	err := tx.Find(&setting, settingID)
	if err != nil {
		return error(err)
	}
	c.Set("setting", setting)
	return c.Render(200, r.HTML("admin/settingsTextEdit"))
}
func AdminTextSettingsPostHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	var setting models.Setting
	settingID := c.Param("setting_id")
	err := tx.Find(&setting, settingID)
	if err != nil {
		return error(err)
	}
	value := c.Param("value")
	setting.Value = value
	err = tx.Update(&setting)
	if err != nil {
		return error(err)
	}
	c.Flash().Add("success", "تغییرات انجام شد.")

	return c.Redirect(302, c.Request().Header.Get("Referer"))
}
