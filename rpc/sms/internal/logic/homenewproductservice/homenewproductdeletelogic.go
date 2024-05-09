package homenewproductservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

// HomeNewProductDeleteLogic 首页新鲜好物
/*
Author: LiuFeiHua
Date: 2024/5/8 10:19
*/
type HomeNewProductDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeNewProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeNewProductDeleteLogic {
	return &HomeNewProductDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// HomeNewProductDelete 删除首页新鲜好物
func (l *HomeNewProductDeleteLogic) HomeNewProductDelete(in *smsclient.HomeNewProductDeleteReq) (*smsclient.HomeNewProductDeleteResp, error) {
	q := query.SmsHomeNewProduct
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &smsclient.HomeNewProductDeleteResp{}, nil
}
