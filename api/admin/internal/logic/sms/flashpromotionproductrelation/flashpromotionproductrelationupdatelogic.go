package flashpromotionproductrelation

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// FlashPromotionProductRelationUpdateLogic 限时购和商品关系
/*
Author: LiuFeiHua
Date: 2024/5/14 15:31
*/
type FlashPromotionProductRelationUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionProductRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionProductRelationUpdateLogic {
	return &FlashPromotionProductRelationUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// FlashPromotionProductRelationUpdate 更新限时购和商品关系
func (l *FlashPromotionProductRelationUpdateLogic) FlashPromotionProductRelationUpdate(req *types.UpdateFlashPromotionProductRelationReq) (resp *types.UpdateFlashPromotionProductRelationResp, err error) {
	_, err = l.svcCtx.FlashPromotionProductRelationService.UpdateFlashPromotionProductRelation(l.ctx, &smsclient.UpdateFlashPromotionProductRelationReq{
		Id:                      req.ID,
		FlashPromotionId:        req.FlashPromotionProductRelationID,        // 限时购id
		FlashPromotionSessionId: req.FlashPromotionProductRelationSessionID, // 编号
		ProductId:               req.ProductID,                              // 商品id
		FlashPromotionPrice:     req.FlashPromotionProductRelationPrice,     // 限时购价格
		FlashPromotionCount:     req.FlashPromotionProductRelationCount,     // 限时购数量
		FlashPromotionLimit:     req.FlashPromotionProductRelationLimit,     // 每人限购数量
		Sort:                    req.Sort,                                   // 排序
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新限时购和商品关系息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateFlashPromotionProductRelationResp{
		Code:    "000000",
		Message: "更新限时购和商品关系成功",
	}, nil
}
