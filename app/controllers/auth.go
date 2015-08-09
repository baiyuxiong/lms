package controllers

import (
	"github.com/revel/revel"
	"github.com/baiyuxiong/lms/app/utils"
	"github.com/baiyuxiong/lms/app/models"
	"github.com/baiyuxiong/lms/app"
	"time"
	"fmt"
)

type Auth struct {
	BaseController
}

func (c Auth) Login() revel.Result {
	return c.Render()
}

func (c Auth) Logout() revel.Result {
	c.Session[utils.SESSION_KEY_UID] = ""
	return c.Redirect("/auth/login")
}

func (c Auth) DoLogin(username, password string) revel.Result {
	revel.INFO.Println("DoLogin")

	c.Validation.Required(username).Message("用户名不能为空")
	c.Validation.Required(password).Message("密码不能为空")

	if c.Validation.HasErrors() {
		c.Flash.Error(utils.ValidationErrorToString(c.Validation.Errors))
		c.Validation.Keep()
		c.FlashParams()
		revel.INFO.Println("DoLogin Validation error")
		return c.Redirect("/auth/login")
	}else{
		errorMsg := "";

		var u = &models.Users{Username: username}
		has, _ := app.Engine.Get(u)


		if !has {
			errorMsg = "用户名或密码出错"
		}

		if utils.EncryptPassword(u.Salt, password) != u.Password {
			errorMsg = "用户名或密码出错"
		}

		if len(errorMsg) >0 {
			c.Flash.Error(errorMsg)
			revel.INFO.Println("DoLogin Validation error - ",errorMsg)
			c.Validation.Keep()
			c.FlashParams()
			return c.Redirect("/auth/login")
		}

		c.Session[utils.SESSION_KEY_UID] = fmt.Sprintf("%d",u.Id)

		u.UpdatedAt = time.Now()
		app.Engine.Id(u.Id).Cols("updated_at").Update(u)

		if username == utils.ADMIN_USERNAME{
			return c.Redirect("/admin")
		}else{
			return c.Redirect("/")
		}
	}
}

