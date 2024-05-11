package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

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
		logc.Errorf(l.ctx, "更新人气推荐商品信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新人气推荐商品失败")
	}

	return &types.UpdateHomeRecommendProductResp{
		Code:    "000000",
		Message: "更新人气推荐商品成功",
	}, nil
}
