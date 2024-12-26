package utils

import (
	"github.com/lining4069/ops-go/backend/global"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

/*
密码加密及验证密码的方法
*/

func BcryptMake(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		global.App.Log.Error("encode error", zap.Any("err", err))
	}
	return string(hash)
}

func BcryptMakeCheck(pwd []byte, hashedPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, pwd)
	if err != nil {
		return false
	}
	return true
}
