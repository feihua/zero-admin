package product_sku

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddProductSkuLogic 添加商品SKU
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type AddProductSkuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProductSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductSkuLogic {
	return &AddProductSkuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddProductSku 添加商品SKU
func (l *AddProductSkuLogic) AddProductSku(req *types.AddProductSkuReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.ProductSkuService.AddProductSku(l.ctx, &pmsclient.AddProductSkuReq{
		Name:               req.Name,               // SKU名称
		MainPic:            req.MainPic,            // 主图
		AlbumPics:          req.AlbumPics,          // 图片集
		Price:              req.Price,              // 价格
		PromotionPrice:     req.PromotionPrice,     // 单品促销价格
		PromotionStartTime: req.PromotionStartTime, // 促销开始时间
		PromotionEndTime:   req.PromotionEndTime,   // 促销结束时间
		Stock:              req.Stock,              // 库存
		LowStock:           req.LowStock,           // 预警库存
		SpecData:           req.SpecData,           // 规格数据
		Weight:             req.Weight,             // 重量(kg)
		PublishStatus:      req.PublishStatus,      // 上架状态：0-下架，1-上架
		VerifyStatus:       req.VerifyStatus,       // 审核状态：0-未审核，1-审核通过，2-审核不通过
		Sort:               req.Sort,               // 排序
		CreateBy:           userId,                 // 创建人ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加商品SKU失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "添加商品SKU成功",
	}, nil
}
