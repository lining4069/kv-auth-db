package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lining4069/kv-auth-db/backend/app/common/request"
	"github.com/lining4069/kv-auth-db/backend/app/common/response"
	"github.com/lining4069/kv-auth-db/backend/app/services"
)

func BitcaskPut(c *gin.Context) {
	var form request.BitcaskPutRequest
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err := services.BitcaskSBService.Put(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, nil)
	}
}

func BitcaskGet(c *gin.Context) {
	key := c.DefaultQuery("key", "")
	if err, value := services.BitcaskSBService.Get(key); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, value)
	}

}
