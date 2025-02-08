package flashpromotionproductrelation

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

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
	result, err := l.svcCtx.FlashPromotionProductRelationService.QueryFlashPromotionProductRelationList(l.ctx, &smsclient.QueryFlashPromotionProductRelationListReq{
		PageNum:                 req.Current,
		PageSize:                req.PageSize,
		FlashPromotionId:        req.FlashPromotionProductRelationID,
		FlashPromotionSessionId: req.FlashPromotionProductRelationSessionID,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,分页查询不同场次关联及商品信息列表异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var ids []int64
	for _, item := range result.List {
		ids = append(ids, item.ProductId)
	}
	var list []*types.ListFlashPromotionProductRelationData

	products, err := l.svcCtx.ProductService.QueryProductListByIds(l.ctx, &pmsclient.QueryProductByIdsReq{Ids: ids})
	for _, item := range products.List {
		var r *smsclient.FlashPromotionProductRelationListData
		for _, data := range result.List {
			if data.ProductId == item.Id {
				r = data
			}
		}
		if r != nil {
			list = append(list, &types.ListFlashPromotionProductRelationData{
				Id:                                     r.Id,                      // 编号
				FlashPromotionProductRelationID:        r.FlashPromotionId,        // 限时购id
				FlashPromotionProductRelationSessionID: r.FlashPromotionSessionId, // 编号
				FlashPromotionProductRelationPrice:     r.FlashPromotionPrice,     // 限时购价格
				FlashPromotionProductRelationCount:     r.FlashPromotionCount,     // 限时购数量
				FlashPromotionProductRelationLimit:     r.FlashPromotionLimit,     // 每人限购数量
				Sort:                                   r.Sort,                    // 排序
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
