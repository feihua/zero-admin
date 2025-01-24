package memberrulesettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberRuleSettingListLogic 查询会员积分成长规则表列表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:01
*/
type QueryMemberRuleSettingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberRuleSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberRuleSettingListLogic {
	return &QueryMemberRuleSettingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberRuleSettingList 查询会员积分成长规则表列表
func (l *QueryMemberRuleSettingListLogic) QueryMemberRuleSettingList(in *umsclient.QueryMemberRuleSettingListReq) (*umsclient.QueryMemberRuleSettingListResp, error) {
	q := query.UmsMemberRuleSetting.WithContext(l.ctx)
	if in.Status != 2 {
		q = q.Where(query.UmsMemberRuleSetting.Status.Eq(in.Status))
	}
	if in.RuleType != 2 {
		q = q.Where(query.UmsMemberRuleSetting.RuleType.Eq(in.RuleType))
	}
	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询积分规则列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*umsclient.MemberRuleSettingListData
	for _, item := range result {

		list = append(list, &umsclient.MemberRuleSettingListData{
			Id:                item.ID,                                 //
			ContinueSignDay:   item.ContinueSignDay,                    // 连续签到天数
			ContinueSignPoint: item.ContinueSignPoint,                  // 连续签到赠送数量
			ConsumePerPoint:   item.ConsumePerPoint,                    // 每消费多少元获取1个点
			LowOrderAmount:    item.LowOrderAmount,                     // 最低获取点数的订单金额
			MaxPointPerOrder:  item.MaxPointPerOrder,                   // 每笔订单最高获取点数
			RuleType:          item.RuleType,                           // 类型：0->积分规则；1->成长值规则
			Status:            item.Status,                             // 状态：0->禁用；1->启用
			CreateBy:          item.CreateBy,                           // 创建者
			CreateTime:        time_util.TimeToStr(item.CreateTime),    // 创建时间
			UpdateBy:          item.UpdateBy,                           // 更新者
			UpdateTime:        time_util.TimeToString(item.UpdateTime), // 更新时间

		})
	}

	return &umsclient.QueryMemberRuleSettingListResp{
		Total: count,
		List:  list,
	}, nil
}
