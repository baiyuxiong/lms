package tests

import (
	"github.com/revel/revel/testing"
	"github.com/baiyuxiong/lms/app"
)

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) TestInsertAdmin() {
	t.ClearUserTable();
	t.ClearUserProfileTable();

	//service.AddUser(utils.ADMIN_USERNAME,utils.ADMIN_PASSWORD ,0 ,"admin",time.Now(),"127.0.0.1")
}

func (t *AppTest) ClearUserTable() {
	sql := "truncate table users"
	_, err := app.Engine.Exec(sql)
	t.AssertEqual(nil, err)
}
func (t *AppTest) ClearUserProfileTable() {
	sql := "truncate table user_profiles"
	_, err := app.Engine.Exec(sql)
	t.AssertEqual(nil, err)
}