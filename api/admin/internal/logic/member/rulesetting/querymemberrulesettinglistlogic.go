package rulesetting

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberRuleSettingListLogic 查询会员积分规则列表
/*
Author: LiuFeiHua
Date: 2024/5/23 9:25
*/
type QueryMemberRuleSettingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberRuleSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberRuleSettingListLogic {
	return &QueryMemberRuleSettingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberRuleSettingList 查询会员积分规则列表
func (l *QueryMemberRuleSettingListLogic) QueryMemberRuleSettingList(req *types.ListMemberRuleSettingReq) (resp *types.ListMemberRuleSettingResp, err error) {
	result, err := l.svcCtx.MemberRuleSettingService.MemberRuleSettingList(l.ctx, &umsclient.MemberRuleSettingListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询会员积分规则列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询会员积分规则失败")
	}

	var list []*types.ListMemberRuleSettingData

	for _, item := range result.List {
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
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询会员积分规则成功",
	}, nil
}
