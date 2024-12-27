package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/lining4069/ops-go/backend/global"
	"time"
)

type jwtService struct {
}

var JwtService = new(jwtService)

// JwtUser 所有需要办法 token 的用户模型必须实现这个接口
type JwtUser interface {
	GetUid() string
}

type CustomClaims struct {
	jwt.StandardClaims
}

const (
	TokenType    = "Bearer" // token类型 JWT token前缀默认为Bearer
	AppGuardName = "ops-go" // token 发行者名称
)

type TokenOutPut struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// CreateToken 创建token jwtService方法
func (jwtService *jwtService) CreateToken(GuardName string, user JwtUser) (tokenData TokenOutPut, err error, token *jwt.Token) {
	// 使用jwt-go库生成token
	token = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + global.App.Config.Jwt.JwtTtl, // token 过期时间戳
				Id:        user.GetUid(),                                    // 用户唯一标识符
				Issuer:    GuardName,                                        // 签名发行者名称                                       //
				NotBefore: time.Now().Unix() - 60,                           //避免时钟偏差
			},
		})

	tokenStr, err := token.SignedString([]byte(global.App.Config.Jwt.Secret)) // 生成带签名的token

	tokenData = TokenOutPut{
		AccessToken: tokenStr,
		ExpiresIn:   int(global.App.Config.Jwt.JwtTtl),
		TokenType:   TokenType,
	}
	return
}
