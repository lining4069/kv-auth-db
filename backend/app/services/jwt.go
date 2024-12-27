package services

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/lining4069/ops-go/backend/global"
	"github.com/lining4069/ops-go/backend/utils"
	"strconv"
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

// GetBlackListKey 获取黑名单缓存 key
func (jwtService *jwtService) GetBlackListKey(tokenStr string) string {
	return "jwt_black_list" + utils.Md5([]byte(tokenStr))
}

// JoinBlackList Token加入黑名单
func (jwtService *jwtService) JoinBlackList(token *jwt.Token) (err error) {
	nowUnix := time.Now().Unix()
	timer := time.Duration(token.Claims.(*CustomClaims).ExpiresAt-nowUnix) * time.Second
	// 将Token剩余时间设置为缓存有效期，并将当前时间作为缓存value值
	err = global.App.Redis.SetNX(context.Background(), jwtService.GetBlackListKey(token.Raw), nowUnix, timer).Err()
	return
}

// IsInBlacklist Token是否在黑名单中
func (jwtService *jwtService) IsInBlacklist(tokenStr string) bool {
	joinUnixStr, err := global.App.Redis.Get(context.Background(), jwtService.GetBlackListKey(tokenStr)).Result()
	joinUnix, err := strconv.ParseInt(joinUnixStr, 10, 64) // base:字符串表示的数字的基数 bitSize:结果的位大小 通常64位
	if joinUnixStr == "" || err != nil {
		return false
	}

	// JoinBlacklistGracePeriod 为黑名单宽容时间，避免并发请求失败
	if time.Now().Unix()-joinUnix < global.App.Config.Jwt.JwtBlacklistGracePeriod {
		return false
	}
	return true

}
