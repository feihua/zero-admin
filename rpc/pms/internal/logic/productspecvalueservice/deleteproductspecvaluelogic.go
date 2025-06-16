package productspecvalueservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductSpecValueLogic 删除商品规格值
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type DeleteProductSpecValueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductSpecValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductSpecValueLogic {
	return &DeleteProductSpecValueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductSpecValue 删除商品规格值
func (l *DeleteProductSpecValueLogic) DeleteProductSpecValue(in *pmsclient.DeleteProductSpecValueReq) (*pmsclient.DeleteProductSpecValueResp, error) {
	q := query.PmsProductSpecValue

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除商品规格值失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除商品规格值失败")
	}

	return &pmsclient.DeleteProductSpecValueResp{}, nil
}
