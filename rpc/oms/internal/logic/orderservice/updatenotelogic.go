package orderservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateNoteLogic 备注订单
/*
Author: LiuFeiHua
Date: 2024/5/14 17:34
*/
type UpdateNoteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNoteLogic {
	return &UpdateNoteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateNote 备注订单
func (l *UpdateNoteLogic) UpdateNote(in *omsclient.UpdateNoteReq) (*omsclient.UpdateNoteResp, error) {
	//1.备注订单
	q := query.OmsOrder
	order, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.OrderId)).First()
	if err != nil {
		return nil, err
	}

	order.Note = in.Note
	order.ModifyTime = time.Now()

	_, err = q.WithContext(l.ctx).Updates(order)

	if err != nil {
		return nil, err
	}

	//2.添加操作记录
	var note = "修改备注信息" + in.Note
	err = query.OmsOrderOperateHistory.WithContext(l.ctx).Create(&model.OmsOrderOperateHistory{
		OrderID:     in.OrderId,
		OperateMan:  in.OperateMan,
		CreateTime:  time.Now(),
		OrderStatus: in.Status,
		Note:        &note,
	})
	if err != nil {
		return nil, err
	}

	return &omsclient.UpdateNoteResp{}, nil
}
