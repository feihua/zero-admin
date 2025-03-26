package flashpromotionproductrelationservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteFlashPromotionProductRelationLogic 删除商品限时购与商品关系
/*
Author: LiuFeiHua
Date: 2024/6/12 17:40
*/
type DeleteFlashPromotionProductRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFlashPromotionProductRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFlashPromotionProductRelationLogic {
	return &DeleteFlashPromotionProductRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteFlashPromotionProductRelation 删除商品限时购与商品关系
func (l *DeleteFlashPromotionProductRelationLogic) DeleteFlashPromotionProductRelation(in *smsclient.DeleteFlashPromotionProductRelationReq) (*smsclient.DeleteFlashPromotionProductRelationResp, error) {
	q := query.SmsFlashPromotionProductRelation
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...), q.FlashPromotionID.Eq(in.FlashPromotionId)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除商品限时购与商品关系失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除商品限时购与商品关系失败")
	}

	return &smsclient.DeleteFlashPromotionProductRelationResp{}, nil
}
