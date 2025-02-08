package homenewproduct

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryHomeNewProductDetailLogic 查询新鲜好物表详情
/*
Author: 刘飞华
Date: 2025/02/07 10:11:43
*/
type QueryHomeNewProductDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryHomeNewProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHomeNewProductDetailLogic {
	return &QueryHomeNewProductDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryHomeNewProductDetail 查询新鲜好物表详情
func (l *QueryHomeNewProductDetailLogic) QueryHomeNewProductDetail(req *types.QueryHomeNewProductDetailReq) (resp *types.QueryHomeNewProductDetailResp, err error) {

	detail, err := l.svcCtx.HomeNewProductService.QueryHomeNewProductDetail(l.ctx, &smsclient.QueryHomeNewProductDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询新鲜好物表详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryHomeNewProductDetailData{
		Id:              detail.Id,              //
		ProductId:       detail.ProductId,       // 商品id
		ProductName:     detail.ProductName,     // 商品名称
		RecommendStatus: detail.RecommendStatus, // 推荐状态：0->不推荐;1->推荐
		Sort:            detail.Sort,            // 排序
	}
	return &types.QueryHomeNewProductDetailResp{
		Code:    "000000",
		Message: "查询新鲜好物表成功",
		Data:    data,
	}, nil
}
