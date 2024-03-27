package memberrulesettingservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

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
	all, err := l.svcCtx.UmsMemberRuleSettingModel.FindAll(l.ctx, in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsMemberRuleSettingModel.Count(l.ctx)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询积分规则列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*umsclient.MemberRuleSettingListData
	for _, item := range *all {

		list = append(list, &umsclient.MemberRuleSettingListData{
			Id:                item.Id,
			ContinueSignDay:   item.ContinueSignDay,
			ContinueSignPoint: item.ContinueSignPoint,
			ConsumePerPoint:   int64(item.ConsumePerPoint),
			LowOrderAmount:    int64(item.LowOrderAmount),
			MaxPointPerOrder:  item.MaxPointPerOrder,
			Type:              item.Type,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询会员规则设置列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &umsclient.MemberRuleSettingListResp{
		Total: count,
		List:  list,
	}, nil

}
