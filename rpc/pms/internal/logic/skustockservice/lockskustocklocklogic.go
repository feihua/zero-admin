package skustockservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// LockSkuStockLockLogic 库存
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

	sql := "update pms_sku_stock set stock=stock-?,lock_stock=lock_stock+? where `id` = ? and stock>=?"
	for _, item := range in.Data {
		db := l.svcCtx.DB
		err := db.Exec(sql, item.ProductSkuId, item.ProductQuantity).Error
		if err != nil {
			return nil, err
		}
	}

	return &pmsclient.LockSkuStockLockResp{}, nil
}
