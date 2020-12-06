package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductOperateLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductOperateLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductOperateLogListLogic {
	return &ProductOperateLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductOperateLogListLogic) ProductOperateLogList(in *pms.ProductOperateLogListReq) (*pms.ProductOperateLogListResp, error) {
	// todo: add your logic here and delete this line

	return &pms.ProductOperateLogListResp{}, nil
}
