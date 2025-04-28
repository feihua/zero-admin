package product

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateRecommendStatusLogic 批量推荐商品
/*
Author: LiuFeiHua
Date: 2025/2/5 14:34
*/
type UpdateRecommendStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRecommendStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRecommendStatusLogic {
	return &UpdateRecommendStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateRecommendStatus 批量推荐商品
func (l *UpdateRecommendStatusLogic) UpdateRecommendStatus(req *types.UpdateProductStatusReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.ProductService.UpdateRecommendStatus(l.ctx, &pmsclient.UpdateProductStatusReq{
		Ids:    req.Ids,
		Status: req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "批量推荐商品失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("批量推荐商品失败")
	}

	return res.Success()
}
