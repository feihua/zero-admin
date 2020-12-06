package logic

import (
	"context"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeNewProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeNewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeNewProductListLogic {
	return &HomeNewProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeNewProductListLogic) HomeNewProductList(in *sms.HomeNewProductListReq) (*sms.HomeNewProductListResp, error) {
	// todo: add your logic here and delete this line

	return &sms.HomeNewProductListResp{}, nil
}
