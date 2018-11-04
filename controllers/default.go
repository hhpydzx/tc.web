package controllers

import (
	"github.com/astaxie/beego"
)

// MainController struct
type MainController struct {
	beego.Controller
}

// Get func
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// ProductController struct
type ProductController struct {
	beego.Controller
}

// Get func
func (c *ProductController) Get() {
	c.TplName = "product.html"
}

// IndexController struct
type IndexController struct {
	beego.Controller
}

// Get func
func (c *IndexController) Get() {
	c.TplName = "index.html"
}

// ContactController struct
type ContactController struct {
	beego.Controller
}

// Get func
func (c *ContactController) Get() {
	c.TplName = "contact.html"
}

// CheckoutController struct
type CheckoutController struct {
	beego.Controller
}

// Get func
func (c *CheckoutController) Get() {
	c.TplName = "checkout.html"
}

// CategoriesController struct
type CategoriesController struct {
	beego.Controller
}

// Get func
func (c *CategoriesController) Get() {
	c.TplName = "categories.html"
}

// CartController struct
type CartController struct {
	beego.Controller
}

// Get func
func (c *CartController) Get() {
	c.TplName = "cart.html"
}
