package homenewproductservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteHomeNewProductLogic 删除新鲜好物表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:54
*/
type DeleteHomeNewProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteHomeNewProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHomeNewProductLogic {
	return &DeleteHomeNewProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteHomeNewProduct 删除新鲜好物表
func (l *DeleteHomeNewProductLogic) DeleteHomeNewProduct(in *smsclient.DeleteHomeNewProductReq) (*smsclient.DeleteHomeNewProductResp, error) {
	q := query.SmsHomeNewProduct
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &smsclient.DeleteHomeNewProductResp{}, nil
}
