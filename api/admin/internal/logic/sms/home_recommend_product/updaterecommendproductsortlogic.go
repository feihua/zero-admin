package home_recommend_product

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateRecommendProductSortLogic 人气推荐商品
/*
Author: LiuFeiHua
Date: 2024/5/14 9:41
*/
type UpdateRecommendProductSortLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRecommendProductSortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRecommendProductSortLogic {
	return &UpdateRecommendProductSortLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateRecommendProductSort 修改推荐排序
func (l *UpdateRecommendProductSortLogic) UpdateRecommendProductSort(req *types.UpdateRecommendProductSortReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.ProductSpuService.UpdateRecommendStatusSort(l.ctx, &pmsclient.UpdateProductSortReq{
		Id:   req.ProductId,
		Sort: req.Sort,
	})

	if err != nil {
		logc.Errorf(l.ctx, "修改推荐排序失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
