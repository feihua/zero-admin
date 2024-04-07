package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeRecommendProductUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeRecommendProductUpdateLogic {
	return HomeRecommendProductUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeRecommendProductUpdateLogic) HomeRecommendProductUpdate(req types.UpdateHomeRecommendProductReq) (*types.UpdateHomeRecommendProductResp, error) {
	_, err := l.svcCtx.HomeRecommendProductService.HomeRecommendProductUpdate(l.ctx, &smsclient.HomeRecommendProductUpdateReq{
		Id:              req.Id,
		ProductId:       req.ProductId,
		ProductName:     req.ProductName,
		RecommendStatus: req.RecommendStatus,
		Sort:            req.Sort,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新人气推荐商品信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新人气推荐商品失败")
	}

	return &types.UpdateHomeRecommendProductResp{
		Code:    "000000",
		Message: "更新人气推荐商品成功",
	}, nil
}
