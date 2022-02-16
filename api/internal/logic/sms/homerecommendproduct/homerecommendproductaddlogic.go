package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sms/smsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeRecommendProductAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeRecommendProductAddLogic {
	return HomeRecommendProductAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeRecommendProductAddLogic) HomeRecommendProductAdd(req types.AddHomeRecommendProductReq) (*types.AddHomeRecommendProductResp, error) {
	_, err := l.svcCtx.Sms.HomeRecommendProductAdd(l.ctx, &smsclient.HomeRecommendProductAddReq{
		ProductId:       req.ProductId,
		ProductName:     req.ProductName,
		RecommendStatus: req.RecommendStatus,
		Sort:            req.Sort,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加人气推荐商品信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加人气推荐商品失败")
	}

	return &types.AddHomeRecommendProductResp{
		Code:    "000000",
		Message: "添加人气推荐商品成功",
	}, nil
}
