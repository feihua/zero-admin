package cart

import (
	"context"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CartProductLogic
/*
Author: LiuFeiHua
Date: 2023/12/6 16:07
*/
type CartProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartProductLogic {
	return &CartProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CartProduct 获取购物车中某个商品的规格,用于重选规格
func (l *CartProductLogic) CartProduct(req *types.CartProductReq) (resp *types.CartProductResp, err error) {
	productResp, err := l.svcCtx.ProductSpuService.QueryProductSpuDetail(l.ctx, &pmsclient.QueryProductSpuDetailReq{
		Id: req.ProductId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "获取购物车中某个商品的规格,用于重选规格失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.CartProductResp{
		Code:    0,
		Message: "操作成功",
		Data: types.CartProductData{
			ProductAttributeList: buildProductAttributeListData(productResp),
			SkuStockList:         buildSkuStockListData(productResp),
		},
	}, nil
}

// 3.获取商品属性信息
func buildProductAttributeListData(resp *pmsclient.QueryProductSpuDetailResp) []types.CartItemProductAttributeList {
	var list []types.CartItemProductAttributeList
	for _, detail := range resp.ProductAttributeList {
		list = append(list, types.CartItemProductAttributeList{
			Id:           detail.Id,           // 主键id
			GroupId:      detail.GroupId,      // 属性分组ID
			Name:         detail.Name,         // 属性名称
			InputType:    detail.InputType,    // 输入类型：1-手动输入，2-单选，3-多选
			ValueType:    detail.ValueType,    // 值类型：1-文本，2-数字，3-日期
			InputList:    detail.InputList,    // 可选值列表，用逗号分隔
			Unit:         detail.Unit,         // 单位
			IsRequired:   detail.IsRequired,   // 是否必填
			IsSearchable: detail.IsSearchable, // 是否支持搜索
			IsShow:       detail.IsShow,       // 是否显示
			Sort:         detail.Sort,         // 排序
			CreateBy:     detail.CreateBy,     // 创建人ID
			CreateTime:   detail.CreateTime,   // 创建时间
			UpdateBy:     detail.UpdateBy,     // 更新人ID
			UpdateTime:   detail.UpdateTime,   // 更新时间
		})
	}

	return list
}

// 5.获取商品SKU库存信息
func buildSkuStockListData(resp *pmsclient.QueryProductSpuDetailResp) []types.CartItemSkuStockList {
	var list []types.CartItemSkuStockList
	for _, detail := range resp.SkuStockList {

		list = append(list, types.CartItemSkuStockList{
			Id:                 detail.Id,                      // 商品SpuId
			SpuId:              detail.SpuId,                   // 商品SpuId
			Name:               detail.Name,                    // SKU名称
			SkuCode:            detail.SkuCode,                 // SKU编码
			MainPic:            detail.MainPic,                 // 主图
			AlbumPics:          detail.AlbumPics,               // 图片集
			Price:              float64(detail.Price),          // 价格
			PromotionPrice:     float64(detail.PromotionPrice), // 单品促销价格
			PromotionStartTime: detail.PromotionStartTime,      // 促销开始时间
			PromotionEndTime:   detail.PromotionEndTime,        // 促销结束时间
			Stock:              detail.Stock,                   // 库存
			LowStock:           detail.LowStock,                // 预警库存
			SpecData:           detail.SpecData,                // 规格数据
			Weight:             float64(detail.Weight),         // 重量(kg)
			PublishStatus:      detail.PublishStatus,           // 上架状态：0-下架，1-上架
			VerifyStatus:       detail.VerifyStatus,            // 审核状态：0-未审核，1-审核通过，2-审核不通过
			Sort:               detail.Sort,                    // 排序
			Sales:              detail.Sales,                   // 销量
			CreateBy:           detail.CreateBy,                // 创建人ID
			CreateTime:         detail.CreateTime,              // 创建时间
			UpdateBy:           detail.UpdateBy,                // 更新人ID
			UpdateTime:         detail.UpdateTime,              // 更新时间
		})
	}

	return list
}
