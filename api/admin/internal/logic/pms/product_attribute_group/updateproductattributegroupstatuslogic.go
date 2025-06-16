package product_attribute_group

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

// UpdateProductAttributeGroupStatusLogic 更新商品属性分组状态状态
/*
Author: LiuFeiHua
Date: 2025/06/17 16:14:47
*/
type UpdateProductAttributeGroupStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductAttributeGroupStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductAttributeGroupStatusLogic {
	return &UpdateProductAttributeGroupStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateProductAttributeGroupStatus 更新商品属性分组状态
func (l *UpdateProductAttributeGroupStatusLogic) UpdateProductAttributeGroupStatus(req *types.UpdateProductAttributeGroupStatusReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.ProductAttributeGroupService.UpdateProductAttributeGroupStatus(l.ctx, &pmsclient.UpdateProductAttributeGroupStatusReq{
		Ids:      req.Ids,    // 主键id
		Status:   req.Status, // 状态：0->禁用；1->启用
		UpdateBy: userId,     // 更新人ID

	})

	if err != nil {
		logc.Errorf(l.ctx, "更新商品属性分组状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新商品属性分组状态成功",
	}, nil
}
