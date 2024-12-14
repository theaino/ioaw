package routers

import (
	"ioaw/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/register", &controllers.UserController{}, "get:RegisterForm;post:Register")
	web.Router("/login", &controllers.UserController{}, "get:LoginForm;post:Login")
	web.Router("/logout", &controllers.UserController{}, "get:Logout")

	web.Router("/articles", &controllers.ArticleController{}, "get:List")
	web.Router("/articles/:id", &controllers.ArticleController{}, "get:View")
	web.Router("/articles/create", &controllers.ArticleController{}, "get:New;post:Create")
	web.Router("/articles/edit/:id", &controllers.ArticleController{}, "get:Edit;post:Update")
	web.Router("/articles/delete/:id", &controllers.ArticleController{}, "get:Delete")

	web.Router("/", &controllers.DashboardController{}, "get:Index")
}
