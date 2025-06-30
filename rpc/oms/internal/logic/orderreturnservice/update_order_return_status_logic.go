package orderreturnservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

// UpdateOrderReturnStatusLogic 更新退货/售后
/*
Author: LiuFeiHua
Date: 2025/06/30 15:49:28
*/
type UpdateOrderReturnStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderReturnStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderReturnStatusLogic {
	return &UpdateOrderReturnStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateOrderReturnStatus 更新退货/售后状态
func (l *UpdateOrderReturnStatusLogic) UpdateOrderReturnStatus(in *omsclient.UpdateOrderReturnStatusReq) (*omsclient.OrderReturnResp, error) {
	re := query.OmsOrderReturn
	item, err := re.WithContext(l.ctx).Where(re.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "退货/售后不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("退货/售后不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询退货/售后异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询退货/售后异常")
	}

	item.Status = in.Status
	operTime := time.Now()
	// 确认退货
	if in.Status == 1 {
		item.CompanyAddress = in.CompanyAddress
		item.HandleTime = &operTime
		item.HandleNote = in.HandleNote
		item.HandleMan = in.HandleMan
		item.RefundAmount = float64(in.RefundAmount)
	}
	// 完成退货
	if in.Status == 2 {
		item.CompanyAddress = in.CompanyAddress
		item.ReceiveMan = in.ReceiveMan
		item.ReceiveTime = &operTime
		item.ReceiveNote = in.ReceiveNote
	}
	// 拒绝退货
	if in.Status == 3 {
		item.HandleTime = &operTime
		item.HandleNote = in.HandleNote
		item.HandleMan = in.HandleMan
	}

	_, err = re.WithContext(l.ctx).Updates(item)
	if err != nil {
		logc.Errorf(l.ctx, "更新订单退货申请失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新订单退货申请失败")
	}

	return &omsclient.OrderReturnResp{}, nil
}
