package category

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCategoryShowStatusLogic 更新显示状态
/*
Author: LiuFeiHua
Date: 2024/5/15 11:22
*/
type UpdateCategoryShowStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCategoryShowStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryShowStatusLogic {
	return &UpdateCategoryShowStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateCategoryShowStatus 更新显示状态
func (l *UpdateCategoryShowStatusLogic) UpdateCategoryShowStatus(req *types.UpdateProductCategoryStatusReq) (resp *types.UpdateProductCategoryResp, err error) {
	_, err = l.svcCtx.ProductCategoryService.UpdateCategoryNavStatus(l.ctx, &pmsclient.UpdateProductCategoryStatusReq{
		Ids:    req.Ids,
		Status: req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新显示状态失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新显示状态失败")
	}

	return &types.UpdateProductCategoryResp{
		Code:    "000000",
		Message: "更新显示状态成功",
	}, nil
}
