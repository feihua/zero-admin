package productspecservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductSpecLogic 删除商品规格
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type DeleteProductSpecLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductSpecLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductSpecLogic {
	return &DeleteProductSpecLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductSpec 删除商品规格
func (l *DeleteProductSpecLogic) DeleteProductSpec(in *pmsclient.DeleteProductSpecReq) (*pmsclient.DeleteProductSpecResp, error) {
	q := query.PmsProductSpec

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除商品规格失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除商品规格失败")
	}

	return &pmsclient.DeleteProductSpecResp{}, nil
}
