package member_consume_setting

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

// QueryMemberConsumeSettingDetailLogic 查询积分消费设置详情
/*
Author: LiuFeiHua
Date: 2025/05/23 11:32:02
*/
type QueryMemberConsumeSettingDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberConsumeSettingDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberConsumeSettingDetailLogic {
	return &QueryMemberConsumeSettingDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberConsumeSettingDetail 查询积分消费设置详情
func (l *QueryMemberConsumeSettingDetailLogic) QueryMemberConsumeSettingDetail(req *types.QueryMemberConsumeSettingDetailReq) (resp *types.QueryMemberConsumeSettingDetailResp, err error) {

	detail, err := l.svcCtx.MemberConsumeSettingService.QueryMemberConsumeSettingDetail(l.ctx, &umsclient.QueryMemberConsumeSettingDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询积分消费设置详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryMemberConsumeSettingDetailData{
		Id:                 detail.Id,                 //
		DeductionPerAmount: detail.DeductionPerAmount, // 每一元需要抵扣的积分数量
		MaxPercentPerOrder: detail.MaxPercentPerOrder, // 每笔订单最高抵用百分比
		UseUnit:            detail.UseUnit,            // 每次使用积分最小单位100
		CouponStatus:       detail.CouponStatus,       // 是否可以和优惠券同用；0->不可以；1->可以
		Status:             detail.Status,             // 状态：0->禁用；1->启用
		CreateBy:           detail.CreateBy,           // 创建人ID
		CreateTime:         detail.CreateTime,         // 创建时间
		UpdateBy:           detail.UpdateBy,           // 更新人ID
		UpdateTime:         detail.UpdateTime,         // 更新时间

	}
	return &types.QueryMemberConsumeSettingDetailResp{
		Code:    "000000",
		Message: "查询积分消费设置成功",
		Data:    data,
	}, nil
}
