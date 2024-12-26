package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lining4069/ops-go/backend/app/common/request"
	"github.com/lining4069/ops-go/backend/app/common/response"
	"github.com/lining4069/ops-go/backend/app/services"
)

/*
MVC C层
校验入参等工作
调用app/services下，对应的服务，完成对应逻辑
*/

// Register 用户注册
func Register(c *gin.Context) {
	var form request.Register // 用户注册入参 app/common/request/user.go
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}
