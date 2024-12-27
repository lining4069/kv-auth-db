package services

import (
	"errors"
	"github.com/lining4069/ops-go/backend/app/common/request"
	"github.com/lining4069/ops-go/backend/app/models"
	"github.com/lining4069/ops-go/backend/global"
	"github.com/lining4069/ops-go/backend/utils"
)

// userService 用户模型相关服务
type userService struct {
}

// UserService 提供给controllers层的调用入口
var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params request.Register) (err error, user models.User) {
	var result = global.App.DB.Where("mobile = ?", params.Mobile).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在")
		return
	}
	user = models.User{Name: params.Name, Mobile: params.Mobile, Password: utils.BcryptMake([]byte(params.Password))}
	err = global.App.DB.Create(&user).Error
	return
}

// Login 登录
func (userService *userService) Login(params request.Login) (err error, user *models.User) {
	err = global.App.DB.Where("mobile  = ?", params.Mobile).First(&user).Error
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("用户不存在或密码错误")
	}
	return
}
