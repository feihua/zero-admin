package homenewproductservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateHomeNewProductStatusLogic 更新新鲜好物状态
/*
Author: LiuFeiHua
Date: 2024/6/12 17:55
*/
type UpdateHomeNewProductStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHomeNewProductStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomeNewProductStatusLogic {
	return &UpdateHomeNewProductStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateHomeNewProductStatus 更新新鲜好物状态
func (l *UpdateHomeNewProductStatusLogic) UpdateHomeNewProductStatus(in *smsclient.UpdateHomeNewProductStatusReq) (*smsclient.UpdateHomeNewProductStatusResp, error) {
	q := query.SmsHomeNewProduct
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.RecommendStatus, in.RecommendStatus)

	if err != nil {
		logc.Errorf(l.ctx, "更新新鲜好物状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新新鲜好物状态失败")
	}

	return &smsclient.UpdateHomeNewProductStatusResp{}, nil
}
