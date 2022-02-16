package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

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

func (l *MemberRuleSettingListLogic) MemberRuleSettingList(in *ums.MemberRuleSettingListReq) (*ums.MemberRuleSettingListResp, error) {
	all, err := l.svcCtx.UmsMemberRuleSettingModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsMemberRuleSettingModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询积分规则列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*ums.MemberRuleSettingListData
	for _, item := range *all {

		list = append(list, &ums.MemberRuleSettingListData{
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
	return &ums.MemberRuleSettingListResp{
		Total: count,
		List:  list,
	}, nil

}
