package common

import (
	"context"
	"encoding/json"

	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logc"
)

func GetUserId(ctx context.Context) (int64, error) {

	rawUserId := ctx.Value("userId")

	if rawUserId == nil {
		logc.Errorf(ctx, "获取当前操作者id失败，上下文中未找到键 'userId'")
		return 0, errorx.NewDefaultError("获取当前操作者id失败")
	}

	jsonNumber, ok := rawUserId.(json.Number)
	if !ok {
		logc.Errorf(ctx, "获取当前操作者id失败，类型断言失败。期望 json.Number，实际类型：%T", rawUserId)
		return 0, errorx.NewDefaultError("获取当前操作者id失败")
	}

	userId, err := jsonNumber.Int64()
	if err != nil {
		logc.Errorf(ctx, "获取当前操作者id失败，转换为 int64 时出错：%s", err.Error())
		return 0, errorx.NewDefaultError("获取当前操作者id失败")
	}

	logc.Infof(ctx, "当前操作会员的id: %d", userId)
	return userId, nil
}

func GetUserName(ctx context.Context) (string, error) {

	rawUserName := ctx.Value("userName")

	if rawUserName == nil {
		logc.Errorf(ctx, "获取当前操作者name失败，上下文中未找到键 'userName', 确认api中是否已经配置了jwt:Auth")
		return "", errorx.NewDefaultError("获取当前操作者name失败")
	}

	userName, ok := rawUserName.(string)
	if !ok {
		logc.Errorf(ctx, "获取当前操作者name失败，类型断言失败。期望 string，实际类型：%T", rawUserName)
		return "", errorx.NewDefaultError("获取当前操作者name失败")
	}

	logc.Infof(ctx, "当前操作会员的name: %s", userName)
	return userName, nil
}
