package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo-pop/pop/popmw"
	"github.com/gobuffalo/envy"
	csrf "github.com/gobuffalo/mw-csrf"
	forcessl "github.com/gobuffalo/mw-forcessl"
	i18n "github.com/gobuffalo/mw-i18n"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/gobuffalo/packr/v2"
	"github.com/gorilla/sessions"
	"github.com/sadegh-shahbazi/vzk/models"
	"github.com/unrolled/secure"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "production")  //production OR development
var PORT = envy.Get("PORT", "0.0.0.0:3001") //0.0.0.0:3000 OR 127.0.0.1:3000

var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionName:  "_project_session",
			Addr:         PORT,
			SessionStore: sessions.NewCookieStore([]byte("u8h1gd42rae4fu0sw5qfe76fhw5esr5tz8yg54dl6yu9y1t73")),
		})

		// Automatically redirect to SSL
		//app.Use(forceSSL())

		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		// Protect against CSRF attacks. https://www.owasp.org/index.php/Cross-Site_Request_Forgery_(CSRF)
		// Remove to disable this.
		app.Use(csrf.New)

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))

		// Setup and use translations:
		app.Use(translations())
		app.Use(SetCurrentUser)

		index := app.Group("")

		auth := app.Group("")
		auth.Use(Authorize)

		throttle := app.Group("")
		throttle.Use(ThrottleMiddleware)
		throttle.Use(Authorize)

		//start user route

		app.GET("/register", UsersNew).Name("register")
		throttle.POST("/users", UsersCreate)
		app.GET("/login", AuthNew).Name("signin")
		throttle.POST("/signin", AuthCreate)
		app.GET("/recovery", UserRecovery).Name("recoveryPath")
		throttle.POST("/requestRecovery", UserRequestRecovery).Name("recoveryPostPath")
		app.GET("/recover", UserRecover).Name("recoverPath")
		throttle.POST("/requestRecover", UserRequestRecover).Name("recoverPostPath")

		auth.GET("/logout", AuthDestroy).Name("signout") //method should be DELETE /not GET
		//end user route
		throttle.Middleware.Skip(Authorize,
			UsersCreate, AuthCreate,
			UserRequestRecovery, UserRequestRecover,
		)

		editor := app.Group("")
		editor.Use(EditorCanMiddleware)

		app.GET("/upload/randomName", UploadHandler).Name("indexUploadPath")
		app.POST("/upload/randomName", UploadPostHandler).Name("indexUploadPostPath") //it has inner editor middleware
		app.GET("/upload/specific", UploadSpecificHandler).Name("indexUploadSpecificPath")
		app.POST("/upload/specific", UploadSpecificPostHandler).Name("indexUploadSpecificPostPath") //it has inner editor middleware
		app.GET("/upload/apk", UploadApkHandler).Name("indexUploadApkPath")
		app.POST("/upload/apk", UploadApkPostHandler).Name("indexUploadApkPostPath") //it has inner editor middleware

		//start admin route
		admin := editor.Group("/admin")
		admin.GET("/", AdminHandler).Name("adminPath")
		admin.GET("/post/add", AdminPostAddHandler).Name("adminPostAddPath")
		admin.POST("/post/add", AdminPostAddPostHandler).Name("adminPostAddPostPath")
		admin.GET("/license/edit/{license_id}", AdminLicenseAddHandler).Name("adminLicensePath")
		admin.POST("/license/edit/{license_id}", AdminLicenseAddPostHandler).Name("adminLicenseAddPostPath")
		admin.GET("/add_vip_account", AddVipAccount)
		admin.POST("/add_vip_account_post", AddVipAccountPost).Name("addVipAccountPostPath")
		admin.GET("/mataleb", AdminMataleb)
		admin.GET("/change_license_updated_date", ChangeLicenseUpdateDateHandler).Name("changeLicenseUpdateDatePath")
		admin.POST("/change_license_updated_date/{license_id}", ChangeLicenseUpdateDatePostHandler).Name("changeLicenseUpdateDatePostPath")
		admin.GET("/report", AdminReportHandler).Name("adminReportPath")
		admin.GET("/settings/texts", AdminAllTextsSettingsHandler).Name("adminAllTextsSettingsPath")
		admin.GET("/settings/text/{setting_id}", AdminTextSettingsHandler).Name("adminTextSettingsPath")
		admin.POST("/settings/text/{setting_id}", AdminTextSettingsPostHandler).Name("adminTextSettingsPostPath")
		//end admin route

		////if app has any problem, these lines are the reason
		//app.GET("/{antivirus_name}", IndexLicense).Name("thisIsJustForRedirect") // should enable it at the end//no it makes un useful pages on wrong urls!!! like localhost:3001/b !!!
		////

		app.Resource("/antis", AntisResource{})
		app.Resource("/licenses", LicensesResource{})
		app.ServeFiles("/", assetsBox) // serve files from the public directory
	}

	return app
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(packr.New("app:locales", "../locales"), "en-US"); err != nil {
		app.Stop(err)
	}
	return T.Middleware()
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
