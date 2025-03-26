package orderreturnapplyservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateOrderReturnApplyLogic 更新订单退货申请
/*
Author: LiuFeiHua
Date: 2024/6/12 10:16
*/
type UpdateOrderReturnApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderReturnApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderReturnApplyLogic {
	return &UpdateOrderReturnApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateOrderReturnApply 更新订单退货申请
func (l *UpdateOrderReturnApplyLogic) UpdateOrderReturnApply(in *omsclient.UpdateOrderReturnApplyReq) (*omsclient.UpdateOrderReturnApplyResp, error) {
	q := query.OmsOrderReturnApply
	returnApply, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).First()
	if err != nil {
		logc.Errorf(l.ctx, "更新订单退货申请失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新订单退货申请失败")
	}

	returnApply.Status = in.Status
	operTime := time.Now()
	// 确认退货
	if in.Status == 1 {
		returnApply.CompanyAddressID = in.CompanyAddressId
		returnApply.HandleTime = &operTime
		returnApply.HandleNote = in.HandleNote
		returnApply.HandleMan = in.HandleMan
		returnApply.ReturnAmount = in.ReturnAmount
	}
	// 完成退货
	if in.Status == 2 {
		returnApply.CompanyAddressID = in.CompanyAddressId
		returnApply.ReceiveMan = in.ReceiveMan
		returnApply.ReceiveTime = &operTime
		returnApply.ReceiveNote = in.ReceiveNote
	}
	// 拒绝退货
	if in.Status == 3 {
		returnApply.HandleTime = &operTime
		returnApply.HandleNote = in.HandleNote
		returnApply.HandleMan = in.HandleMan
	}

	_, err = q.WithContext(l.ctx).Updates(returnApply)
	if err != nil {
		logc.Errorf(l.ctx, "更新订单退货申请失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新订单退货申请失败")
	}
	return &omsclient.UpdateOrderReturnApplyResp{}, nil
}
