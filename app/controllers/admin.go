package controllers

import (
	"github.com/revel/revel"
	"github.com/baiyuxiong/lms/app/models"
	"github.com/baiyuxiong/lms/app/service"
	"github.com/baiyuxiong/lms/app"
	"github.com/baiyuxiong/lms/app/utils"
	"time"
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
)

type Admin struct {
	BaseController
}

/**
用户管理
 */
func (c Admin) ListUser(limit, offset int) revel.Result {

	if limit <= 0 {
		limit = utils.PERPAGE;
	}
	if offset < 0 {
		offset = 0;
	}


	user := new(models.Users)
	totalCount, err := app.Engine.Where("is_delete=?", 0).Count(user)
	if err != nil {
		return nil
	}

	users := make([]models.Users, 0);
	err = app.Engine.Where("is_delete=?", 0).Limit(limit, offset).Find(&users)
	if err != nil {
		return nil
	}

	userLevel := utils.UserLevel

	profiles := make(map[int]models.UserProfiles)
	depts := make(map[int]models.Department)

	for _, user := range users {
		_, exists := profiles[user.Id]
		if !exists {
			userProfile := new(models.UserProfiles)
			has, _ := app.Engine.Id(user.Id).Get(userProfile)
			if has {
				profiles[user.Id] =*userProfile
			}
		}

		_, exists = depts[user.Id]
		if !exists {
			dept := new(models.Department)
			has, _ := app.Engine.Id(user.Id).Get(dept)
			if has {
				depts[user.Id] =*dept
			}
		}
	}

	pages := utils.Pagination(limit, offset, totalCount)

	return c.Render(users, profiles, depts, pages, userLevel)
}

func (c Admin) AddUser() revel.Result {

	userLevel := utils.UserLevel
	employmentType := utils.EmploymentType

	//部门列表
	depts := make([]models.Department, 0);
	app.Engine.Where("is_delete=?", 0).Find(&depts)

	return c.Render(depts, userLevel, employmentType)
}

func (c Admin) DoAddUser(username, password, name, phone, phone1, position, email string, department_id, employment_type, gender, level, is_leader int, join_date, birth_date time.Time) revel.Result {

	c.Validation.Required(username).Message("用户名不能为空")
	c.Validation.Required(password).Message("密码不能为空")
	c.Validation.Required(name).Message("姓名不能为空")

	if c.Validation.HasErrors() {
		c.Flash.Error(utils.ValidationErrorToString(c.Validation.Errors))
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/admin/addUser")
	}

	//判断部门有没有领导
	if is_leader == 1 {
		dept := &models.Department{}
		app.Engine.Id(department_id).Get(dept)
		if dept.LeaderId != 0 {
			up := &models.UserProfiles{}
			app.Engine.Id(dept.LeaderId).Get(up)

			c.Flash.Error("该部门当前审批负责人为"+up.Name+"，一个部门只能指定一个审批负责人")
			c.Validation.Keep()
			c.FlashParams()
			return c.Redirect("/admin/addUser")
		}
	}

	err := service.AddUser(username, password, department_id, employment_type, gender, level, name, phone, phone1, position, email, join_date, birth_date, c.Request.RemoteAddr)
	if err != nil {
		c.Flash.Error(err.Error())
		return c.Redirect("/admin/addUser")
	}

	return c.Redirect("/admin/listUser")
}


func (c Admin) EditUser(id int) revel.Result {

	//部门列表
	depts := make([]models.Department, 0);
	app.Engine.Where("is_delete=?", 0).Find(&depts)

	user := &models.Users{}
	app.Engine.Id(id).Get(user)

	userProfiles := &models.UserProfiles{}
	app.Engine.Id(id).Get(userProfiles)

	userLevel := utils.UserLevel
	employmentType := utils.EmploymentType

	isDepartmentLeader := false
	for _, v := range depts {
		if v.Id ==userProfiles.DepartmentId {
			if v.LeaderId == user.Id {
				isDepartmentLeader = true
			}
		}
	}

	return c.Render(user, userProfiles, depts, isDepartmentLeader, userLevel, employmentType)
}

func (c Admin) DoEditUser(id int, username, password, name, phone, phone1, position, email string, department_id, employment_type, gender, level, is_leader int, join_date, birth_date time.Time) revel.Result {

	c.Validation.Required(username).Message("用户名不能为空")
	c.Validation.Required(name).Message("姓名不能为空")

	if c.Validation.HasErrors() {
		c.Flash.Error(utils.ValidationErrorToString(c.Validation.Errors))
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/admin/editUser/?id="+fmt.Sprintf("%d", id))
	}

	//判断部门有没有领导
	if is_leader == 1 {
		dept := &models.Department{}
		app.Engine.Id(department_id).Get(dept)
		if dept.LeaderId != 0 && dept.LeaderId != c.User.Id {
			up := &models.UserProfiles{}
			app.Engine.Id(dept.LeaderId).Get(up)

			c.Flash.Error("该部门当前审批负责人为"+up.Name+"，一个部门只能指定一个审批负责人")
			c.Validation.Keep()
			c.FlashParams()
			return c.Redirect("/admin/editUser?id="+fmt.Sprintf("%d", id))
		}
	}

	service.EditUser(id, username, password, department_id, employment_type, gender, level, name, phone, phone1, position, email, join_date, birth_date, c.Request.RemoteAddr)

	return c.Redirect("/admin/listUser")
}

func (c Admin) DelUser(id int) revel.Result {

	if id != 1 {
		u := new(models.Users)
		u.IsDelete = 1
		app.Engine.Id(id).Cols("is_delete").Update(u)
	}

	return c.Redirect("/admin/ListUser")
}

