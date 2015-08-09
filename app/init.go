package app

import (
	"github.com/revel/revel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
	"fmt"
	"github.com/baiyuxiong/lms/app/utils"
	"github.com/toolkits/file"
	"encoding/json"
	"os"
)

var Engine *xorm.Engine

var Flows map[string]utils.Flow

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter, // Recover from panics and display an error page instead.
		revel.RouterFilter, // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter, // Parse parameters into Controller.Params.
		revel.SessionFilter, // Restore and write the session cookie.
		revel.FlashFilter, // Restore and write the flash cookie.
		revel.ValidationFilter, // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter, // Resolve the requested language
		HeaderFilter, // Add some security based headers
		revel.InterceptorFilter, // Run interceptors around the action.
		revel.CompressFilter, // Compress the result.
		revel.ActionInvoker, // Invoke the action.
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	// revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
	revel.OnAppStart(InitDB)
	revel.OnAppStart(InitFlows)

	revel.TemplateFuncs["formatDate"] = func(date time.Time, format string) string {
		return date.Format(format)
	}
	revel.TemplateFuncs["leaveStatus"] = func(status int) string {
		if status == utils.LEAVE_STATUS_ASKING{
			return "请假中"
		}else if status == utils.LEAVE_STATUS_ASK_OK{
			return "请假成功"
		}else if status == utils.LEAVE_STATUS_ASK_FAIL{
			return "请假失败"
		}else if status == utils.LEAVE_STATUS_CANCELING{
			return "销假中"
		}else if status == utils.LEAVE_STATUS_CANCEL_OK{
			return "销假成功"
		}else if status == utils.LEAVE_STATUS_CANCEL_FAIL{
			return "销假失败"
		}
		return "未知状态，出错"
	}

	revel.TemplateFuncs["checkStatus"] = func(status int) string {
		if status == utils.LEAVE_CHECKING{
			return "审批中"
		}else if status == utils.LEAVE_CHECK_OK{
			return "已批准"
		}else if status == utils.LEAVE_CHECK_FAIL{
			return "未批准"
		}
		return "未知状态，出错"
	}

	revel.TemplateFuncs["addSpaces"] = func(count int,times int) string {
		str := ""
		total := count*times
		for ; total>0; total-- {
			str += "&nbsp;"
		}
		return str
	}
	revel.TemplateFuncs["myeq"] = func(a interface{}, b interface{}) (result bool) {
		if nil == a && nil == b{
			result = true;
		} else if a== nil || b==nil{
			result = false
		} else{
			aStr := "";
			switch x := a.(type) {
			case int:
				aStr = fmt.Sprintf("%d",x)
			case float64:
				aStr = fmt.Sprintf("%f",x)
			default:
				aStr = a.(string)
			}

			bStr := "";
			switch x := b.(type) {
			case int:
				bStr = fmt.Sprintf("%d",x)
			case float64:
				bStr = fmt.Sprintf("%f",x)
			default:
				bStr = b.(string)
			}
			result = (aStr == bStr)
		}
		return
	}
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

func InitDB() {
	dns := revel.Config.StringDefault("db.dns", "root:@tcp(127.0.0.1:3306)/lms?charset=utf8")

	var err error
	Engine, err = xorm.NewEngine("mysql", dns)
	if err != nil {
		revel.ERROR.Fatalln("InitDB - NewEngine error ", err)
	}
	err = Engine.Ping()
	if err != nil {
		revel.ERROR.Fatalln("InitDB - Ping error ", err)
	}
	Engine.ShowSQL = true
	Engine.ShowErr = true
	Engine.ShowDebug = true
	Engine.ShowWarn = true
	Engine.ShowInfo = true

	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 5000)
	Engine.SetDefaultCacher(cacher)

	go pingDB()
}

func pingDB() {
	for {
		Engine.Ping()
		time.Sleep(time.Minute * 10) //每10分钟ping一下，保持连接有效
	}
}

func InitFlows() {
	cfg := revel.BasePath+"/conf/flow.json"

	if !file.IsExist(cfg) {
		file, _ := os.Getwd()
		revel.ERROR.Fatalln("config file:", cfg, "is not existent. getcwd",file)
	}
	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		revel.ERROR.Fatalln("read config file:", cfg, "fail:", err)
	}

	err = json.Unmarshal([]byte(configContent), &Flows)
	if err != nil {
		revel.ERROR.Fatalln("parse config file:", cfg, "fail:", err)
	}else {
		revel.INFO.Println("InitFlows ok")
	}
}
