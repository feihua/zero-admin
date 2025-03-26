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

// DeleteHomeNewProductLogic 删除新鲜好物
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

// DeleteHomeNewProduct 删除新鲜好物
func (l *DeleteHomeNewProductLogic) DeleteHomeNewProduct(in *smsclient.DeleteHomeNewProductReq) (*smsclient.DeleteHomeNewProductResp, error) {
	q := query.SmsHomeNewProduct
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除新鲜好物失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除新鲜好物失败")
	}

	return &smsclient.DeleteHomeNewProductResp{}, nil
}
