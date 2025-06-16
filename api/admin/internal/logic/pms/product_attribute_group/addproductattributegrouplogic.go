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

// AddProductAttributeGroupLogic 添加商品属性分组
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type AddProductAttributeGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProductAttributeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductAttributeGroupLogic {
	return &AddProductAttributeGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddProductAttributeGroup 添加商品属性分组
func (l *AddProductAttributeGroupLogic) AddProductAttributeGroup(req *types.AddProductAttributeGroupReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.ProductAttributeGroupService.AddProductAttributeGroup(l.ctx, &pmsclient.AddProductAttributeGroupReq{
		CategoryId: req.CategoryId, // 分类ID
		Name:       req.Name,       // 分组名称
		Sort:       req.Sort,       // 排序
		Status:     req.Status,     // 状态：0->禁用；1->启用
		CreateBy:   userId,         // 创建人ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加商品属性分组失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "添加商品属性分组成功",
	}, nil
}
