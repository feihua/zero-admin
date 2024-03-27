package skustockservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// ReleaseSkuStockLockLogic
/*
Author: LiuFeiHua
Date: 2023/12/7 13:55
*/
type ReleaseSkuStockLockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReleaseSkuStockLockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReleaseSkuStockLockLogic {
	return &ReleaseSkuStockLockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ReleaseSkuStockLock 更新商品库存
func (l *ReleaseSkuStockLockLogic) ReleaseSkuStockLock(in *pmsclient.ReleaseSkuStockLockReq) (*pmsclient.ReleaseSkuStockLockResp, error) {
	for _, item := range in.Data {
		err := l.svcCtx.PmsSkuStockModel.ReleaseSkuStockLock(l.ctx, item.ProductSkuId, item.ProductQuantity)
		if err != nil {
			return nil, err
		}
	}

	return &pmsclient.ReleaseSkuStockLockResp{}, nil
}
