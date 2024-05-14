package flashpromotionproductrelation

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// FlashPromotionProductRelationListLogic 限时购和商品关系
/*
Author: LiuFeiHua
Date: 2024/5/14 15:31
*/
type FlashPromotionProductRelationListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionProductRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionProductRelationListLogic {
	return &FlashPromotionProductRelationListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// FlashPromotionProductRelationList 分页查询不同场次关联及商品信息
func (l *FlashPromotionProductRelationListLogic) FlashPromotionProductRelationList(req *types.ListFlashPromotionProductRelationReq) (resp *types.ListFlashPromotionProductRelationResp, err error) {
	result, err := l.svcCtx.FlashPromotionProductRelationService.FlashPromotionProductRelationList(l.ctx, &smsclient.FlashPromotionProductRelationListReq{
		Current:                 req.Current,
		PageSize:                req.PageSize,
		FlashPromotionId:        req.FlashPromotionProductRelationID,
		FlashPromotionSessionId: req.FlashPromotionProductRelationSessionID,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,分页查询不同场次关联及商品信息列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("分页查询不同场次关联及商品信息失败")
	}

	var ids []int64
	for _, item := range result.List {
		ids = append(ids, item.ProductId)
	}
	var list []*types.ListFlashPromotionProductRelationData

	products, err := l.svcCtx.ProductService.ProductListByIds(l.ctx, &pmsclient.ProductByIdsReq{Ids: ids})
	for _, item := range products.List {
		var r *smsclient.FlashPromotionProductRelationListData
		for _, data := range result.List {
			if data.ProductId == item.Id {
				r = data
			}
		}
		if r != nil {
			list = append(list, &types.ListFlashPromotionProductRelationData{
				Id:                                     r.Id,
				FlashPromotionProductRelationID:        r.FlashPromotionId,
				FlashPromotionProductRelationSessionID: r.FlashPromotionSessionId,
				ProductID:                              r.ProductId,
				FlashPromotionProductRelationPrice:     r.FlashPromotionPrice,
				FlashPromotionProductRelationCount:     r.FlashPromotionCount,
				FlashPromotionProductRelationLimit:     r.FlashPromotionLimit,
				Sort:                                   r.Sort,
				ProductId:                              item.Id,
				Name:                                   item.Name,
				ProductSn:                              item.ProductSn,
				Price:                                  item.Price,
				Stock:                                  item.Stock,
			})
		}
	}

	return &types.ListFlashPromotionProductRelationResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
		Code:     "000000",
		Message:  "分页查询不同场次关联及商品信息成功",
	}, nil
}
