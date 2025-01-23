package orderreturnreasonservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddOrderReturnReasonLogic 添加退货原因表
/*
Author: LiuFeiHua
Date: 2024/6/12 10:18
*/
type AddOrderReturnReasonLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderReturnReasonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderReturnReasonLogic {
	return &AddOrderReturnReasonLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddOrderReturnReason 添加退货原因表
func (l *AddOrderReturnReasonLogic) AddOrderReturnReason(in *omsclient.AddOrderReturnReasonReq) (*omsclient.AddOrderReturnReasonResp, error) {
	err := query.OmsOrderReturnReason.WithContext(l.ctx).Create(&model.OmsOrderReturnReason{
		Name:   in.Name,   // 退货类型
		Sort:   in.Sort,   // 排序
		Status: in.Status, // 状态：0->不启用；1->启用
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.AddOrderReturnReasonResp{}, nil
}
