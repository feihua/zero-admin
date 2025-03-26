package common

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/zeromicro/go-zero/core/logc"
)

func GetMemberId(ctx context.Context) (int64, error) {

	rawMemberId := ctx.Value("memberId")

	if rawMemberId == nil {
		logc.Errorf(ctx, "获取当前操作的会员id失败，上下文中未找到键 'memberId'")
		return 0, errorx.NewDefaultError("获取当前操作的会员id失败")
	}

	jsonNumber, ok := rawMemberId.(json.Number)
	if !ok {
		logc.Errorf(ctx, "获取当前操作的会员id失败，类型断言失败。期望 json.Number，实际类型：%T", rawMemberId)
		return 0, errorx.NewDefaultError("获取当前操作的会员id失败")
	}

	memberId, err := jsonNumber.Int64()
	if err != nil {
		logc.Errorf(ctx, "获取当前操作的会员id失败，转换为 int64 时出错：%s", err.Error())
		return 0, errorx.NewDefaultError("获取当前操作的会员id失败")
	}

	logc.Infof(ctx, "当前操作会员的id: %d", memberId)
	return memberId, nil
}
