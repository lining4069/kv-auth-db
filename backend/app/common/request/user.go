package request

/*
用户相关的请求结构体，并实现Validator接口
*/

// Register 注册请求结构体
type Register struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binding:"required"`
}

// GetMessages 实现参数验证接口 (Validator interface)
func (register Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":     "用户名不能为空",
		"mobile.required":   "手机号不能为空",
		"mobile.mobile":     "手机号码格式不对",
		"password.required": "用户密码不能为空",
	}
}

// Login 登录请求结构体
type Login struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"mobile.required":   "手机号码不能为空",
		"mobile.mobile":     "手机号码格式不正确",
		"password.required": "用户密码不能为空",
	}
}
