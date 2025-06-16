package productattributevalueservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateProductAttributeValueStatusLogic 更新商品属性值
/*
Author: LiuFeiHua
Date: 2025/06/17 16:14:47
*/
type UpdateProductAttributeValueStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductAttributeValueStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductAttributeValueStatusLogic {
	return &UpdateProductAttributeValueStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductAttributeValueStatus 更新商品属性值状态
func (l *UpdateProductAttributeValueStatusLogic) UpdateProductAttributeValueStatus(in *pmsclient.UpdateProductAttributeValueStatusReq) (*pmsclient.UpdateProductAttributeValueStatusResp, error) {
	q := query.PmsProductAttributeValue

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "更新商品属性值状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新商品属性值状态失败")
	}

	return &pmsclient.UpdateProductAttributeValueStatusResp{}, nil
}
