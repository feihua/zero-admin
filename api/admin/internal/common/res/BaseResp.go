package res

import (
	"github.com/feihua/zero-admin/api/admin/internal/types"
)

// Success 返回成功
/*
Author: LiuFeiHua
Date: 2025/4/28 10:03
*/
func Success() (*types.BaseResp, error) {
	return &types.BaseResp{
		Code:    "000000",
		Message: "操作成功",
	}, nil
}

// Error 返回失败(默认)
func Error(msg string) *types.BaseResp {
	return &types.BaseResp{
		Code:    "111111",
		Message: msg,
	}
}
