package product_attribute

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

// UpdateProductAttributeStatusLogic 更新商品属性状态状态
/*
Author: LiuFeiHua
Date: 2025/06/17 16:14:47
*/
type UpdateProductAttributeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductAttributeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductAttributeStatusLogic {
	return &UpdateProductAttributeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateProductAttributeStatus 更新商品属性状态
func (l *UpdateProductAttributeStatusLogic) UpdateProductAttributeStatus(req *types.UpdateProductAttributeStatusReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.ProductAttributeService.UpdateProductAttributeStatus(l.ctx, &pmsclient.UpdateProductAttributeStatusReq{
		Ids:      req.Ids,    // 主键id
		Status:   req.Status, // 状态：0->禁用；1->启用
		UpdateBy: userId,     // 更新人ID

	})

	if err != nil {
		logc.Errorf(l.ctx, "更新商品属性状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新商品属性状态成功",
	}, nil
}
