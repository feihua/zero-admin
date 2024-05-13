package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberRuleSettingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberRuleSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberRuleSettingListLogic {
	return MemberRuleSettingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberRuleSettingListLogic) MemberRuleSettingList(req types.ListMemberRuleSettingReq) (*types.ListMemberRuleSettingResp, error) {
	resp, err := l.svcCtx.MemberRuleSettingService.MemberRuleSettingList(l.ctx, &umsclient.MemberRuleSettingListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询会员积分规则列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询会员积分规则失败")
	}

	var list []*types.ListMemberRuleSettingData

	for _, item := range resp.List {
		list = append(list, &types.ListMemberRuleSettingData{
			Id:                item.Id,
			ContinueSignDay:   item.ContinueSignDay,
			ContinueSignPoint: item.ContinueSignPoint,
			ConsumePerPoint:   item.ConsumePerPoint,
			LowOrderAmount:    item.LowOrderAmount,
			MaxPointPerOrder:  item.MaxPointPerOrder,
			Type:              item.RuleType,
		})
	}

	return &types.ListMemberRuleSettingResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询会员积分规则成功",
	}, nil
}
