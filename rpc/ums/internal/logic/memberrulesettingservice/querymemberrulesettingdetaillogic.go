package memberrulesettingservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QueryMemberRuleSettingDetailLogic 查询会员积分成长规则详情
/*
Author: LiuFeiHua
Date: 2025/05/23 11:19:51
*/
type QueryMemberRuleSettingDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberRuleSettingDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberRuleSettingDetailLogic {
	return &QueryMemberRuleSettingDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberRuleSettingDetail 查询会员积分成长规则详情
func (l *QueryMemberRuleSettingDetailLogic) QueryMemberRuleSettingDetail(in *umsclient.QueryMemberRuleSettingDetailReq) (*umsclient.QueryMemberRuleSettingDetailResp, error) {
	item, err := query.UmsMemberRuleSetting.WithContext(l.ctx).Where(query.UmsMemberRuleSetting.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "会员积分成长规则不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("会员积分成长规则不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询会员积分成长规则异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询会员积分成长规则异常")
	}

	data := &umsclient.QueryMemberRuleSettingDetailResp{
		Id:               item.ID,                                          //
		ConsumePerPoint:  item.ConsumePerPoint,                             // 每消费多少元获取1个点
		LowOrderAmount:   item.LowOrderAmount,                              // 最低获取点数的订单金额
		MaxPointPerOrder: item.MaxPointPerOrder,                            // 每笔订单最高获取点数
		RuleType:         item.RuleType,                                    // 类型：0->积分规则；1->成长值规则
		Status:           item.Status,                                      // 状态：0->禁用；1->启用
		CreateBy:         item.CreateBy,                                    // 创建人ID
		CreateTime:       time_util.TimeToStr(item.CreateTime),             // 创建时间
		UpdateBy:         pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
		UpdateTime:       time_util.TimeToString(item.UpdateTime),          // 更新时间
	}

	return data, nil
}
