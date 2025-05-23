package memberconsumesettingservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

// UpdateMemberConsumeSettingLogic 更新积分消费设置
/*
Author: LiuFeiHua
Date: 2025/05/23 11:32:02
*/
type UpdateMemberConsumeSettingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberConsumeSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberConsumeSettingLogic {
	return &UpdateMemberConsumeSettingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberConsumeSetting 更新积分消费设置
func (l *UpdateMemberConsumeSettingLogic) UpdateMemberConsumeSetting(in *umsclient.UpdateMemberConsumeSettingReq) (*umsclient.UpdateMemberConsumeSettingResp, error) {
	q := query.UmsMemberConsumeSetting

	// 1.根据积分消费设置id查询积分消费设置是否已存在
	setting, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "积分消费设置不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("积分消费设置不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询积分消费设置异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询积分消费设置异常")
	}

	now := time.Now()
	item := &model.UmsMemberConsumeSetting{
		ID:                 in.Id,                 //
		DeductionPerAmount: in.DeductionPerAmount, // 每一元需要抵扣的积分数量
		MaxPercentPerOrder: in.MaxPercentPerOrder, // 每笔订单最高抵用百分比
		UseUnit:            in.UseUnit,            // 每次使用积分最小单位100
		CouponStatus:       in.CouponStatus,       // 是否可以和优惠券同用；0->不可以；1->可以
		Status:             in.Status,             // 状态：0->禁用；1->启用
		CreateBy:           setting.CreateBy,      // 创建人ID
		CreateTime:         setting.CreateTime,    // 创建时间
		UpdateBy:           &in.UpdateBy,          // 更新人ID
		UpdateTime:         &now,                  // 更新时间
	}

	if in.Status == 1 {
		_, _ = q.WithContext(l.ctx).Where(q.ID.Neq(in.Id), q.Status.Eq(in.Status)).Update(q.Status, 0)
	}
	// 2.积分消费设置存在时,则直接更新积分消费设置
	err = l.svcCtx.DB.Model(&model.UmsMemberConsumeSetting{}).WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新积分消费设置失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新积分消费设置失败")
	}

	return &umsclient.UpdateMemberConsumeSettingResp{}, nil
}
