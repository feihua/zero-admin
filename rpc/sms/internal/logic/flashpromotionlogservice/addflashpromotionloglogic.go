package flashpromotionlogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddFlashPromotionLogLogic 添加限时购通知记录
/*
Author: LiuFeiHua
Date: 2024/6/12 17:38
*/
type AddFlashPromotionLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFlashPromotionLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFlashPromotionLogLogic {
	return &AddFlashPromotionLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddFlashPromotionLog 添加限时购通知记录
func (l *AddFlashPromotionLogLogic) AddFlashPromotionLog(in *smsclient.AddFlashPromotionLogReq) (*smsclient.AddFlashPromotionLogResp, error) {
	SubscribeTime, _ := time.Parse("2006-01-02 15:04:05", in.SubscribeTime)
	err := query.SmsFlashPromotionLog.WithContext(l.ctx).Create(&model.SmsFlashPromotionLog{
		MemberID:      in.MemberId,    // 会员id
		ProductID:     in.ProductId,   // 商品id
		MemberPhone:   in.MemberPhone, // 会员电话
		ProductName:   in.ProductName, // 商品名称
		SubscribeTime: SubscribeTime,  // 会员订阅时间
		SendTime:      time.Now(),     // 发送时间
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加限时购通知记录失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加限时购通知记录失败")
	}

	return &smsclient.AddFlashPromotionLogResp{}, nil
}
