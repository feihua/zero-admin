package member_consume_setting

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddMemberConsumeSettingLogic 添加积分消费设置
/*
Author: LiuFeiHua
Date: 2025/05/23 11:32:02
*/
type AddMemberConsumeSettingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMemberConsumeSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberConsumeSettingLogic {
	return &AddMemberConsumeSettingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddMemberConsumeSetting 添加积分消费设置
func (l *AddMemberConsumeSettingLogic) AddMemberConsumeSetting(req *types.AddMemberConsumeSettingReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MemberConsumeSettingService.AddMemberConsumeSetting(l.ctx, &umsclient.AddMemberConsumeSettingReq{
		DeductionPerAmount: req.DeductionPerAmount, // 每一元需要抵扣的积分数量
		MaxPercentPerOrder: req.MaxPercentPerOrder, // 每笔订单最高抵用百分比
		UseUnit:            req.UseUnit,            // 每次使用积分最小单位100
		CouponStatus:       req.CouponStatus,       // 是否可以和优惠券同用；0->不可以；1->可以
		Status:             req.Status,             // 状态：0->禁用；1->启用
		CreateBy:           userId,                 // 创建人ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加积分消费设置失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "添加积分消费设置成功",
	}, nil
}
