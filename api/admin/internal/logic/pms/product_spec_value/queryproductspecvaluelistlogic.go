package product_spec_value

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

// QueryProductSpecValueListLogic 查询商品规格值列表
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type QueryProductSpecValueListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryProductSpecValueListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductSpecValueListLogic {
	return &QueryProductSpecValueListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryProductSpecValueList 查询商品规格值列表
func (l *QueryProductSpecValueListLogic) QueryProductSpecValueList(req *types.QueryProductSpecValueListReq) (resp *types.QueryProductSpecValueListResp, err error) {
	result, err := l.svcCtx.ProductSpecValueService.QueryProductSpecValueList(l.ctx, &pmsclient.QueryProductSpecValueListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
		SpecId:   req.SpecId, // 规格ID
		Status:   req.Status, // 状态：0->禁用；1->启用
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字商品规格值列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryProductSpecValueListData

	for _, detail := range result.List {
		list = append(list, &types.QueryProductSpecValueListData{
			Id:         detail.Id,         //
			SpecId:     detail.SpecId,     // 规格ID
			Value:      detail.Value,      // 规格值
			Sort:       detail.Sort,       // 排序
			Status:     detail.Status,     // 状态：0->禁用；1->启用
			CreateBy:   detail.CreateBy,   // 创建人ID
			CreateTime: detail.CreateTime, // 创建时间
			UpdateBy:   detail.UpdateBy,   // 更新人ID
			UpdateTime: detail.UpdateTime, // 更新时间

		})
	}

	return &types.QueryProductSpecValueListResp{
		Code:     "000000",
		Message:  "查询商品规格值列表成功",
		Data:     list,
		Current:  req.Current,
		PageSize: req.PageSize,
		Total:    result.Total,
		Success:  true,
	}, nil
}
