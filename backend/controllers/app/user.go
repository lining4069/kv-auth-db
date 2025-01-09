package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lining4069/kv-auth-db/backend/app/common/request"
	"github.com/lining4069/kv-auth-db/backend/app/common/response"
	"github.com/lining4069/kv-auth-db/backend/app/services"
)

/*
Controllers层
校验入参,处理响应等工作
调用app/services下，对应的服务，完成对应逻辑
*/

// Register 用户注册
func Register(c *gin.Context) {
	var form request.Register // 用户注册入参 app/common/request/user.go
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	// 注册
	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}

// Info 获取当前用户信息
func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}
