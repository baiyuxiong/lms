package controllers

import (
	"github.com/revel/revel"
	"github.com/baiyuxiong/lms/app/models"
	"github.com/baiyuxiong/lms/app/utils"
	"github.com/baiyuxiong/lms/app/service"
	"github.com/baiyuxiong/lms/app"
	"time"
	"strconv"
	"errors"
	"strings"
	"fmt"
)

type App struct {
	BaseController
}

func (c App) ListLeave() revel.Result {
	leaves := make([]models.Leave, 0)
	app.Engine.Where("user_id=?", c.User.Id).OrderBy("updated_at desc").Find(&leaves)
	leaveType := utils.LeaveType
	return c.Render(leaves,leaveType)
}

func (c App) ViewLeave(id int)revel.Result{
	leave := new(models.Leave)
	app.Engine.Id(id).Get(leave)

	dept := new(models.Department)
	app.Engine.Id(c.UserProfile.DepartmentId).Get(dept)

	if leave.UserId != c.User.Id {
		return c.Redirect("/app/listLeave")
	}
	transfers := make([]models.LeaveTransfer,0)
	app.Engine.Where("leave_id=?",id).OrderBy("created_at asc").Find(&transfers)

	//获取每个参与审批人的信息
	transferUsers := make(map[int]models.UserProfiles)
	for _, transfer := range transfers {
		_, exists := transferUsers[transfer.AssignTo]
		if !exists {
			userProfile := new(models.UserProfiles)
			has, _ := app.Engine.Id(transfer.AssignTo).Get(userProfile)
			if has {
				transferUsers[transfer.AssignTo] =*userProfile
			}
		}
	}

	userProfile := c.UserProfile
	userLevel := utils.UserLevel
	userGender := utils.UserGender
	leaveType := utils.LeaveType
	employmentType:=utils.EmploymentType

	return c.Render(leave,transfers,transferUsers,userProfile,userLevel,userGender,dept,employmentType,leaveType)
}

func (c App) AddLeave() revel.Result {
	leaveType := utils.LeaveType

	return c.Render(leaveType)
}

func (c App) DoAddLeave(reason string, leave_type int, startdate, enddate time.Time, address string, onway_date int, filepath string) revel.Result {

	leave := new(models.Leave)
	leave.UserId = c.User.Id
	leave.Status = utils.LEAVE_STATUS_ASKING
	leave.LeaveType = leave_type
	leave.Reason = reason
	leave.Address = address
	leave.OnwayDate = onway_date
	leave.Filepath = filepath
	leave.Startdate = startdate
	leave.Enddate = enddate
	leave.Duration = fmt.Sprintf("%.1f",enddate.Sub(startdate).Seconds()/86400)
	leave.CreatedAt = time.Now()
	leave.UpdatedAt = leave.CreatedAt
	revel.INFO.Println("DoAddLeave",startdate,enddate)

	leave.FlowId = c.getLeaveFlowId(c.UserProfile, leave)
	leave.StepId = "0" //第0步，
	err :=c.getLeaveInchargeProfile(c.UserProfile,leave)

	if err != nil{
		c.Flash.Error(err.Error())
		c.FlashParams()
		return c.Redirect("/app/addLeave")
	}

	app.Engine.Insert(leave)

	transfer := new(models.LeaveTransfer)
	transfer.LeaveId = leave.Id
	transfer.AssignFr = leave.UserId
	transfer.AssignTo = leave.InChargeUserId
	transfer.IsAgree = 0
	transfer.CreatedAt = time.Now()
	transfer.UpdatedAt = transfer.CreatedAt

	app.Engine.Insert(transfer)
	//获取处理负责人

	return c.Redirect("/app/listLeave")
}

func (c App) CheckLeave() revel.Result {
	
	return c.Render()
}

func (c App) DoCheckLeave() revel.Result {
	return c.Render()
}

func (c App) EditProfile() revel.Result {
	user := c.User

	userProfiles := c.UserProfile

	dept := &models.Department{}
	app.Engine.Id(userProfiles.DepartmentId).Get(dept)

	employmentType, exists := utils.EmploymentType[userProfiles.EmploymentType]
	if !exists {
		employmentType = ""
	}

	return c.Render(user, userProfiles, dept, employmentType)
}

func (c App) DoEditProfile(password, phone, phone1, email string) revel.Result {

	service.EditSelf(c.User.Id, password, phone, phone1, email)

	return c.Redirect("/app/editProfile")
}

