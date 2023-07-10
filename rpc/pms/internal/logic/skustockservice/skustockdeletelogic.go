package skustockservicelogic

import (
	"context"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SkuStockDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSkuStockDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SkuStockDeleteLogic {
	return &SkuStockDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SkuStockDeleteLogic) SkuStockDelete(in *pmsclient.SkuStockDeleteReq) (*pmsclient.SkuStockDeleteResp, error) {
	err := l.svcCtx.PmsSkuStockModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &pmsclient.SkuStockDeleteResp{}, nil
}
