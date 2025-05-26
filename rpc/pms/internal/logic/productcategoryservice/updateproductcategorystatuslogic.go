package productcategoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateProductCategoryStatusLogic 更新产品分类
/*
Author: LiuFeiHua
Date: 2025/05/26 10:33:54
*/
type UpdateProductCategoryStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductCategoryStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductCategoryStatusLogic {
	return &UpdateProductCategoryStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductCategoryStatus 更新产品分类状态
func (l *UpdateProductCategoryStatusLogic) UpdateProductCategoryStatus(in *pmsclient.UpdateProductCategoryStatusReq) (*pmsclient.UpdateProductCategoryStatusResp, error) {
	q := query.PmsProductCategory

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.IsEnabled, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "更新产品分类状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新产品分类状态失败")
	}

	return &pmsclient.UpdateProductCategoryStatusResp{}, nil
}
