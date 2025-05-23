package member_rule_setting

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
func (l *QueryMemberRuleSettingListLogic) QueryMemberRuleSettingList(req *types.QueryMemberRuleSettingListReq) (resp *types.QueryMemberRuleSettingListResp, err error) {
	result, err := l.svcCtx.MemberRuleSettingService.QueryMemberRuleSettingList(l.ctx, &umsclient.QueryMemberRuleSettingListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
		RuleType: req.RuleType,
		Status:   req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询会员积分规则列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询会员积分规则失败")
	}

	var list []*types.QueryMemberRuleSettingListData

	for _, item := range result.List {
		list = append(list, &types.QueryMemberRuleSettingListData{
			Id:               item.Id,               //
			ConsumePerPoint:  item.ConsumePerPoint,  // 每消费多少元获取1个点
			LowOrderAmount:   item.LowOrderAmount,   // 最低获取点数的订单金额
			MaxPointPerOrder: item.MaxPointPerOrder, // 每笔订单最高获取点数
			RuleType:         item.RuleType,         // 类型：0->积分规则；1->成长值规则
			Status:           item.Status,           // 状态：0->禁用；1->启用
			CreateBy:         item.CreateBy,         // 创建者
			CreateTime:       item.CreateTime,       // 创建时间
			UpdateBy:         item.UpdateBy,         // 更新者
			UpdateTime:       item.UpdateTime,       // 更新时间
		})
	}

	return &types.QueryMemberRuleSettingListResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询会员积分规则成功",
	}, nil
}
