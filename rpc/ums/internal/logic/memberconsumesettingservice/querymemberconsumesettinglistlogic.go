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
)

// QueryMemberConsumeSettingListLogic 查询积分消费设置列表
/*
Author: LiuFeiHua
Date: 2025/05/23 11:32:02
*/
type QueryMemberConsumeSettingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberConsumeSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberConsumeSettingListLogic {
	return &QueryMemberConsumeSettingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberConsumeSettingList 查询积分消费设置列表
func (l *QueryMemberConsumeSettingListLogic) QueryMemberConsumeSettingList(in *umsclient.QueryMemberConsumeSettingListReq) (*umsclient.QueryMemberConsumeSettingListResp, error) {
	memberConsumeSetting := query.UmsMemberConsumeSetting
	q := memberConsumeSetting.WithContext(l.ctx)

	if in.CouponStatus != 2 {
		q = q.Where(memberConsumeSetting.CouponStatus.Eq(in.CouponStatus))
	}
	if in.Status != 2 {
		q = q.Where(memberConsumeSetting.Status.Eq(in.Status))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询积分消费设置列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询积分消费设置列表失败")
	}

	var list []*umsclient.MemberConsumeSettingListData

	for _, item := range result {
		list = append(list, &umsclient.MemberConsumeSettingListData{
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

		})
	}

	return &umsclient.QueryMemberConsumeSettingListResp{
		Total: count,
		List:  list,
	}, nil
}
