package flashpromotionlogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteFlashPromotionLogLogic 删除限时购通知记录
/*
Author: LiuFeiHua
Date: 2024/6/12 17:39
*/
type DeleteFlashPromotionLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFlashPromotionLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFlashPromotionLogLogic {
	return &DeleteFlashPromotionLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteFlashPromotionLog 删除限时购通知记录
func (l *DeleteFlashPromotionLogLogic) DeleteFlashPromotionLog(in *smsclient.DeleteFlashPromotionLogReq) (*smsclient.DeleteFlashPromotionLogResp, error) {
	q := query.SmsFlashPromotionLog
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除限时购通知记录失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除限时购通知记录失败")
	}

	return &smsclient.DeleteFlashPromotionLogResp{}, nil
}
