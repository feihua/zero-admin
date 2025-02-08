package flashpromotionservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddFlashPromotionLogic 添加限时购
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

// AddFlashPromotion 添加限时购
func (l *AddFlashPromotionLogic) AddFlashPromotion(in *smsclient.AddFlashPromotionReq) (*smsclient.AddFlashPromotionResp, error) {
	q := query.SmsFlashPromotion
	title := in.Title
	count, err := q.WithContext(l.ctx).Where(q.Title.Eq(title)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "添加限时购失败,参数：%+v, 异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("添加限时购失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("活动标题：%s,已存在", title))
	}

	StartDate, _ := time.Parse("2006-01-02", in.StartDate)
	EndDate, _ := time.Parse("2006-01-02", in.EndDate)

	err = q.WithContext(l.ctx).Create(&model.SmsFlashPromotion{
		Title:     in.Title,  // 标题
		StartDate: StartDate, // 开始日期
		EndDate:   EndDate,   // 结束日期
		Status:    in.Status, // 上下线状态
	})

	if err != nil {
		return nil, err
	}

	return &smsclient.AddFlashPromotionResp{}, nil
}