func (c App)getLeaveFlowId(up *models.UserProfiles, leave *models.Leave) string {
	revel.INFO.Println("getLeaveFlowId duration:",leave.Duration,",gender:",up.Gender,",LeaveType:",leave.LeaveType,",level:",up.Level)
	//一般员工
	duration,_ := strconv.Atoi(leave.Duration)
	if up.Level == 1 {
		if  duration<= 1 {
			return "1"; //普通员工一天之内
		}else {
			if up.Gender == 2 && (leave.LeaveType == 4 || leave.LeaveType == 6) {//女
				if duration > 1 && duration <=3 {
					return "6"
				}else{
					return "7"
				}
			}else {
				if duration > 1 && duration <=3 {
					return "2"
				}else {
					return "3"
				}
			}
		}
	}else {//科级员工
		if up.Gender == 2 && (leave.LeaveType == 4 || leave.LeaveType == 6) {//女
			if duration <=1 {
				return "8"
			}else {
				return "9"
			}
		}else {
			if duration <=1 {
				return "4"
			}else {
				return "5"
			}
		}
	}
	return ""
}
/**
根据假条，获取下一审批人ID，第一个参数为发起人信息
 */
func (c App)getLeaveInchargeProfile(up *models.UserProfiles,leave *models.Leave) error{
	currentStepId,err := strconv.Atoi(leave.StepId)
	if err != nil{
		return err
	}

	nextStepId := currentStepId + 1
	flowId := leave.FlowId

	steps,exists := app.Flows[flowId]
	if !exists{
		return errors.New("处理流程"+flowId+"不存在")
	}

	nextStep,exists := steps.Steps[fmt.Sprintf("%d",nextStepId)]
	if !exists{
		return errors.New("处理已结束")
	}

	leave.InChargeUserId,err = c.getLeaveInchargeUserId(up,flowId,fmt.Sprintf("%d",nextStepId))
	if err != nil{
		return err
	}

	//当前一步，维护状态用
	currentStep,exists := steps.Steps[fmt.Sprintf("%d",currentStepId)]

	if currentStepId!=0 && !exists{
		leave.Status = utils.LEAVE_STATUS_CANCEL_OK //销假成功，结束了
	}else if nextStep.Type=="2" && currentStep.Type=="1"{
		leave.Status = utils.LEAVE_STATUS_ASK_OK //请假成功，该消假了
	}
	return nil
}

func (c App)getLeaveInchargeUserId(up *models.UserProfiles,flowId,stepId string) (int,error) {
	steps,exists := app.Flows[flowId]
	if !exists{
		revel.INFO.Println("getLeaveInchargeUserId , flowId",flowId,"not exists")
		return 0,errors.New("流程"+flowId+"不存在")
	}

	step,exists := steps.Steps[stepId]
	if !exists{
		revel.INFO.Println("getLeaveInchargeUserId , flowId",flowId,"stepId",stepId,"not exists")
		return 0,errors.New("步骤"+flowId+"-"+stepId+"不存在")
	}

	charge := step.Charge

	if charge=="starter"{
		return up.UserId,nil
	}
	if charge == "starterDept"{
		department := new(models.Department)
		has,_ := app.Engine.Id(up.DepartmentId).Get(department)
		if !has{
			return 0,errors.New(fmt.Sprintf("部门%d不存在",up.DepartmentId))
		}
		return department.LeaderId,nil
	}

	if strings.HasPrefix(charge,"deptById"){
		splits := strings.Split(charge,":")
		deptId := splits[1]

		department := new(models.Department)
		app.Engine.Id(deptId).Get(department)

		return department.LeaderId,nil
	}

	if strings.HasPrefix(charge,"userByUsername"){
		splits := strings.Split(charge,":")
		username := splits[1]

		user := new(models.Users)
		has,_ := app.Engine.Where("username=?",username).Get(user)

		if !has{
			return 0,errors.New("用户"+username+"不存在")
		}
		return user.Id,nil
	}

	if strings.HasPrefix(charge,"deptParentByStep"){
		splits := strings.Split(charge,":")
		parentStepId := splits[1]
		parentUserId,err := c.getLeaveInchargeUserId(up,flowId,parentStepId)//上一级的用户ID
		if err != nil {
			return 0,err
		}

		//根据parentUserId查询负责的部门
		department := new(models.Department)
		has,_ := app.Engine.Where("leader_id=?",parentUserId).Get(department)
		if !has{
			revel.INFO.Println("parentUserId",parentUserId,"is not leader of department")
			return 0,errors.New(fmt.Sprintf("用户%d不是任何部门的负责人",parentUserId))
		}
		parentDepartment := new(models.Department)
		has,_ = app.Engine.Id(department.ParentId).Get(parentDepartment)
		if !has{
			revel.INFO.Println("No department of id ",department.ParentId,"found")
			return 0,errors.New(fmt.Sprintf("部门%d不存在",department.ParentId))
		}
		return parentDepartment.LeaderId,nil
	}

	return 0,nil
}