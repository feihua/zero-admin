package memberconsumesettingservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QueryMemberConsumeSettingDetailLogic 查询积分消费设置详情
/*
Author: LiuFeiHua
Date: 2025/05/23 11:32:02
*/
type QueryMemberConsumeSettingDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberConsumeSettingDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberConsumeSettingDetailLogic {
	return &QueryMemberConsumeSettingDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberConsumeSettingDetail 查询积分消费设置详情
func (l *QueryMemberConsumeSettingDetailLogic) QueryMemberConsumeSettingDetail(in *umsclient.QueryMemberConsumeSettingDetailReq) (*umsclient.QueryMemberConsumeSettingDetailResp, error) {
	item, err := query.UmsMemberConsumeSetting.WithContext(l.ctx).Where(query.UmsMemberConsumeSetting.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "积分消费设置不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("积分消费设置不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询积分消费设置异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询积分消费设置异常")
	}

	data := &umsclient.QueryMemberConsumeSettingDetailResp{
		Id:                 item.ID,                                          //
		DeductionPerAmount: item.DeductionPerAmount,                          // 每一元需要抵扣的积分数量
		MaxPercentPerOrder: item.MaxPercentPerOrder,                          // 每笔订单最高抵用百分比
		UseUnit:            item.UseUnit,                                     // 每次使用积分最小单位100
		CouponStatus:       item.CouponStatus,                                // 是否可以和优惠券同用；0->不可以；1->可以
		Status:             item.Status,                                      // 状态：0->禁用；1->启用
		CreateBy:           item.CreateBy,                                    // 创建人ID
		CreateTime:         time_util.TimeToStr(item.CreateTime),             // 创建时间
		UpdateBy:           pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
		UpdateTime:         time_util.TimeToString(item.UpdateTime),          // 更新时间
	}

	return data, nil
}
