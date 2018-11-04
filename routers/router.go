package routers

import (
	"github.com/astaxie/beego"
	"tc.web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/cart", &controllers.CartController{})
	beego.Router("/categories", &controllers.CategoriesController{})
	beego.Router("/checkout", &controllers.CheckoutController{})
	beego.Router("/contact", &controllers.ContactController{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/product", &controllers.ProductController{})
}