/**
部门管理
 */

func (c Admin) ListDepartment() revel.Result {

	depts := make([]models.Department, 0);
	err := app.Engine.Where("is_delete=?", 0).OrderBy("depth asc").Find(&depts)
	if err != nil {
		return nil
	}

	//获取每个部门负责人的信息
	leaders := make(map[int]models.UserProfiles)
	//部门名称列表
	nameMap := make(map[int]string)
	for _, dept := range depts {
		nameMap[dept.Id] = dept.Name
		_, exists := leaders[dept.LeaderId]
		if !exists {
			userProfile := new(models.UserProfiles)
			has, _ := app.Engine.Id(dept.LeaderId).Get(userProfile)
			if has {
				leaders[dept.LeaderId] =*userProfile
			}
		}
	}
	return c.Render(depts, leaders, nameMap)
}

func (c Admin) AddDepartment() revel.Result {

	depts := make([]models.Department, 0);
	app.Engine.Where("is_delete=?", 0).OrderBy("depth asc").Find(&depts)

	return c.Render(depts)
}

func (c Admin) DoAddDepartment(name, info string, parent_id int) revel.Result {

	c.Validation.Required(name).Message("名称不能为空")

	if c.Validation.HasErrors() {
		c.Flash.Error(utils.ValidationErrorToString(c.Validation.Errors))
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect("/admin/addDepartment")
	}

	//获取上一级的depth，本级加1
	depth := 0
	if parent_id > 0 {
		department := &models.Department{}
		app.Engine.Id(parent_id).Get(department)
		if nil != department {
			depth = department.Depth
		}
	}

	l := new(models.Department)
	l.Name = name
	l.Info = info
	l.Depth = (depth+1)
	l.ParentId = parent_id
	l.UpdatedAt = time.Now()
	l.CreatedAt = time.Now()

	app.Engine.Insert(l)
	return c.Redirect("/admin/listDepartment")
}


func (c Admin) EditDepartment(id int) revel.Result {

	department := &models.Department{}
	app.Engine.Id(id).Get(department)

	depts := make([]models.Department, 0);
	app.Engine.Where("is_delete=?", 0).OrderBy("depth asc").Find(&depts)

	return c.Render(department, depts)
}

func (c Admin) DoEditDepartment(id int, name, info string, parent_id int) revel.Result {
	//获取上一级的depth，本级加1
	depth := 0
	if parent_id > 0 {
		department := &models.Department{}
		app.Engine.Id(parent_id).Get(department)
		if nil != department {
			depth = department.Depth
		}
	}

	d := new(models.Department)
	d.Depth = (depth+1)
	d.ParentId = parent_id
	d.Name = name
	d.Info = info
	app.Engine.Id(id).Cols("name").Cols("info").Cols("parent_id").Cols("depth").Update(d)

	return c.Redirect("/admin/listDepartment")
}

func (c Admin) DelDepartment(id int) revel.Result {
	l := new(models.Department)
	l.IsDelete = 1
	app.Engine.Id(id).Cols("is_delete").Update(l)
	return c.Redirect("/admin/ListDepartment")
}

func (c Admin) Import() revel.Result {
	if (!utils.DATA_IMPORTED) {
		dataFilePath := "/Users/baiyuxiong/Desktop/users.csv"
		file, err := os.Open(dataFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		line := 1
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if line != 1 {
				fmt.Println(scanner.Text())

				cells := strings.Split(scanner.Text(), ",")
				if len(cells) == 11 {
					employment_type := 0;
					for k, v := range utils.EmploymentType {
						if cells[2] == v {
							employment_type = k
							break;
						}
					}

					gender := 0
					for k, v := range utils.UserGender {
						if cells[6] == v {
							gender = k
							break;
						}
					}

					level := 1
					if cells[5] != "一般员工" {
						level = 2
					}

					department_id := 0
					depts := make([]models.Department, 0)
					app.Engine.Find(&depts)
					for _, v := range depts {
						if cells[4] == v.Name {
							department_id = v.Id
							break;
						}
					}

					if department_id == 0 {
						dept := new(models.Department)
						dept.ParentId = 0
						dept.Depth = 2
						if cells[4] == "领导" {
							if cells[5] == "经理" {
								dept.Depth = 0
							}else {
								dept.Depth = 1
							}
						}
						dept.LeaderId = 0
						dept.Name = cells[4]
						dept.Info = cells[4]
						dept.UpdatedAt = time.Now()
						dept.CreatedAt = time.Now()

						app.Engine.Insert(dept)
						department_id = dept.Id
					}

					join_date, _ := time.Parse("2006.01.02", cells[7])
					birth_date, _ := time.Parse("2006.01.02", cells[8])
					service.AddUser(cells[1], cells[1], department_id, employment_type, gender, level, cells[3], "", "", cells[5], "", join_date, birth_date, "")
				}
			}
			line++
		}



		if err := scanner.Err(); err != nil {
			return c.RenderText(err.Error())
		}
		return c.RenderText("ok")
	}else{
		return c.RenderText("Imported")
	}
}

func (c Admin) SetDepartmentLeader() revel.Result {
	//将主任改为部门审批负责人
	users := make([]models.UserProfiles,0)
	app.Engine.Where("position=?","主任").Find(&users)
	for _,u := range users {
		department := &models.Department{LeaderId:u.UserId}
		app.Engine.Id(u.DepartmentId).Cols("leader_id").Update(department)
	}
	return c.RenderText("ok")
}
