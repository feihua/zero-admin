package productattributeservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductAttributeLogic 删除商品属性
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type DeleteProductAttributeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductAttributeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductAttributeLogic {
	return &DeleteProductAttributeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductAttribute 删除商品属性
func (l *DeleteProductAttributeLogic) DeleteProductAttribute(in *pmsclient.DeleteProductAttributeReq) (*pmsclient.DeleteProductAttributeResp, error) {
	q := query.PmsProductAttribute

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除商品属性失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除商品属性失败")
	}

	return &pmsclient.DeleteProductAttributeResp{}, nil
}
