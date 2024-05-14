package homenewproductservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateNewProductStatusLogic 首页新品
/*
Author: LiuFeiHua
Date: 2024/5/14 9:31
*/
type UpdateNewProductStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNewProductStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNewProductStatusLogic {
	return &UpdateNewProductStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateNewProductStatus 批量修改首页新品状态
func (l *UpdateNewProductStatusLogic) UpdateNewProductStatus(in *smsclient.UpdateNewProductStatusReq) (*smsclient.UpdateNewProductStatusResp, error) {
	q := query.SmsHomeNewProduct
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.RecommendStatus, in.RecommendStatus)

	if err != nil {
		return nil, err
	}

	return &smsclient.UpdateNewProductStatusResp{}, nil
}
