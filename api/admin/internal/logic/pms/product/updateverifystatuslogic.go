package product

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateVerifyStatusLogic 批量修改审核状态
/*
Author: LiuFeiHua
Date: 2025/2/5 14:35
*/
type UpdateVerifyStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateVerifyStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVerifyStatusLogic {
	return &UpdateVerifyStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateVerifyStatus 批量修改审核状态
func (l *UpdateVerifyStatusLogic) UpdateVerifyStatus(req *types.UpdateProductStatusReq) (resp *types.UpdateProductStatusResp, err error) {
	_, err = l.svcCtx.ProductService.UpdateVerifyStatus(l.ctx, &pmsclient.UpdateProductStatusReq{
		Ids:    req.Ids,
		Status: req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "批量修改审核状态失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("批量修改审核状态失败")
	}

	return &types.UpdateProductStatusResp{
		Code:    "000000",
		Message: "批量修改审核状态成功",
	}, nil
}
