package flashpromotionsessionservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateFlashPromotionSessionLogic 更新限时购场次
/*
Author: LiuFeiHua
Date: 2024/6/12 17:50
*/
type UpdateFlashPromotionSessionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFlashPromotionSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFlashPromotionSessionLogic {
	return &UpdateFlashPromotionSessionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateFlashPromotionSession 更新限时购场次
func (l *UpdateFlashPromotionSessionLogic) UpdateFlashPromotionSession(in *smsclient.UpdateFlashPromotionSessionReq) (*smsclient.UpdateFlashPromotionSessionResp, error) {
	q := query.SmsFlashPromotionSession
	_, err := q.WithContext(l.ctx).Updates(&model.SmsFlashPromotionSession{
		ID:        in.Id,        // 编号
		Name:      in.Name,      // 场次名称
		StartTime: in.StartTime, // 每日开始时间
		EndTime:   in.EndTime,   // 每日结束时间
		Status:    in.Status,    // 启用状态：0->不启用；1->启用
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新限时购场次失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新限时购场次失败")
	}

	return &smsclient.UpdateFlashPromotionSessionResp{}, nil
}
