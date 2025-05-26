package product_category

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

// QueryProductCategoryListLogic 查询产品分类列表
/*
Author: LiuFeiHua
Date: 2025/05/26 13:56:03
*/
type QueryProductCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryProductCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductCategoryListLogic {
	return &QueryProductCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryProductCategoryList 查询产品分类列表
func (l *QueryProductCategoryListLogic) QueryProductCategoryList(req *types.QueryProductCategoryListReq) (resp *types.QueryProductCategoryListResp, err error) {
	result, err := l.svcCtx.ProductCategoryService.QueryProductCategoryList(l.ctx, &pmsclient.QueryProductCategoryListReq{
		PageNum:   req.Current,
		PageSize:  req.PageSize,
		ParentId:  req.ParentId,  // 上级分类的编号：0表示一级分类
		Name:      req.Name,      // 商品分类名称
		NavStatus: req.NavStatus, // 是否显示在导航栏：0->不显示；1->显示
		Keywords:  req.Keywords,  // 关键字
		IsEnabled: req.IsEnabled, // 是否启用
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字产品分类列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryProductCategoryListData

	for _, detail := range result.List {
		list = append(list, &types.QueryProductCategoryListData{
			Id:           detail.Id,           //
			ParentId:     detail.ParentId,     // 上级分类的编号：0表示一级分类
			Name:         detail.Name,         // 商品分类名称
			Level:        detail.Level,        // 分类级别：0->1级；1->2级
			ProductCount: detail.ProductCount, // 商品数量
			ProductUnit:  detail.ProductUnit,  // 商品单位
			NavStatus:    detail.NavStatus,    // 是否显示在导航栏：0->不显示；1->显示
			Sort:         detail.Sort,         // 排序
			Icon:         detail.Icon,         // 图标
			Keywords:     detail.Keywords,     // 关键字
			Description:  detail.Description,  // 描述
			IsEnabled:    detail.IsEnabled,    // 是否启用
			CreateBy:     detail.CreateBy,     // 创建人ID
			CreateTime:   detail.CreateTime,   // 创建时间
			UpdateBy:     detail.UpdateBy,     // 更新人ID
			UpdateTime:   detail.UpdateTime,   // 更新时间

		})
	}

	return &types.QueryProductCategoryListResp{
		Code:     "000000",
		Message:  "查询产品分类列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
