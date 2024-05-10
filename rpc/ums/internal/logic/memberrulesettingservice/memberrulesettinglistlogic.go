package memberrulesettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *MemberRuleSettingListLogic) MemberRuleSettingList(in *umsclient.MemberRuleSettingListReq) (*umsclient.MemberRuleSettingListResp, error) {
	q := query.UmsMemberRuleSetting.WithContext(l.ctx)
	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

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
			ConsumePerPoint:   float32(item.ConsumePerPoint),
			LowOrderAmount:    float32(item.LowOrderAmount),
			MaxPointPerOrder:  item.MaxPointPerOrder,
			RuleType:          item.RuleType,
			CreateBy:          item.CreateBy,
			CreateTime:        item.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateBy:          *item.UpdateBy,
			UpdateTime:        item.UpdateTime.Format("2006-01-02 15:04:05"),
		})
	}

	logc.Infof(l.ctx, "查询会员规则设置列表信息,参数：%+v,响应：%+v", in, list)
	return &umsclient.MemberRuleSettingListResp{
		Total: count,
		List:  list,
	}, nil

}
