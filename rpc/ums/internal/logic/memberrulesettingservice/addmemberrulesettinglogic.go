package memberrulesettingservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberRuleSettingLogic 添加会员积分成长规则
/*
Author: LiuFeiHua
Date: 2025/05/23 11:19:51
*/
type AddMemberRuleSettingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMemberRuleSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberRuleSettingLogic {
	return &AddMemberRuleSettingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddMemberRuleSetting 添加会员积分成长规则
func (l *AddMemberRuleSettingLogic) AddMemberRuleSetting(in *umsclient.AddMemberRuleSettingReq) (*umsclient.AddMemberRuleSettingResp, error) {
	q := query.UmsMemberRuleSetting

	item := &model.UmsMemberRuleSetting{
		ConsumePerPoint:  in.ConsumePerPoint,  // 每消费多少元获取1个点
		LowOrderAmount:   in.LowOrderAmount,   // 最低获取点数的订单金额
		MaxPointPerOrder: in.MaxPointPerOrder, // 每笔订单最高获取点数
		RuleType:         in.RuleType,         // 类型：0->积分规则；1->成长值规则
		Status:           in.Status,           // 状态：0->禁用；1->启用
		CreateBy:         in.CreateBy,         // 创建人ID
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加会员积分成长规则失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加会员积分成长规则失败")
	}

	return &umsclient.AddMemberRuleSettingResp{}, nil
}
