package flashpromotionservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
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
	q := query.SmsFlashPromotion

	flash, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).First()
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errors.New("活动不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询活动异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询活动异常")
	}

	title := in.Title
	count, err := q.WithContext(l.ctx).Where(q.Title.Eq(title), q.ID.Neq(in.Id)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "更新限时购失败,参数：%+v, 异常:%s", in, err.Error())
		return nil, errors.New(fmt.Sprintf("更新加限时购失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("活动标题：%s,已存在", title))
	}

	StartDate, _ := time.Parse("2006-01-02", in.StartDate)
	EndDate, _ := time.Parse("2006-01-02", in.EndDate)

	err = l.svcCtx.DB.Model(&model.SmsFlashPromotion{}).WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Save(&model.SmsFlashPromotion{
		ID:         in.Id,            // 编号
		Title:      in.Title,         // 标题
		StartDate:  StartDate,        // 开始日期
		EndDate:    EndDate,          // 结束日期
		Status:     in.Status,        // 上下线状态
		CreateTime: flash.CreateTime, // 创建时间
	}).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新限时购表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新限时购表失败")
	}

	return &smsclient.UpdateFlashPromotionResp{}, nil
}
