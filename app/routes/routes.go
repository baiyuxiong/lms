// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tBaseController struct {}
var BaseController tBaseController


func (_ tBaseController) Before(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("BaseController.Before", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tUpload struct {}
var Upload tUpload


func (_ tUpload) DoUploadFile(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Upload.DoUploadFile", args).Url
}


type tAdmin struct {}
var Admin tAdmin


func (_ tAdmin) ListUser(
		limit int,
		offset int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "limit", limit)
	revel.Unbind(args, "offset", offset)
	return revel.MainRouter.Reverse("Admin.ListUser", args).Url
}

func (_ tAdmin) AddUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.AddUser", args).Url
}

func (_ tAdmin) DoAddUser(
		username string,
		password string,
		name string,
		phone string,
		phone1 string,
		position string,
		email string,
		department_id int,
		employment_type int,
		gender int,
		level int,
		is_leader int,
		join_date interface{},
		birth_date interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "name", name)
	revel.Unbind(args, "phone", phone)
	revel.Unbind(args, "phone1", phone1)
	revel.Unbind(args, "position", position)
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "department_id", department_id)
	revel.Unbind(args, "employment_type", employment_type)
	revel.Unbind(args, "gender", gender)
	revel.Unbind(args, "level", level)
	revel.Unbind(args, "is_leader", is_leader)
	revel.Unbind(args, "join_date", join_date)
	revel.Unbind(args, "birth_date", birth_date)
	return revel.MainRouter.Reverse("Admin.DoAddUser", args).Url
}

func (_ tAdmin) EditUser(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Admin.EditUser", args).Url
}

func (_ tAdmin) DoEditUser(
		id int,
		username string,
		password string,
		name string,
		phone string,
		phone1 string,
		position string,
		email string,
		department_id int,
		employment_type int,
		gender int,
		level int,
		is_leader int,
		join_date interface{},
		birth_date interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "name", name)
	revel.Unbind(args, "phone", phone)
	revel.Unbind(args, "phone1", phone1)
	revel.Unbind(args, "position", position)
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "department_id", department_id)
	revel.Unbind(args, "employment_type", employment_type)
	revel.Unbind(args, "gender", gender)
	revel.Unbind(args, "level", level)
	revel.Unbind(args, "is_leader", is_leader)
	revel.Unbind(args, "join_date", join_date)
	revel.Unbind(args, "birth_date", birth_date)
	return revel.MainRouter.Reverse("Admin.DoEditUser", args).Url
}

func (_ tAdmin) DelUser(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Admin.DelUser", args).Url
}

func (_ tAdmin) ListDepartment(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.ListDepartment", args).Url
}

func (_ tAdmin) AddDepartment(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.AddDepartment", args).Url
}

func (_ tAdmin) DoAddDepartment(
		name string,
		info string,
		parent_id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "name", name)
	revel.Unbind(args, "info", info)
	revel.Unbind(args, "parent_id", parent_id)
	return revel.MainRouter.Reverse("Admin.DoAddDepartment", args).Url
}

func (_ tAdmin) EditDepartment(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Admin.EditDepartment", args).Url
}

func (_ tAdmin) DoEditDepartment(
		id int,
		name string,
		info string,
		parent_id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "name", name)
	revel.Unbind(args, "info", info)
	revel.Unbind(args, "parent_id", parent_id)
	return revel.MainRouter.Reverse("Admin.DoEditDepartment", args).Url
}

func (_ tAdmin) DelDepartment(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Admin.DelDepartment", args).Url
}

func (_ tAdmin) Import(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.Import", args).Url
}

func (_ tAdmin) SetDepartmentLeader(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Admin.SetDepartmentLeader", args).Url
}


type tApp struct {}
var App tApp


func (_ tApp) ListLeave(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.ListLeave", args).Url
}

func (_ tApp) ViewLeave(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.ViewLeave", args).Url
}

func (_ tApp) AddLeave(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.AddLeave", args).Url
}

func (_ tApp) DoAddLeave(
		reason string,
		leave_type int,
		startdate interface{},
		enddate interface{},
		address string,
		onway_date int,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "reason", reason)
	revel.Unbind(args, "leave_type", leave_type)
	revel.Unbind(args, "startdate", startdate)
	revel.Unbind(args, "enddate", enddate)
	revel.Unbind(args, "address", address)
	revel.Unbind(args, "onway_date", onway_date)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("App.DoAddLeave", args).Url
}

func (_ tApp) CheckLeave(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.CheckLeave", args).Url
}

func (_ tApp) ViewCheckDetail(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.ViewCheckDetail", args).Url
}

func (_ tApp) DoCheckLeave(
		id int,
		action string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "action", action)
	return revel.MainRouter.Reverse("App.DoCheckLeave", args).Url
}

func (_ tApp) EditProfile(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.EditProfile", args).Url
}

func (_ tApp) DoEditProfile(
		password string,
		phone string,
		phone1 string,
		email string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "phone", phone)
	revel.Unbind(args, "phone1", phone1)
	revel.Unbind(args, "email", email)
	return revel.MainRouter.Reverse("App.DoEditProfile", args).Url
}


type tAuth struct {}
var Auth tAuth


func (_ tAuth) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Auth.Login", args).Url
}

func (_ tAuth) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Auth.Logout", args).Url
}

func (_ tAuth) DoLogin(
		username string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("Auth.DoLogin", args).Url
}


