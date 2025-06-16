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

// UpdateProductAttributeStatusLogic 更新商品属性
/*
Author: LiuFeiHua
Date: 2025/06/17 16:14:47
*/
type UpdateProductAttributeStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductAttributeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductAttributeStatusLogic {
	return &UpdateProductAttributeStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductAttributeStatus 更新商品属性状态
func (l *UpdateProductAttributeStatusLogic) UpdateProductAttributeStatus(in *pmsclient.UpdateProductAttributeStatusReq) (*pmsclient.UpdateProductAttributeStatusResp, error) {
	q := query.PmsProductAttribute

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "更新商品属性状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新商品属性状态失败")
	}

	return &pmsclient.UpdateProductAttributeStatusResp{}, nil
}
