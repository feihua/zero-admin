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

// UpdateProductAttributeGroupStatusLogic 更新商品属性分组
/*
Author: LiuFeiHua
Date: 2025/06/17 16:14:47
*/
type UpdateProductAttributeGroupStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductAttributeGroupStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductAttributeGroupStatusLogic {
	return &UpdateProductAttributeGroupStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductAttributeGroupStatus 更新商品属性分组状态
func (l *UpdateProductAttributeGroupStatusLogic) UpdateProductAttributeGroupStatus(in *pmsclient.UpdateProductAttributeGroupStatusReq) (*pmsclient.UpdateProductAttributeGroupStatusResp, error) {
	q := query.PmsProductAttributeGroup

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "更新商品属性分组状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新商品属性分组状态失败")
	}

	return &pmsclient.UpdateProductAttributeGroupStatusResp{}, nil
}
