package common

// 设置返回给前端的msg
const (
	RegisterSuccess = "注册成功"
	RegisterFail    = "注册失败"

	UserNameExist   = "用户名已存在"
	UserNotExist    = "用户不存在"
	LoginSuccess    = "登录成功"
	LoginFail       = "登录失败"
	DataFormatError = "传递数据格式有误"

	UserNameOrPasswordError = "用户名或密码错误"

	GetUserListFail = "无法获取用户列表"

	GetUserInfoFail = "无法获取用户信息"

	UpdateUserSuccess = "更新用户信息成功"
	UpdateUserFail    = "更新用户信息失败"

	NotLogin = "请登录"

	TokenExpired = "token已失效"

	ToeknError = "token错误"
)
