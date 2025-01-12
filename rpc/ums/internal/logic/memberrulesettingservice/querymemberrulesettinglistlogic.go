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
			ConsumePerPoint:   item.ConsumePerPoint,
			ContinueSignDay:   item.ContinueSignDay,
			ContinueSignPoint: item.ContinueSignPoint,
			CreateBy:          item.CreateBy,
			CreateTime:        item.CreateTime.Format("2006-01-02 15:04:05"),
			Id:                item.ID,
			LowOrderAmount:    item.LowOrderAmount,
			MaxPointPerOrder:  item.MaxPointPerOrder,
			RuleType:          item.RuleType,
			Status:            item.Status,
			UpdateBy:          item.UpdateBy,
			UpdateTime:        time_util.TimeToString(item.UpdateTime),
		})
	}

	return &umsclient.QueryMemberRuleSettingListResp{
		Total: count,
		List:  list,
	}, nil
}
