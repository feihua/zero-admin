package flashpromotionservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteFlashPromotionLogic 删除限时购表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:42
*/
type DeleteFlashPromotionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFlashPromotionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFlashPromotionLogic {
	return &DeleteFlashPromotionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteFlashPromotion 删除限时购表
func (l *DeleteFlashPromotionLogic) DeleteFlashPromotion(in *smsclient.DeleteFlashPromotionReq) (*smsclient.DeleteFlashPromotionResp, error) {
	q := query.SmsFlashPromotion
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除限时购表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除限时购表失败")
	}

	return &smsclient.DeleteFlashPromotionResp{}, nil
}
