package orderreturnreasonservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateOrderReturnReasonLogic 更新退货原因
type UpdateOrderReturnReasonLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderReturnReasonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderReturnReasonLogic {
	return &UpdateOrderReturnReasonLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateOrderReturnReason 更新退货原因
func (l *UpdateOrderReturnReasonLogic) UpdateOrderReturnReason(in *omsclient.UpdateOrderReturnReasonReq) (*omsclient.UpdateOrderReturnReasonResp, error) {
	q := query.OmsOrderReturnReason
	item := &model.OmsOrderReturnReason{
		ID:         in.Id,      //
		Name:       in.Name,    // 退货类型
		Sort:       in.Sort,    // 排序
		Status:     in.Status,  // 状态：0->不启用；1->启用
		CreateTime: time.Now(), // 创建时间
	}

	err := l.svcCtx.DB.Model(&model.OmsOrderReturnReason{}).WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新退货原因失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新退货原因失败")
	}

	return &omsclient.UpdateOrderReturnReasonResp{}, nil
}
