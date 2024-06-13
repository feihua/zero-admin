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

// AddFlashPromotionLogic 添加限时购表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:42
*/
type AddFlashPromotionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFlashPromotionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFlashPromotionLogic {
	return &AddFlashPromotionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddFlashPromotion 添加限时购表
func (l *AddFlashPromotionLogic) AddFlashPromotion(in *smsclient.AddFlashPromotionReq) (*smsclient.AddFlashPromotionResp, error) {
	StartDate, _ := time.Parse("2006-01-02", in.StartDate)
	EndDate, _ := time.Parse("2006-01-02", in.EndDate)

	err := query.SmsFlashPromotion.WithContext(l.ctx).Create(&model.SmsFlashPromotion{
		Title:      in.Title,
		StartDate:  StartDate,
		EndDate:    EndDate,
		Status:     in.Status,
		CreateTime: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &smsclient.AddFlashPromotionResp{}, nil
}
