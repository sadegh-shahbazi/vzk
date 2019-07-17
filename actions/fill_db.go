package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"github.com/sadegh-shahbazi/vzk/models"
	"time"
)

// FillDBFillDB default implementation.
func FillDB(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	//fill database with important data
	count, _ := tx.Count(models.Role{})
	if count == 0 {
		err := tx.Create(&models.Role{Name: "user"})
		err = tx.Create(&models.Role{Name: "vip"})
		err = tx.Create(&models.Role{Name: "writer"})
		err = tx.Create(&models.Role{Name: "editor"})
		err = tx.Create(&models.Role{Name: "admin"})
		if err != nil {
			return error(err)
		}
	}

	count, _ = tx.Count(models.AntivirushaType{})
	if count == 0 {
		err := tx.Create(&models.AntivirushaType{Name: "none", NameFa: "هیچ کدام"})
		err = tx.Create(&models.AntivirushaType{Name: "pc", NameFa: "کامپیوتر"})
		err = tx.Create(&models.AntivirushaType{Name: "mofid", NameFa: "مفید"})
		err = tx.Create(&models.AntivirushaType{Name: "mobile", NameFa: "موبایل"})
		err = tx.Create(&models.AntivirushaType{Name: "windows", NameFa: "ویندوز"})
		err = tx.Create(&models.AntivirushaType{Name: "software", NameFa: "سایر نرم افزار ها"})
		if err != nil {
			return error(err)
		}
	}

	count, _ = tx.Count(models.User{})
	if count == 0 {
		err := tx.Create(&models.User{
			Email:        "sadeghsadegh20@yahoo.com",
			PasswordHash: "$2a$10$BM9KP9uBsIGw6sbrpdEIfuBFX9bFYn4rPgOK9O7kLSoNS.Cskz4TC",
			RoleID:       5,
			Disliked:     2, Liked: 1, IsActive: true, Name: "sadegh", Image: "/uploads/AnonymousMale.jpg", RememberToken: "ddd", Bio: "about me", VipEndTime: time.Now(),
		})
		err = tx.Create(&models.User{
			Email:        "sara@sara.com",
			PasswordHash: "$2a$10$BM9KP9uBsIGw6sbrpdEIfuBFX9bFYn4rPgOK9O7kLSoNS.Cskz4TC",
			RoleID:       1,
			Disliked:     0, Liked: 0, IsActive: true, Name: "sara", Image: "/uploads/AnonymousMale.jpg", RememberToken: "ddd", Bio: "about sara", VipEndTime: time.Now(),
		})
		err = tx.Create(&models.User{
			Email:        "sadegh@sadegh.com",
			PasswordHash: "$2a$10$BM9KP9uBsIGw6sbrpdEIfuBFX9bFYn4rPgOK9O7kLSoNS.Cskz4TC",
			RoleID:       1,
			Disliked:     0, Liked: 10, IsActive: true, Name: "ss", Image: "/uploads/AnonymousMale.jpg", RememberToken: "ddd", Bio: "about ss", VipEndTime: time.Now(),
		})
		if err != nil {
			return error(err)
		}
	}

	count, _ = tx.Count(models.Antivirusha{})
	if count == 0 {
		err := tx.Create(&models.Antivirusha{
			Name:           "software",
			NameFa:         "software",
			IsActive:       true,
			IsFree:         true,
			AntivirusOrder: 100,
			Type:           "software",
			Image:          "/uploads/images/bitdefender.jpg",
		})
		err = tx.Create(&models.Antivirusha{
			Name:           "bitdefender",
			NameFa:         "بیت دفندر",
			IsActive:       true,
			IsFree:         true,
			AntivirusOrder: 100,
			Type:           "pc",
			Image:          "/uploads/images/bitdefender.jpg",
		})
		err = tx.Create(&models.Antivirusha{
			Name:           "kaspersky",
			NameFa:         "کسپراسکی",
			IsActive:       true,
			IsFree:         true,
			AntivirusOrder: 100,
			Type:           "pc",
			Image:          "/uploads/images/kaspersky.jpg",
		})
		err = tx.Create(&models.Antivirusha{
			Name:           "eset",
			NameFa:         "ایست",
			IsActive:       true,
			IsFree:         true,
			AntivirusOrder: 100,
			Type:           "pc",
			Image:          "/uploads/images/eset.jpg",
		})
		err = tx.Create(&models.Antivirusha{
			Name:           "eset",
			NameFa:         "ایست",
			IsActive:       true,
			IsFree:         false,
			AntivirusOrder: 100,
			Type:           "mobile",
			Image:          "/uploads/images/eset.jpg",
		})
		err = tx.Create(&models.Antivirusha{
			Name:           "eset",
			NameFa:         "ایست",
			IsActive:       true,
			IsFree:         false,
			AntivirusOrder: 100,
			Type:           "mobile",
			Image:          "/uploads/images/eset.jpg",
		})
		if err != nil {
			return error(err)
		}
	}

	count, _ = tx.Count(models.Setting{})
	if count == 0 {
		err := tx.Create(&models.Setting{
			Name:  "rules_text",
			Value: "این قوانننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننننین هست.",
		})
		err = tx.Create(&models.Setting{
			Name:  "contact_us_text",
			Value: "این تماس با ما هست.",
		})
		err = tx.Create(&models.Setting{
			Name:  "license_can_comment_text",
			Value: "در صورتی که شما لایسنس جدیدی دارید، آن را با دیگران به اشتراک بگذارید.<br>در زیر میتوانید لایسنس هایی که دیگران با شما به اشتراک گذاشته اند را ببنید.<br>در صورتی که لایسنس ها فعال هستند، در قسمت پاسخ به دیگران اطلاع دهید.<br>",
		})
		err = tx.Create(&models.Setting{
			Name:  "index_home_text",
			Value: "<h1>آنتی لایسنس خوش آمدید.</h1><p>صفحه اصلی اینجاست</p>",
		})
		err = tx.Create(&models.Setting{
			Name:  "index_antivirus_text",
			Value: "<h1>آنتی لایسنس خوش آمدید.</h1><p>صفحه آنتی ویروس ها اینجاست</p>",
		})
		err = tx.Create(&models.Setting{
			Name:  "index_windows_text",
			Value: "<h1>آنتی لایسنس خوش آمدید.</h1><p>صفحه ویندوز اینجاست</p>",
		})
		err = tx.Create(&models.Setting{
			Name:  "merchant_id",
			Value: "1ff7fab4-89b5-11e9-a460-000c29344814",
		})
		err = tx.Create(&models.Setting{
			Name:  "vip_monthly_price",
			Value: "1000",
		})
		err = tx.Create(&models.Setting{
			Name:  "vip_yearly_price",
			Value: "5000",
		})
		err = tx.Create(&models.Setting{
			Name:  "index_home_mofid_text",
			Value: "نرم افزار های مفید:",
		})
		err = tx.Create(&models.Setting{
			Name:  "buy_vip_text",
			Value: "خرید اکانت ویژه:",
		})
		if err != nil {
			return error(err)
		}
	}

	count, _ = tx.Count(models.PostType{})
	if count == 0 {
		err := tx.Create(&models.PostType{ID: 1, Name: "پست"})
		err = tx.Create(&models.PostType{ID: 2, Name: "سوال"})
		err = tx.Create(&models.PostType{ID: 3, Name: "مقاله و لایسنس سایر نرم افزار ها"})
		err = tx.Create(&models.PostType{ID: 4, Name: "دانلود"})
		err = tx.Create(&models.PostType{ID: 5, Name: "درباره"})
		err = tx.Create(&models.PostType{ID: 6, Name: "کامنت برای لایسنس"})
		if err != nil {
			return error(err)
		}
	}

	count, _ = tx.Count(models.Post{})
	if count == 0 {
		err := tx.Create(&models.Post{Title: "پست اول", Content: "متن پست اول", PostType: models.PostType{ID: 1}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "پست دوم", Content: "متن پست دوم", PostType: models.PostType{ID: 1}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "پست سوم", Content: "متن پست سوم", PostType: models.PostType{ID: 1}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "پست چهارم", Content: "متن پست چهارم", PostType: models.PostType{ID: 1}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "پست پنجم", Content: "متن پست پنجم", PostType: models.PostType{ID: 1}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "سوال اول", Content: "متن سوال اول", PostType: models.PostType{ID: 2}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "سوال دوم", Content: "متن سوال دوم", PostType: models.PostType{ID: 2}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "سوال سوم", Content: "متن سوال سوم", PostType: models.PostType{ID: 2}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "سوال چهارم", Content: "متن سوال چهارم", PostType: models.PostType{ID: 2}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "سوال پنجم", Content: "متن سوال پنجم", PostType: models.PostType{ID: 2}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "سوال ششم", Content: "متن سوال ششم", PostType: models.PostType{ID: 2}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "سوال هفتم", Content: "متن سوال هفتم", PostType: models.PostType{ID: 2}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "سوال هشتم", Content: "متن سوال هشتم", PostType: models.PostType{ID: 2}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "سوال نهم", Content: "متن سوال نهم", PostType: models.PostType{ID: 2}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "سوال دهم", Content: "متن سوال دهم", PostType: models.PostType{ID: 2}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "مقاله اول", Content: "متن مقاله اول", PostType: models.PostType{ID: 3}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "مقاله دوم", Content: "متن مقاله دوم", PostType: models.PostType{ID: 3}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "مقاله سوم", Content: "متن مقاله سوم", PostType: models.PostType{ID: 3}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "مقاله چهارم", Content: "متن مقاله چهارم", PostType: models.PostType{ID: 3}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "مقاله پنجم", Content: "متن مقاله پنجم", PostType: models.PostType{ID: 3}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "مقاله ششم", Content: "متن مقاله ششم", PostType: models.PostType{ID: 3}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "مقاله هفتم", Content: "متن مقاله هفتم", PostType: models.PostType{ID: 3}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "مقاله هشتم", Content: "متن مقاله هشتم", PostType: models.PostType{ID: 3}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "مقاله نهم", Content: "متن مقاله نهم", PostType: models.PostType{ID: 3}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "مقاله دهم", Content: "متن مقاله دهم", PostType: models.PostType{ID: 3}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		err = tx.Create(&models.Post{Title: "مقاله یازدهم", Content: "متن مقاله یازدهم", PostType: models.PostType{ID: 3}, User: models.User{ID: 1}, IsActive: true, Liked: 0, Disliked: 0, AntivirushaID: 1, CanComment: true})
		if err != nil {
			return error(err)
		}
	}

	count, _ = tx.Count(models.Comment{})
	if count == 0 {
		err := tx.Create(&models.Comment{Content: "اولین جواب به پست اول", Post: models.Post{ID: 1}, User: models.User{ID: 1}, IsReplayed: false})
		err = tx.Create(&models.Comment{Content: "دومین جواب به پست اول", Post: models.Post{ID: 1}, User: models.User{ID: 1}, IsReplayed: false})
		err = tx.Create(&models.Comment{Content: "اولین جواب به سوال اول", Post: models.Post{ID: 6}, User: models.User{ID: 1}, IsReplayed: false})
		err = tx.Create(&models.Comment{Content: "دومین جواب به سوال اول", Post: models.Post{ID: 6}, User: models.User{ID: 1}, IsReplayed: false})
		err = tx.Create(&models.Comment{Content: "اولین جواب به مقاله اول", Post: models.Post{ID: 16}, User: models.User{ID: 1}, IsReplayed: false})
		err = tx.Create(&models.Comment{Content: "دومین جواب به مقاله اول", Post: models.Post{ID: 16}, User: models.User{ID: 1}, IsReplayed: false})
		if err != nil {
			return error(err)
		}
	}

	count, _ = tx.Count(models.License{})
	if count == 0 {
		err := tx.Create(&models.License{AntivirushaID: 1, ContentOne: "this is content one of antivirus 1", ContentOneVip: "this is content one vip 11111111111111", ContentTwo: "this is content two"})
		err = tx.Create(&models.License{AntivirushaID: 2, ContentOne: "this is content one of antivirus 2", ContentOneVip: "this is content one vip 222222222222222", ContentTwo: "this is content two"})
		if err != nil {
			return error(err)
		}
	}

	//fill database with test data

	return c.Render(200, r.String("database seeded."))
}
