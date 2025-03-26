package skustockservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddSkuStockLogic 添加sku的库存
/*
Author: LiuFeiHua
Date: 2024/6/12 17:10
*/
type AddSkuStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSkuStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSkuStockLogic {
	return &AddSkuStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddSkuStock 添加sku的库存
func (l *AddSkuStockLogic) AddSkuStock(in *pmsclient.AddSkuStockReq) (*pmsclient.AddSkuStockResp, error) {
	err := query.PmsSkuStock.WithContext(l.ctx).Create(&model.PmsSkuStock{
		ProductID:      in.ProductId,      // 商品id
		SkuCode:        in.SkuCode,        // sku编码
		Price:          in.Price,          // 价格
		Stock:          in.Stock,          // 库存
		LowStock:       in.LowStock,       // 预警库存
		Pic:            in.Pic,            // 展示图片
		Sale:           in.Sale,           // 销量
		PromotionPrice: in.PromotionPrice, // 单品促销价格
		LockStock:      in.LockStock,      // 锁定库存
		SpData:         in.SpData,         // 商品销售属性，json格式
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加sku的库存失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加sku的库存失败")
	}

	return &pmsclient.AddSkuStockResp{}, nil
}
