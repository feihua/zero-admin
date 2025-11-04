package orderreturnservicelogic

import (
	"context"
	"errors"
	"time"

	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// UpdateOrderReturnLogic 更新退货/售后
/*
Author: LiuFeiHua
Date: 2025/06/30 15:49:28
*/
type UpdateOrderReturnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderReturnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderReturnLogic {
	return &UpdateOrderReturnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateOrderReturn 更新退货/售后
func (l *UpdateOrderReturnLogic) UpdateOrderReturn(in *omsclient.OrderReturnReq) (*omsclient.OrderReturnResp, error) {
	q := query.OmsOrderReturn.WithContext(l.ctx)

	// 1.根据退货/售后id查询退货/售后是否已存在
	returnApply, err := q.Where(query.OmsOrderReturn.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "退货/售后不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("退货/售后不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询退货/售后异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询退货/售后异常")
	}

	returnApply.Status = in.Status
	operTime := time.Now()
	// 确认退货
	if in.Status == 1 {
		// returnApply.CompanyAddressID = in.CompanyAddressID
		returnApply.HandleTime = &operTime
		returnApply.HandleNote = in.HandleNote
		returnApply.HandleMan = in.HandleMan
		returnApply.RefundAmount = float64(in.RefundAmount)
	}
	// 完成退货
	if in.Status == 2 {
		// returnApply.CompanyAddressID = in.CompanyAddressId
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

	// 2.退货/售后存在时,则直接更新退货/售后
	_, err = q.Where(query.OmsOrderReturn.ID.Eq(in.Id)).Updates(returnApply)

	if err != nil {
		logc.Errorf(l.ctx, "更新退货/售后失败,参数:%+v,异常:%s", returnApply, err.Error())
		return nil, errors.New("更新退货/售后失败")
	}

	return &omsclient.OrderReturnResp{}, nil
}
