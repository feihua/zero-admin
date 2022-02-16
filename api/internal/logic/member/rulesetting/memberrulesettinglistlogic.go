package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

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
	resp, err := l.svcCtx.Ums.MemberRuleSettingList(l.ctx, &umsclient.MemberRuleSettingListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询会员积分规则列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询会员积分规则失败")
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtMemberRuleSettingData

	for _, item := range resp.List {
		list = append(list, &types.ListtMemberRuleSettingData{
			Id:                item.Id,
			ContinueSignDay:   item.ContinueSignDay,
			ContinueSignPoint: item.ContinueSignPoint,
			ConsumePerPoint:   float64(item.ConsumePerPoint),
			LowOrderAmount:    float64(item.LowOrderAmount),
			MaxPointPerOrder:  item.MaxPointPerOrder,
			Type:              item.Type,
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
