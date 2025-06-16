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

// QueryProductSpecValueDetailLogic 查询商品规格值详情
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type QueryProductSpecValueDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryProductSpecValueDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductSpecValueDetailLogic {
	return &QueryProductSpecValueDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryProductSpecValueDetail 查询商品规格值详情
func (l *QueryProductSpecValueDetailLogic) QueryProductSpecValueDetail(req *types.QueryProductSpecValueDetailReq) (resp *types.QueryProductSpecValueDetailResp, err error) {

	detail, err := l.svcCtx.ProductSpecValueService.QueryProductSpecValueDetail(l.ctx, &pmsclient.QueryProductSpecValueDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询商品规格值详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryProductSpecValueDetailData{
		Id:         detail.Id,         //
		SpecId:     detail.SpecId,     // 规格ID
		Value:      detail.Value,      // 规格值
		Sort:       detail.Sort,       // 排序
		Status:     detail.Status,     // 状态：0->禁用；1->启用
		CreateBy:   detail.CreateBy,   // 创建人ID
		CreateTime: detail.CreateTime, // 创建时间
		UpdateBy:   detail.UpdateBy,   // 更新人ID
		UpdateTime: detail.UpdateTime, // 更新时间

	}
	return &types.QueryProductSpecValueDetailResp{
		Code:    "000000",
		Message: "查询商品规格值成功",
		Data:    data,
	}, nil
}
