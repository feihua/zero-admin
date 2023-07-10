package productoperatelogservicelogic

import (
	"context"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/rpc/pms/internal/svc"

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

func (l *ProductOperateLogDeleteLogic) ProductOperateLogDelete(in *pmsclient.ProductOperateLogDeleteReq) (*pmsclient.ProductOperateLogDeleteResp, error) {
	err := l.svcCtx.PmsProductOperateLogModel.Delete(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductOperateLogDeleteResp{}, nil
}
