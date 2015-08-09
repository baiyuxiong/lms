package controllers

import (
	"github.com/revel/revel"
	"github.com/baiyuxiong/lms/app/utils"
	"github.com/baiyuxiong/lms/app/models"
	"github.com/baiyuxiong/lms/app"
	"strconv"
	"strings"
)

// init is called when the first request into the controller is made
func init() {
	revel.InterceptMethod((*BaseController).Before, revel.BEFORE)
}

type G struct {
	Path string
}

type BaseController struct {
	*revel.Controller
	User *models.Users
	UserProfile *models.UserProfiles
	IsAdmin bool
	g *G
}

// Before is called prior to the controller method
func (c *BaseController) Before() revel.Result {

	c.g = &G{}
	c.g.Path = c.Request.URL.Path

	c.IsAdmin = false;
	noTokenPath := []string{"/auth/login","/auth/doLogin"}
	if !utils.StringInSlice(c.Request.URL.Path,noTokenPath){

		uid := c.Session[utils.SESSION_KEY_UID]
		if len(uid) < 1{
			return c.Redirect("/auth/login")
		}else {
			u := &models.Users{}
			has, _ := app.Engine.Id(uid).Get(u)

			if has {
				c.User = u
				c.IsAdmin = (u.Username == utils.ADMIN_USERNAME)

				if !c.IsAdmin && strings.Contains(c.g.Path,"admin"){
					return c.Redirect("/")
				}

				uidInt,err := strconv.Atoi(uid)
				if err== nil{
					up := &models.UserProfiles{UserId:uidInt}
					app.Engine.Get(up)
					c.UserProfile = up
				}

			}else{
				return c.Redirect("/auth/login")
			}
		}
	}

	return nil
}

func (c *BaseController) GRender(extraRenderArgs ...interface{}) revel.Result{
	g := c.g
	extraRenderArgs = append(extraRenderArgs,g)
	return  c.Render(extraRenderArgs...)
}