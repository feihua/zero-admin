package flashpromotionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateFlashPromotionLogic 更新限时购表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:44
*/
type UpdateFlashPromotionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFlashPromotionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFlashPromotionLogic {
	return &UpdateFlashPromotionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateFlashPromotion 更新限时购表
func (l *UpdateFlashPromotionLogic) UpdateFlashPromotion(in *smsclient.UpdateFlashPromotionReq) (*smsclient.UpdateFlashPromotionResp, error) {
	StartDate, _ := time.Parse("2006-01-02", in.StartDate)
	EndDate, _ := time.Parse("2006-01-02", in.EndDate)

	q := query.SmsFlashPromotion
	_, err := q.WithContext(l.ctx).Updates(&model.SmsFlashPromotion{
		ID:        in.Id,     // 编号
		Title:     in.Title,  // 标题
		StartDate: StartDate, // 开始日期
		EndDate:   EndDate,   // 结束日期
		Status:    in.Status, // 上下线状态
	})

	if err != nil {
		return nil, err
	}

	return &smsclient.UpdateFlashPromotionResp{}, nil
}
