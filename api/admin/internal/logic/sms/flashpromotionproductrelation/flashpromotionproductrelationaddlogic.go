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
			FlashPromotionId:        item.FlashPromotionProductRelationID,        // 限时购id
			FlashPromotionSessionId: item.FlashPromotionProductRelationSessionID, // 编号
			ProductId:               item.ProductID,                              // 商品id
			FlashPromotionPrice:     item.FlashPromotionProductRelationPrice,     // 限时购价格
			FlashPromotionCount:     item.FlashPromotionProductRelationCount,     // 限时购数量
			FlashPromotionLimit:     item.FlashPromotionProductRelationLimit,     // 每人限购数量
			Sort:                    item.Sort,                                   // 排序
		})
	}

	_, err = l.svcCtx.FlashPromotionProductRelationService.AddFlashPromotionProductRelation(l.ctx, &smsclient.AddFlashPromotionProductRelationReq{
		Data: list,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加限时购和商品关系失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddFlashPromotionProductRelationResp{
		Code:    "000000",
		Message: "添加限时购和商品关系成功",
	}, nil

}
