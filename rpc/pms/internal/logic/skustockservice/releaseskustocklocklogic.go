package skustockservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

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
	sql := "update pms_sku_stock set stock=stock+?,lock_stock=lock_stock-? where `id` = ? and lock_stock>=?"
	for _, item := range in.Data {
		db := l.svcCtx.DB
		err := db.Exec(sql, item.ProductSkuId, item.ProductQuantity).Error
		if err != nil {
			logc.Errorf(l.ctx, "更新商品库存失败,参数:%+v,异常:%s", in, err.Error())
			return nil, errors.New("更新商品库存失败")
		}
	}

	return &pmsclient.ReleaseSkuStockLockResp{}, nil
}
