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
	"gorm.io/gorm"
	"time"
)

// UpdateMemberRuleSettingLogic 更新会员积分成长规则
/*
Author: LiuFeiHua
Date: 2025/05/23 11:19:51
*/
type UpdateMemberRuleSettingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberRuleSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberRuleSettingLogic {
	return &UpdateMemberRuleSettingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberRuleSetting 更新会员积分成长规则
func (l *UpdateMemberRuleSettingLogic) UpdateMemberRuleSetting(in *umsclient.UpdateMemberRuleSettingReq) (*umsclient.UpdateMemberRuleSettingResp, error) {
	q := query.UmsMemberRuleSetting.WithContext(l.ctx)

	// 1.根据会员积分成长规则id查询会员积分成长规则是否已存在
	rule, err := q.Where(query.UmsMemberRuleSetting.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "会员积分成长规则不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("会员积分成长规则不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询会员积分成长规则异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询会员积分成长规则异常")
	}

	now := time.Now()
	item := &model.UmsMemberRuleSetting{
		ID:               in.Id,               //
		ConsumePerPoint:  in.ConsumePerPoint,  // 每消费多少元获取1个点
		LowOrderAmount:   in.LowOrderAmount,   // 最低获取点数的订单金额
		MaxPointPerOrder: in.MaxPointPerOrder, // 每笔订单最高获取点数
		RuleType:         in.RuleType,         // 类型：0->积分规则；1->成长值规则
		Status:           in.Status,           // 状态：0->禁用；1->启用
		CreateBy:         rule.CreateBy,       // 创建人ID
		CreateTime:       rule.CreateTime,     // 创建时间
		UpdateBy:         &in.UpdateBy,        // 更新人ID
		UpdateTime:       &now,                // 更新时间
	}

	// 2.会员积分成长规则存在时,则直接更新会员积分成长规则
	err = l.svcCtx.DB.Model(&model.UmsMemberRuleSetting{}).WithContext(l.ctx).Where(query.UmsMemberRuleSetting.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新会员积分成长规则失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新会员积分成长规则失败")
	}

	return &umsclient.UpdateMemberRuleSettingResp{}, nil
}
