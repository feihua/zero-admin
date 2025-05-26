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

// UpdateProductCategoryLogic 更新产品分类
/*
Author: LiuFeiHua
Date: 2025/05/26 13:56:03
*/
type UpdateProductCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductCategoryLogic {
	return &UpdateProductCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateProductCategory 更新产品分类
func (l *UpdateProductCategoryLogic) UpdateProductCategory(req *types.UpdateProductCategoryReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.ProductCategoryService.UpdateProductCategory(l.ctx, &pmsclient.UpdateProductCategoryReq{
		Id:          req.Id,          //
		ParentId:    req.ParentId,    // 上级分类的编号：0表示一级分类
		Name:        req.Name,        // 商品分类名称
		Level:       req.Level,       // 分类级别：0->1级；1->2级
		ProductUnit: req.ProductUnit, // 商品单位
		NavStatus:   req.NavStatus,   // 是否显示在导航栏：0->不显示；1->显示
		Sort:        req.Sort,        // 排序
		Icon:        req.Icon,        // 图标
		Keywords:    req.Keywords,    // 关键字
		Description: req.Description, // 描述
		IsEnabled:   req.IsEnabled,   // 是否启用
		UpdateBy:    userId,          // 更新人ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新产品分类失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新产品分类成功",
	}, nil
}
