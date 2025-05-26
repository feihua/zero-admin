package product_category

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateProductCategoryStatusLogic 更新产品分类状态状态
/*
Author: LiuFeiHua
Date: 2025/05/26 13:56:03
*/
type UpdateProductCategoryStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductCategoryStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductCategoryStatusLogic {
	return &UpdateProductCategoryStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateProductCategoryStatus 更新产品分类状态
func (l *UpdateProductCategoryStatusLogic) UpdateProductCategoryStatus(req *types.UpdateProductCategoryStatusReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.ProductCategoryService.UpdateProductCategoryStatus(l.ctx, &pmsclient.UpdateProductCategoryStatusReq{
		Ids:      req.Ids,    //
		Status:   req.Status, //
		UpdateBy: userId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新产品分类状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新产品分类状态成功",
	}, nil
}
