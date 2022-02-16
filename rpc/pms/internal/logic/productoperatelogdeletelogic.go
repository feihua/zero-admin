package logic

import (
	"context"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductOperateLogDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductOperateLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductOperateLogDeleteLogic {
	return &ProductOperateLogDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductOperateLogDeleteLogic) ProductOperateLogDelete(in *pms.ProductOperateLogDeleteReq) (*pms.ProductOperateLogDeleteResp, error) {
	err := l.svcCtx.PmsProductOperateLogModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.ProductOperateLogDeleteResp{}, nil
}
