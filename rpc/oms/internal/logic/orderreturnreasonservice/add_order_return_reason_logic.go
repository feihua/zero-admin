package orderreturnreasonservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddOrderReturnReasonLogic 添加退货原因
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:44
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

// AddOrderReturnReason 添加退货原因
func (l *AddOrderReturnReasonLogic) AddOrderReturnReason(in *omsclient.AddOrderReturnReasonReq) (*omsclient.AddOrderReturnReasonResp, error) {
	q := query.OmsOrderReturnReason

	count, _ := q.WithContext(l.ctx).Where(q.Name.Eq(in.Name)).Count()

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("退货原因：%s,已存在", in.Name))
	}

	item := &model.OmsOrderReturnReason{
		Name:     in.Name,     // 退货类型
		Sort:     in.Sort,     // 排序
		Status:   in.Status,   // 状态：0->不启用；1->启用
		CreateBy: in.CreateBy, // 创建人ID
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加退货原因失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加退货原因失败")
	}

	return &omsclient.AddOrderReturnReasonResp{}, nil
}
