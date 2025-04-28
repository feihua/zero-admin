package product_category

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

// UpdateCategoryNavStatusLogic 更新导航状态
/*
Author: LiuFeiHua
Date: 2024/5/15 11:27
*/
type UpdateCategoryNavStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCategoryNavStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryNavStatusLogic {
	return &UpdateCategoryNavStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateCategoryNavStatus 更新导航状态
func (l *UpdateCategoryNavStatusLogic) UpdateCategoryNavStatus(req *types.UpdateProductCategoryStatusReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.ProductCategoryService.UpdateCategoryShowStatus(l.ctx, &pmsclient.UpdateProductCategoryStatusReq{
		Ids:    req.Ids,
		Status: req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新显示状态失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新显示状态失败")
	}

	return res.Success()
}
