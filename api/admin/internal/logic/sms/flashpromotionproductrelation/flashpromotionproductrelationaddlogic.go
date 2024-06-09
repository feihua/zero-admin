package flashpromotionproductrelation

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// FlashPromotionProductRelationAddLogic 限时购和商品关系
/*
Author: LiuFeiHua
Date: 2024/5/14 15:30
*/
type FlashPromotionProductRelationAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionProductRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionProductRelationAddLogic {
	return &FlashPromotionProductRelationAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// FlashPromotionProductRelationAdd 添加限时购和商品关系
func (l *FlashPromotionProductRelationAddLogic) FlashPromotionProductRelationAdd(req *types.AddFlashPromotionProductRelationReq) (resp *types.AddFlashPromotionProductRelationResp, err error) {

	var list []*smsclient.FlashPromotionProductRelationAddData
	for _, item := range req.Data {
		list = append(list, &smsclient.FlashPromotionProductRelationAddData{
			FlashPromotionId:        item.FlashPromotionProductRelationID,
			FlashPromotionSessionId: item.FlashPromotionProductRelationSessionID,
			ProductId:               item.ProductID,
			FlashPromotionPrice:     item.FlashPromotionProductRelationPrice,
			FlashPromotionCount:     item.FlashPromotionProductRelationCount,
			FlashPromotionLimit:     item.FlashPromotionProductRelationLimit,
			Sort:                    item.Sort,
		})
	}

	_, err = l.svcCtx.FlashPromotionProductRelationService.FlashPromotionProductRelationAdd(l.ctx, &smsclient.FlashPromotionProductRelationAddReq{
		Data: list,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加限时购和商品关系失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加限时购和商品关系失败")
	}

	return &types.AddFlashPromotionProductRelationResp{
		Code:    "000000",
		Message: "添加限时购和商品关系成功",
	}, nil

}
