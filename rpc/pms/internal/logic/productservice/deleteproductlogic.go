package productservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductLogic 删除商品
/*
Author: LiuFeiHua
Date: 2025/3/26 16:36
*/
type DeleteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProduct 删除商品
func (l *DeleteProductLogic) DeleteProduct(in *pmsclient.DeleteProductReq) (*pmsclient.DeleteProductResp, error) {
	q := query.PmsProduct
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除商品失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除商品失败")
	}

	return &pmsclient.DeleteProductResp{}, nil
}
