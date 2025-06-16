package productattributegroupservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductAttributeGroupLogic 删除商品属性分组
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type DeleteProductAttributeGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductAttributeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductAttributeGroupLogic {
	return &DeleteProductAttributeGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductAttributeGroup 删除商品属性分组
func (l *DeleteProductAttributeGroupLogic) DeleteProductAttributeGroup(in *pmsclient.DeleteProductAttributeGroupReq) (*pmsclient.DeleteProductAttributeGroupResp, error) {
	q := query.PmsProductAttributeGroup

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除商品属性分组失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除商品属性分组失败")
	}

	return &pmsclient.DeleteProductAttributeGroupResp{}, nil
}
