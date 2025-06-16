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

// UpdateProductAttributeGroupLogic 更新商品属性分组
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type UpdateProductAttributeGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductAttributeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductAttributeGroupLogic {
	return &UpdateProductAttributeGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateProductAttributeGroup 更新商品属性分组
func (l *UpdateProductAttributeGroupLogic) UpdateProductAttributeGroup(req *types.UpdateProductAttributeGroupReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.ProductAttributeGroupService.UpdateProductAttributeGroup(l.ctx, &pmsclient.UpdateProductAttributeGroupReq{
		Id:         req.Id,         // 主键id
		CategoryId: req.CategoryId, // 分类ID
		Name:       req.Name,       // 分组名称
		Sort:       req.Sort,       // 排序
		Status:     req.Status,     // 状态：0->禁用；1->启用
		UpdateBy:   userId,         // 更新人ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新商品属性分组失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新商品属性分组成功",
	}, nil
}
