package order_return

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateOrderReturnLogic 更新退货/售后主
/*
Author: LiuFeiHua
Date: 2025/07/01 10:06:44
*/
type UpdateOrderReturnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderReturnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderReturnLogic {
	return &UpdateOrderReturnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateOrderReturn 更新退货/售后主
func (l *UpdateOrderReturnLogic) UpdateOrderReturn(req *types.UpdateOrderReturnReq) (resp *types.BaseResp, err error) {
	userName := l.ctx.Value("userName").(string)
	_, err = l.svcCtx.OrderReturnService.UpdateOrderReturn(l.ctx, &omsclient.OrderReturnReq{
		Id:               req.Id,                    // 主键ID
		Status:           req.Status,                // 退货状态（0待审核 1审核通过 2已收货 3已退款 4已拒绝 5已关闭）
		RefundAmount:     float32(req.RefundAmount), // 退款金额
		HandleNote:       req.HandleNote,            // 处理备注
		HandleMan:        userName,                  // 处理人员
		ReceiveNote:      req.ReceiveNote,           // 收货备注
		ReceiveMan:       userName,                  // 收货人
		CompanyAddressID: req.CompanyAddressId,      // 公司地址ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新退货/售后主失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新退货/售后主成功",
	}, nil
}
