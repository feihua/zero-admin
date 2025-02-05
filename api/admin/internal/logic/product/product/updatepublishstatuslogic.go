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

// UpdatePublishStatusLogic 批量上下架商品
/*
Author: LiuFeiHua
Date: 2025/2/5 14:33
*/
type UpdatePublishStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePublishStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePublishStatusLogic {
	return &UpdatePublishStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdatePublishStatus 批量上下架商品
func (l *UpdatePublishStatusLogic) UpdatePublishStatus(req *types.UpdateProductStatusReq) (resp *types.UpdateProductStatusResp, err error) {
	_, err = l.svcCtx.ProductService.UpdatePublishStatus(l.ctx, &pmsclient.UpdateProductStatusReq{
		Ids:    req.Ids,
		Status: req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "批量上下架商品失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("批量上下架商品失败")
	}

	return &types.UpdateProductStatusResp{
		Code:    "000000",
		Message: "批量上下架商品成功",
	}, nil
}
