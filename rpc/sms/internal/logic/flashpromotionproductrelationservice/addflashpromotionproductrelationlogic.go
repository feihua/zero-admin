package flashpromotionproductrelationservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddFlashPromotionProductRelationLogic 添加商品限时购与商品关系
/*
Author: LiuFeiHua
Date: 2024/6/12 17:40
*/
type AddFlashPromotionProductRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFlashPromotionProductRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFlashPromotionProductRelationLogic {
	return &AddFlashPromotionProductRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddFlashPromotionProductRelation 添加商品限时购与商品关系
func (l *AddFlashPromotionProductRelationLogic) AddFlashPromotionProductRelation(in *smsclient.AddFlashPromotionProductRelationReq) (*smsclient.AddFlashPromotionProductRelationResp, error) {
	var data []*model.SmsFlashPromotionProductRelation
	for _, item := range in.Data {
		data = append(data, &model.SmsFlashPromotionProductRelation{
			FlashPromotionID:        item.FlashPromotionId,        // 限时购id
			FlashPromotionSessionID: item.FlashPromotionSessionId, // 编号
			ProductID:               item.ProductId,               // 商品id
			FlashPromotionPrice:     item.FlashPromotionPrice,     // 限时购价格
			FlashPromotionCount:     item.FlashPromotionCount,     // 限时购数量
			FlashPromotionLimit:     item.FlashPromotionLimit,     // 每人限购数量
			Sort:                    item.Sort,                    // 排序
		})
	}

	err := query.SmsFlashPromotionProductRelation.WithContext(l.ctx).Create(data...)

	if err != nil {
		logc.Errorf(l.ctx, "添加商品限时购与商品关系失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加商品限时购与商品关系失败")
	}

	return &smsclient.AddFlashPromotionProductRelationResp{}, nil
}
