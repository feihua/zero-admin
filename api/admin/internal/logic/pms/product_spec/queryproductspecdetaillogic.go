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

// QueryProductSpecDetailLogic 查询商品规格详情
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type QueryProductSpecDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryProductSpecDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductSpecDetailLogic {
	return &QueryProductSpecDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryProductSpecDetail 查询商品规格详情
func (l *QueryProductSpecDetailLogic) QueryProductSpecDetail(req *types.QueryProductSpecDetailReq) (resp *types.QueryProductSpecDetailResp, err error) {

	detail, err := l.svcCtx.ProductSpecService.QueryProductSpecDetail(l.ctx, &pmsclient.QueryProductSpecDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询商品规格详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryProductSpecDetailData{
		Id:         detail.Id,         //
		CategoryId: detail.CategoryId, // 分类ID
		Name:       detail.Name,       // 规格名称
		Sort:       detail.Sort,       // 排序
		Status:     detail.Status,     // 状态：0->禁用；1->启用
		CreateBy:   detail.CreateBy,   // 创建人ID
		CreateTime: detail.CreateTime, // 创建时间
		UpdateBy:   detail.UpdateBy,   // 更新人ID
		UpdateTime: detail.UpdateTime, // 更新时间

	}
	return &types.QueryProductSpecDetailResp{
		Code:    "000000",
		Message: "查询商品规格成功",
		Data:    data,
	}, nil
}
