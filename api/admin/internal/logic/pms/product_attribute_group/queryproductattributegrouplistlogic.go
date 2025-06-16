package product_attribute_group

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductAttributeGroupListLogic 查询商品属性分组列表
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type QueryProductAttributeGroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryProductAttributeGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductAttributeGroupListLogic {
	return &QueryProductAttributeGroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryProductAttributeGroupList 查询商品属性分组列表
func (l *QueryProductAttributeGroupListLogic) QueryProductAttributeGroupList(req *types.QueryProductAttributeGroupListReq) (resp *types.QueryProductAttributeGroupListResp, err error) {
	result, err := l.svcCtx.ProductAttributeGroupService.QueryProductAttributeGroupList(l.ctx, &pmsclient.QueryProductAttributeGroupListReq{
		PageNum:    req.Current,
		PageSize:   req.PageSize,
		CategoryId: req.CategoryId, // 分类ID
		Name:       req.Name,       // 分组名称
		Status:     req.Status,     // 状态：0->禁用；1->启用
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字商品属性分组列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryProductAttributeGroupListData

	for _, detail := range result.List {
		list = append(list, &types.QueryProductAttributeGroupListData{
			Id:         detail.Id,         // 主键id
			CategoryId: detail.CategoryId, // 分类ID
			Name:       detail.Name,       // 分组名称
			Sort:       detail.Sort,       // 排序
			Status:     detail.Status,     // 状态：0->禁用；1->启用
			CreateBy:   detail.CreateBy,   // 创建人ID
			CreateTime: detail.CreateTime, // 创建时间
			UpdateBy:   detail.UpdateBy,   // 更新人ID
			UpdateTime: detail.UpdateTime, // 更新时间

		})
	}

	return &types.QueryProductAttributeGroupListResp{
		Code:     "000000",
		Message:  "查询商品属性分组列表成功",
		Data:     list,
		Current:  req.Current,
		PageSize: req.PageSize,
		Total:    result.Total,
		Success:  true,
	}, nil
}
