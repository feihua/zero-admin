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
	"gorm.io/gorm"
	"time"
)

// UpdateOrderReturnReasonLogic 更新退货原因
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:44
*/
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
	reason := query.OmsOrderReturnReason
	q := reason.WithContext(l.ctx)

	// 1.根据退货原因id查询退货原因是否已存在
	detail, err := q.Where(reason.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "退货原因不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("退货原因不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询退货原因异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询退货原因异常")
	}

	count, _ := q.Where(reason.ID.Neq(in.Id), reason.Name.Eq(in.Name)).Count()

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("退货原因：%s,已存在", in.Name))
	}

	now := time.Now()
	item := &model.OmsOrderReturnReason{
		ID:         in.Id,             // 主键ID
		Name:       in.Name,           // 退货类型
		Sort:       in.Sort,           // 排序
		Status:     in.Status,         // 状态：0->不启用；1->启用
		CreateBy:   detail.CreateBy,   // 创建者
		CreateTime: detail.CreateTime, // 创建时间
		UpdateBy:   &in.UpdateBy,      // 更新者
		UpdateTime: &now,              // 更新时间
	}

	// 2.退货原因存在时,则直接更新退货原因
	err = l.svcCtx.DB.Model(&model.OmsOrderReturnReason{}).WithContext(l.ctx).Where(reason.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新退货原因失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新退货原因失败")
	}

	return &omsclient.UpdateOrderReturnReasonResp{}, nil
}
