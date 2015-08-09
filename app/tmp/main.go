// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	_ "github.com/baiyuxiong/lms/app"
	controllers "github.com/baiyuxiong/lms/app/controllers"
	tests "github.com/baiyuxiong/lms/tests"
	controllers1 "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers0 "github.com/revel/modules/testrunner/app/controllers"
	time "time"
	"github.com/revel/revel/testing"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.BaseController)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Before",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					70: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					107: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Admin)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "ListUser",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "limit", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "offset", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					73: []string{ 
						"users",
						"profiles",
						"depts",
						"pages",
						"userLevel",
					},
				},
			},
			&revel.MethodType{
				Name: "AddUser",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					85: []string{ 
						"depts",
						"userLevel",
						"employmentType",
					},
				},
			},
			&revel.MethodType{
				Name: "DoAddUser",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "name", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "phone", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "phone1", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "position", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "email", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "department_id", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "employment_type", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "gender", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "level", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "is_leader", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "join_date", Type: reflect.TypeOf((*time.Time)(nil)) },
					&revel.MethodArg{Name: "birth_date", Type: reflect.TypeOf((*time.Time)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "EditUser",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					150: []string{ 
						"user",
						"userProfiles",
						"depts",
						"isDepartmentLeader",
						"userLevel",
						"employmentType",
					},
				},
			},
			&revel.MethodType{
				Name: "DoEditUser",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "name", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "phone", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "phone1", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "position", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "email", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "department_id", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "employment_type", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "gender", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "level", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "is_leader", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "join_date", Type: reflect.TypeOf((*time.Time)(nil)) },
					&revel.MethodArg{Name: "birth_date", Type: reflect.TypeOf((*time.Time)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DelUser",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ListDepartment",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					223: []string{ 
						"depts",
						"leaders",
						"nameMap",
					},
				},
			},
			&revel.MethodType{
				Name: "AddDepartment",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					231: []string{ 
						"depts",
					},
				},
			},
			&revel.MethodType{
				Name: "DoAddDepartment",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "name", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "info", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "parent_id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "EditDepartment",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					276: []string{ 
						"department",
						"depts",
					},
				},
			},
			&revel.MethodType{
				Name: "DoEditDepartment",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "name", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "info", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "parent_id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DelDepartment",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Import",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "SetDepartmentLeader",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.App)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "ListLeave",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					24: []string{ 
						"leaves",
						"leaveType",
					},
				},
			},
			&revel.MethodType{
				Name: "ViewLeave",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					59: []string{ 
						"leave",
						"transfers",
						"transferUsers",
						"userProfile",
						"userLevel",
						"userGender",
						"dept",
						"employmentType",
						"leaveType",
					},
				},
			},
			&revel.MethodType{
				Name: "AddLeave",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					65: []string{ 
						"leaveType",
					},
				},
			},
			&revel.MethodType{
				Name: "DoAddLeave",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "reason", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "leave_type", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "startdate", Type: reflect.TypeOf((*time.Time)(nil)) },
					&revel.MethodArg{Name: "enddate", Type: reflect.TypeOf((*time.Time)(nil)) },
					&revel.MethodArg{Name: "address", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "onway_date", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "CheckLeave",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					113: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "DoCheckLeave",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					117: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "EditProfile",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					133: []string{ 
						"user",
						"userProfiles",
						"dept",
						"employmentType",
					},
				},
			},
			&revel.MethodType{
				Name: "DoEditProfile",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "phone", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "phone1", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "email", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Auth)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Login",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					17: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Logout",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DoLogin",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Upload)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "DoUploadFile",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"github.com/baiyuxiong/lms/app/controllers.Admin.DoAddDepartment": { 
			236: "name",
		},
		"github.com/baiyuxiong/lms/app/controllers.Admin.DoAddUser": { 
			90: "username",
			91: "password",
			92: "name",
		},
		"github.com/baiyuxiong/lms/app/controllers.Admin.DoEditUser": { 
			155: "username",
			156: "name",
		},
		"github.com/baiyuxiong/lms/app/controllers.Auth.DoLogin": { 
			28: "username",
			29: "password",
		},
	}
	testing.TestSuites = []interface{}{ 
		(*tests.AppTest)(nil),
	}

	revel.Run(*port)
}
