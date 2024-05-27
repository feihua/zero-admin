package memberrulesettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/logic/common"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberRuleSettingListLogic 会员积分成长规则
/*
Author: LiuFeiHua
Date: 2024/5/23 11:02
*/
type MemberRuleSettingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberRuleSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberRuleSettingListLogic {
	return &MemberRuleSettingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberRuleSettingList 查询会员积分成长规则
func (l *MemberRuleSettingListLogic) MemberRuleSettingList(in *umsclient.MemberRuleSettingListReq) (*umsclient.MemberRuleSettingListResp, error) {
	q := query.UmsMemberRuleSetting.WithContext(l.ctx)
	result, count, err := q.FindByPage(int((in.Current-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询积分规则列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*umsclient.MemberRuleSettingListData
	for _, item := range result {

		list = append(list, &umsclient.MemberRuleSettingListData{
			Id:                item.ID,
			ContinueSignDay:   item.ContinueSignDay,
			ContinueSignPoint: item.ContinueSignPoint,
			ConsumePerPoint:   item.ConsumePerPoint,
			LowOrderAmount:    item.LowOrderAmount,
			MaxPointPerOrder:  item.MaxPointPerOrder,
			RuleType:          item.RuleType,
			CreateBy:          item.CreateBy,
			CreateTime:        item.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateBy:          item.UpdateBy,
			UpdateTime:        common.TimeToString(item.UpdateTime),
		})
	}

	logc.Infof(l.ctx, "查询会员规则设置列表信息,参数：%+v,响应：%+v", in, list)
	return &umsclient.MemberRuleSettingListResp{
		Total: count,
		List:  list,
	}, nil

}
