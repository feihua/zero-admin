package logic

import "go-zero-admin/shared"

var (
	errorDuplicateUsername  = shared.NewDefaultError("用户名已经注册")
	errorDuplicateMobile    = shared.NewDefaultError("手机号已经被占用")
	errorUsernameUnRegister = shared.NewDefaultError("用户未注册")
	errorRoleUnRegister     = shared.NewDefaultError("角色不存在")
	errorUserRegisterFail   = shared.NewDefaultError("用户注册失败")
	errorIncorrectPassword  = shared.NewDefaultError("用户密码错误")
	errorUserNotFound       = shared.NewDefaultError("用户不存在")

)
