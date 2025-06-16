package product_sku

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

// UpdateProductSkuLogic 更新商品SKU
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type UpdateProductSkuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductSkuLogic {
	return &UpdateProductSkuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateProductSku 更新商品SKU
func (l *UpdateProductSkuLogic) UpdateProductSku(req *types.UpdateProductSkuReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}

	var data []*pmsclient.UpdateProductSkuData
	for _, s := range req.Data {
		data = append(data, &pmsclient.UpdateProductSkuData{
			Id:                 s.Id,                 // 商品SpuId
			SpuId:              s.SpuId,              // 商品SpuId
			Name:               s.Name,               // SKU名称
			SkuCode:            s.SkuCode,            // SKU编码
			MainPic:            s.MainPic,            // 主图
			AlbumPics:          s.AlbumPics,          // 图片集
			Price:              s.Price,              // 价格
			PromotionPrice:     s.PromotionPrice,     // 单品促销价格
			PromotionStartTime: s.PromotionStartTime, // 促销开始时间
			PromotionEndTime:   s.PromotionEndTime,   // 促销结束时间
			Stock:              s.Stock,              // 库存
			LowStock:           s.LowStock,           // 预警库存
			SpecData:           s.SpecData,           // 规格数据
			Weight:             s.Weight,             // 重量(kg)
			PublishStatus:      s.PublishStatus,      // 上架状态：0-下架，1-上架
			VerifyStatus:       s.VerifyStatus,       // 审核状态：0-未审核，1-审核通过，2-审核不通过
			Sort:               s.Sort,               // 排序
			UpdateBy:           userId,               // 更新人ID
		})
	}

	_, err = l.svcCtx.ProductSkuService.UpdateProductSku(l.ctx, &pmsclient.UpdateProductSkuReq{
		Data: data,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新商品SKU失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新商品SKU成功",
	}, nil
}
