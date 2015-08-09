package utils

//common
const (
	SESSION_KEY_UID = "uid" //session存储用户ID键
	ADMIN_USERNAME = "admin" //管理员用户名
	ADMIN_PASSWORD = "admin" //管理员密码

	PERPAGE = 20 //分页
	DATA_IMPORTED = false

	LEAVE_STATUS_ASKING = 1
	LEAVE_STATUS_ASK_OK = 2
	LEAVE_STATUS_ASK_FAIL = 3

	LEAVE_STATUS_CANCELING = 4
	LEAVE_STATUS_CANCEL_OK = 5
	LEAVE_STATUS_CANCEL_FAIL = 6

	LEAVE_CHECKING = 0
	LEAVE_CHECK_OK = 1
	LEAVE_CHECK_FAIL=2

)

//请假类型
var LeaveType = map[int]string{
	1 : "事假",
	2 : "病假",
	3 : "年休假",
	4 : "婚假",
	5 : "丧假",
	6 : "产假",
	7 : "探亲假",
	8 : "因公外出",
	9 : "学习",
}

//员工性别
var UserGender = map[int]string{
	1 : "男",
	2 : "女",
}

//员工级别
var UserLevel = map[int]string{
	1 : "一般员工",
	2 : "科级管理人员",
}

//用工性质
var EmploymentType = map[int]string{
	1 : "长期合同工",
	2 : "劳务工",
}