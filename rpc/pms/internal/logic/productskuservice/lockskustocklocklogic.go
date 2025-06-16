package productskuservicelogic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

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
func (l *LockSkuStockLockLogic) LockSkuStockLock(in *pmsclient.UpdateSkuStockReq) (*pmsclient.UpdateSkuStockLockResp, error) {
	sql := "update pms_product_sku set stock=stock-?,lock_stock=lock_stock+? where `id` = ? and stock>=?"
	for _, item := range in.Data {
		db := l.svcCtx.DB
		err := db.Exec(sql, item.Id, item.ProductQuantity).Error
		if err != nil {
			logc.Errorf(l.ctx, "下单的时候,锁定库存失败,参数:%+v,异常:%s", in, err.Error())
			return nil, errors.New("下单的时候,锁定库存失败")
		}
	}

	return &pmsclient.UpdateSkuStockLockResp{}, nil
}
