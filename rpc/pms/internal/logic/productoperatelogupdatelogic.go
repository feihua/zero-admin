package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductOperateLogUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductOperateLogUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductOperateLogUpdateLogic {
	return &ProductOperateLogUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductOperateLogUpdateLogic) ProductOperateLogUpdate(in *pms.ProductOperateLogUpdateReq) (*pms.ProductOperateLogUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &pms.ProductOperateLogUpdateResp{}, nil
}
