package productskuservicelogic

import (
	"context"
	"errors"
	"time"

	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductSkuListLogic 查询商品SKU列表
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type QueryProductSkuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductSkuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductSkuListLogic {
	return &QueryProductSkuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductSkuList 查询商品SKU列表
func (l *QueryProductSkuListLogic) QueryProductSkuList(in *pmsclient.QueryProductSkuListReq) (*pmsclient.QueryProductSkuListResp, error) {
	productSku := query.PmsProductSku
	q := productSku.WithContext(l.ctx)
	if in.SpuId != 0 {
		q = q.Where(productSku.SpuID.Eq(in.SpuId))
	}
	if len(in.Name) > 0 {
		q = q.Where(productSku.Name.Like("%" + in.Name + "%"))
	}
	if len(in.SkuCode) > 0 {
		q = q.Where(productSku.SkuCode.Like("%" + in.SkuCode + "%"))
	}

	if len(in.PromotionStartTime) > 0 {
		startTime, _ := time.Parse("2006-01-02 15:04:05", in.PromotionStartTime)
		q = q.Where(productSku.PromotionStartTime.Gte(startTime))
	}
	if len(in.PromotionEndTime) > 0 {
		endTime, _ := time.Parse("2006-01-02 15:04:05", in.PromotionEndTime)
		q = q.Where(productSku.PromotionEndTime.Lte(endTime))
	}

	if in.PublishStatus != 2 {
		q = q.Where(productSku.PublishStatus.Eq(in.PublishStatus))
	}
	if in.VerifyStatus != 2 {
		q = q.Where(productSku.VerifyStatus.Eq(in.VerifyStatus))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询商品SKU列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询商品SKU列表失败")
	}

	var list []*pmsclient.ProductSkuListData

	for _, item := range result {
		list = append(list, &pmsclient.ProductSkuListData{
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

		})
	}

	return &pmsclient.QueryProductSkuListResp{
		Total: count,
		List:  list,
	}, nil
}
