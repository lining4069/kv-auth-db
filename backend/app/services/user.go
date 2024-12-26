package services

import (
	"errors"
	"github.com/lining4069/ops-go/backend/app/common/request"
	"github.com/lining4069/ops-go/backend/app/models"
	"github.com/lining4069/ops-go/backend/global"
	"github.com/lining4069/ops-go/backend/utils"
)

type userService struct {
}

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
