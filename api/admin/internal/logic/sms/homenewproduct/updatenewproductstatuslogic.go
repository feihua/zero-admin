package homenewproduct

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateNewProductStatusLogic
/*
Author: LiuFeiHua
Date: 2024/5/14 9:18
*/
type UpdateNewProductStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateNewProductStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNewProductStatusLogic {
	return &UpdateNewProductStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateNewProductStatus 批量修改推荐状态
func (l *UpdateNewProductStatusLogic) UpdateNewProductStatus(req *types.UpdateNewProductStatusReq) (resp *types.UpdateNewProductStatusResp, err error) {
	_, err = l.svcCtx.HomeNewProductService.UpdateNewProductStatus(l.ctx, &smsclient.UpdateNewProductStatusReq{
		Ids:             req.Ids,
		RecommendStatus: req.RecommendStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "批量修改推荐状态失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("批量修改推荐状态失败")
	}

	return &types.UpdateNewProductStatusResp{
		Code:    "000000",
		Message: "批量修改推荐状态成功",
	}, nil

}
