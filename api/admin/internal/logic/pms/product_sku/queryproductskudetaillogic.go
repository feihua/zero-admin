package product_sku

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductSkuDetailLogic 查询商品SKU详情
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type QueryProductSkuDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryProductSkuDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductSkuDetailLogic {
	return &QueryProductSkuDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryProductSkuDetail 查询商品SKU详情
func (l *QueryProductSkuDetailLogic) QueryProductSkuDetail(req *types.QueryProductSkuDetailReq) (resp *types.QueryProductSkuDetailResp, err error) {

	detail, err := l.svcCtx.ProductSkuService.QueryProductSkuDetail(l.ctx, &pmsclient.QueryProductSkuDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询商品SKU详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryProductSkuDetailData{
		Id:                 detail.Id,                 // 商品SpuId
		SpuId:              detail.SpuId,              // 商品SpuId
		Name:               detail.Name,               // SKU名称
		SkuCode:            detail.SkuCode,            // SKU编码
		MainPic:            detail.MainPic,            // 主图
		AlbumPics:          detail.AlbumPics,          // 图片集
		Price:              detail.Price,              // 价格
		PromotionPrice:     detail.PromotionPrice,     // 单品促销价格
		PromotionStartTime: detail.PromotionStartTime, // 促销开始时间
		PromotionEndTime:   detail.PromotionEndTime,   // 促销结束时间
		Stock:              detail.Stock,              // 库存
		LowStock:           detail.LowStock,           // 预警库存
		SpecData:           detail.SpecData,           // 规格数据
		Weight:             detail.Weight,             // 重量(kg)
		PublishStatus:      detail.PublishStatus,      // 上架状态：0-下架，1-上架
		VerifyStatus:       detail.VerifyStatus,       // 审核状态：0-未审核，1-审核通过，2-审核不通过
		Sort:               detail.Sort,               // 排序
		Sales:              detail.Sales,              // 销量
		CreateBy:           detail.CreateBy,           // 创建人ID
		CreateTime:         detail.CreateTime,         // 创建时间
		UpdateBy:           detail.UpdateBy,           // 更新人ID
		UpdateTime:         detail.UpdateTime,         // 更新时间

	}
	return &types.QueryProductSkuDetailResp{
		Code:    "000000",
		Message: "查询商品SKU成功",
		Data:    data,
	}, nil
}
