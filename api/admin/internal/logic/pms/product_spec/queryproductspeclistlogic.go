package product_spec

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

// QueryProductSpecListLogic 查询商品规格列表
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type QueryProductSpecListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryProductSpecListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductSpecListLogic {
	return &QueryProductSpecListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryProductSpecList 查询商品规格列表
func (l *QueryProductSpecListLogic) QueryProductSpecList(req *types.QueryProductSpecListReq) (resp *types.QueryProductSpecListResp, err error) {
	result, err := l.svcCtx.ProductSpecService.QueryProductSpecList(l.ctx, &pmsclient.QueryProductSpecListReq{
		PageNum:    req.Current,
		PageSize:   req.PageSize,
		CategoryId: req.CategoryId, // 分类ID
		Name:       req.Name,       // 规格名称
		Status:     req.Status,     // 状态：0->禁用；1->启用
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字商品规格列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryProductSpecListData

	for _, detail := range result.List {
		list = append(list, &types.QueryProductSpecListData{
			Id:         detail.Id,         //
			CategoryId: detail.CategoryId, // 分类ID
			Name:       detail.Name,       // 规格名称
			Sort:       detail.Sort,       // 排序
			Status:     detail.Status,     // 状态：0->禁用；1->启用
			CreateBy:   detail.CreateBy,   // 创建人ID
			CreateTime: detail.CreateTime, // 创建时间
			UpdateBy:   detail.UpdateBy,   // 更新人ID
			UpdateTime: detail.UpdateTime, // 更新时间

		})
	}

	return &types.QueryProductSpecListResp{
		Code:     "000000",
		Message:  "查询商品规格列表成功",
		Data:     list,
		Current:  req.Current,
		PageSize: req.PageSize,
		Total:    result.Total,
		Success:  true,
	}, nil
}
