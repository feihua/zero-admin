package productskuservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QueryProductSkuDetailLogic 查询商品SKU详情
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type QueryProductSkuDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductSkuDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductSkuDetailLogic {
	return &QueryProductSkuDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductSkuDetail 查询商品SKU详情
func (l *QueryProductSkuDetailLogic) QueryProductSkuDetail(in *pmsclient.QueryProductSkuDetailReq) (*pmsclient.QueryProductSkuDetailResp, error) {
	item, err := query.PmsProductSku.WithContext(l.ctx).Where(query.PmsProductSku.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "商品SKU不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("商品SKU不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询商品SKU异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询商品SKU异常")
	}

	data := &pmsclient.QueryProductSkuDetailResp{
		Id:                 item.ID,                                          // 商品SpuId
		SpuId:              item.SpuID,                                       // 商品SpuId
		Name:               item.Name,                                        // SKU名称
		SkuCode:            item.SkuCode,                                     // SKU编码
		MainPic:            item.MainPic,                                     // 主图
		AlbumPics:          item.AlbumPics,                                   // 图片集
		Price:              float32(item.Price),                              // 价格
		PromotionPrice:     float32(item.PromotionPrice),                     // 单品促销价格
		PromotionStartTime: time_util.TimeToString(item.PromotionStartTime),  // 促销开始时间
		PromotionEndTime:   time_util.TimeToString(item.PromotionEndTime),    // 促销结束时间
		Stock:              item.Stock,                                       // 库存
		LowStock:           item.LowStock,                                    // 预警库存
		SpecData:           item.SpecData,                                    // 规格数据
		Weight:             float32(item.Weight),                             // 重量(kg)
		PublishStatus:      item.PublishStatus,                               // 上架状态：0-下架，1-上架
		VerifyStatus:       item.VerifyStatus,                                // 审核状态：0-未审核，1-审核通过，2-审核不通过
		Sort:               item.Sort,                                        // 排序
		Sales:              item.Sales,                                       // 销量
		CreateBy:           item.CreateBy,                                    // 创建人ID
		CreateTime:         time_util.TimeToStr(item.CreateTime),             // 创建时间
		UpdateBy:           pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
		UpdateTime:         time_util.TimeToString(item.UpdateTime),          // 更新时间
	}

	return data, nil
}
