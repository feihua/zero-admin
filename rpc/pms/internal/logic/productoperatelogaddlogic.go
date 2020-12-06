package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductOperateLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductOperateLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductOperateLogAddLogic {
	return &ProductOperateLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductOperateLogAddLogic) ProductOperateLogAdd(in *pms.ProductOperateLogAddReq) (*pms.ProductOperateLogAddResp, error) {
	// todo: add your logic here and delete this line

	return &pms.ProductOperateLogAddResp{}, nil
}
