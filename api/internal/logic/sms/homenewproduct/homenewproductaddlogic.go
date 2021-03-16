package logic

import (
	"context"
	"go-zero-admin/rpc/sms/smsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeNewProductAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeNewProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeNewProductAddLogic {
	return HomeNewProductAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeNewProductAddLogic) HomeNewProductAdd(req types.AddHomeNewProductReq) (*types.AddHomeNewProductResp, error) {
	_, err := l.svcCtx.Sms.HomeNewProductAdd(l.ctx, &smsclient.HomeNewProductAddReq{
		ProductId:       req.ProductId,
		ProductName:     req.ProductName,
		RecommendStatus: req.RecommendStatus,
		Sort:            req.Sort,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddHomeNewProductResp{
		Code:    "000000",
		Message: "",
	}, nil
}
