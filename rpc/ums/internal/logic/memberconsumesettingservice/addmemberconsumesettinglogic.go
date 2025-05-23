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
)

// AddMemberConsumeSettingLogic 添加积分消费设置
/*
Author: LiuFeiHua
Date: 2025/05/23 11:32:02
*/
type AddMemberConsumeSettingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberConsumeSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberConsumeSettingLogic {
	return &AddMemberConsumeSettingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberConsumeSetting 添加积分消费设置
func (l *AddMemberConsumeSettingLogic) AddMemberConsumeSetting(in *umsclient.AddMemberConsumeSettingReq) (*umsclient.AddMemberConsumeSettingResp, error) {
	q := query.UmsMemberConsumeSetting

	item := &model.UmsMemberConsumeSetting{
		DeductionPerAmount: in.DeductionPerAmount, // 每一元需要抵扣的积分数量
		MaxPercentPerOrder: in.MaxPercentPerOrder, // 每笔订单最高抵用百分比
		UseUnit:            in.UseUnit,            // 每次使用积分最小单位100
		CouponStatus:       in.CouponStatus,       // 是否可以和优惠券同用；0->不可以；1->可以
		Status:             in.Status,             // 状态：0->禁用；1->启用
		CreateBy:           in.CreateBy,           // 创建人ID
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加积分消费设置失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加积分消费设置失败")
	}

	return &umsclient.AddMemberConsumeSettingResp{}, nil
}
