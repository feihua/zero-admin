package skustockservicelogic

import (
	"context"
	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// LockSkuStockLockLogic
/*
Author: LiuFeiHua
Date: 2023/12/13 13:59
*/
type LockSkuStockLockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLockSkuStockLockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LockSkuStockLockLogic {
	return &LockSkuStockLockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// LockSkuStockLock 下单的时候,锁定库存
func (l *LockSkuStockLockLogic) LockSkuStockLock(in *pmsclient.LockSkuStockLockReq) (*pmsclient.LockSkuStockLockResp, error) {
	for _, item := range in.Data {
		err := l.svcCtx.PmsSkuStockModel.LockSkuStockLock(l.ctx, item.ProductSkuId, item.ProductQuantity)
		if err != nil {
			return nil, err
		}
	}

	return &pmsclient.LockSkuStockLockResp{}, nil
}
