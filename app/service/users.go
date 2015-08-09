package service

import (
	"time"

	"github.com/baiyuxiong/lms/app/models"
	"github.com/baiyuxiong/lms/app"
	"github.com/baiyuxiong/lms/app/utils"
	"errors"
	"strings"
)


func AddUser(username, password string, department_id, employment_type, gender, level int, name, phone, phone1, position, email string, join_date,birth_date time.Time, ip_address string) error {
	var u = &models.Users{Username: username}
	has, _ := app.Engine.Get(u)

	if has {
		return errors.New("用户名已被注册，不可用")
	}

	now := time.Now()
	u.IpAddress = strings.Split(ip_address, ":")[0]
	u.Salt = utils.Salt()
	u.Password = utils.EncryptPassword(u.Salt, password)
	u.IsActivited = 1
	u.ActivatedAt = now
	u.CreatedAt = now
	u.UpdatedAt = now

	_, err := app.Engine.Insert(u)
	if err != nil {
		return errors.New("添加用户失败" + err.Error())
	}

	profile := new(models.UserProfiles)
	profile.UserId = u.Id
	profile.DepartmentId = department_id
	profile.Gender = gender
	profile.Level = level
	profile.EmploymentType = employment_type
	profile.Name = name
	profile.Phone = phone
	profile.Phone1 = phone1
	profile.Position = position
	profile.JoinDate = join_date
	profile.Email = email
	profile.BirthDate = birth_date
	profile.CreatedAt = now
	profile.UpdatedAt = now
	_, err = app.Engine.Insert(profile)
	if err != nil {
		app.Engine.Id(u.Id).Delete(u)
		errors.New("添加用户失败"+err.Error())
	}
	return nil;
}

func EditUser(id int, username, password string, department_id, employment_type, gender, level int, name, phone, phone1, position, email string, join_date,birth_date time.Time, ip_address string) error {
	now := time.Now()

	u := new(models.Users)
	u.Username = username
	u.UpdatedAt = now
	u.IpAddress = ip_address

	if len(password) >0 {
		u.Salt = utils.Salt()
		u.Password = utils.EncryptPassword(u.Salt, password)
	}

	_, err := app.Engine.Id(id).Update(u)
	if err != nil {
		return errors.New("编辑用户失败" + err.Error())
	}

	up := new(models.UserProfiles)
	up.DepartmentId = department_id
	up.Gender = gender
	up.Level = level
	up.EmploymentType = employment_type
	up.Name = name
	up.Phone = phone
	up.Phone1 = phone1
	up.Position = position
	up.JoinDate = join_date
	up.Email = email
	up.BirthDate = birth_date
	up.UpdatedAt = now

	_, err = app.Engine.Where("user_id=?", id).Update(up)
	if err != nil {
		return errors.New("编辑用户失败" + err.Error())
	}
	return nil
}

func EditSelf(id int, password string, phone, phone1, email string) error {
	now := time.Now()

	if len(password) >0 {
		u := new(models.Users)
		u.Salt = utils.Salt()
		u.UpdatedAt = now
		u.Password = utils.EncryptPassword(u.Salt, password)

		_, err := app.Engine.Id(id).Update(u)
		if err != nil {
			return errors.New("编辑用户失败" + err.Error())
		}
	}

	up := new(models.UserProfiles)
	up.Phone = phone
	up.Phone1 = phone1
	up.Email = email
	up.UpdatedAt = now

	_, err := app.Engine.Where("user_id=?", id).Update(up)
	if err != nil {
		return errors.New("编辑用户失败" + err.Error())
	}
	return nil
}