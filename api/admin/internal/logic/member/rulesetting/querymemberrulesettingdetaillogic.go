package rulesetting

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberRuleSettingDetailLogic 查询会员积分成长规则表详情
/*
Author: 刘飞华
Date: 2025/02/05 10:34:53
*/
type QueryMemberRuleSettingDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberRuleSettingDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberRuleSettingDetailLogic {
	return &QueryMemberRuleSettingDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberRuleSettingDetail 查询会员积分成长规则表详情
func (l *QueryMemberRuleSettingDetailLogic) QueryMemberRuleSettingDetail(req *types.QueryMemberRuleSettingDetailReq) (resp *types.QueryMemberRuleSettingDetailResp, err error) {

	detail, err := l.svcCtx.MemberRuleSettingService.QueryMemberRuleSettingDetail(l.ctx, &umsclient.QueryMemberRuleSettingDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询会员积分成长规则表详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryMemberRuleSettingDetailData{
		Id:                detail.Id,                //
		ContinueSignDay:   detail.ContinueSignDay,   // 连续签到天数
		ContinueSignPoint: detail.ContinueSignPoint, // 连续签到赠送数量
		ConsumePerPoint:   detail.ConsumePerPoint,   // 每消费多少元获取1个点
		LowOrderAmount:    detail.LowOrderAmount,    // 最低获取点数的订单金额
		MaxPointPerOrder:  detail.MaxPointPerOrder,  // 每笔订单最高获取点数
		RuleType:          detail.RuleType,          // 类型：0->积分规则；1->成长值规则
		Status:            detail.Status,            // 状态：0->禁用；1->启用
		CreateBy:          detail.CreateBy,          // 创建者
		CreateTime:        detail.CreateTime,        // 创建时间
		UpdateBy:          detail.UpdateBy,          // 更新者
		UpdateTime:        detail.UpdateTime,        // 更新时间
	}
	return &types.QueryMemberRuleSettingDetailResp{
		Code:    "000000",
		Message: "查询会员积分成长规则表成功",
		Data:    data,
	}, nil
}
