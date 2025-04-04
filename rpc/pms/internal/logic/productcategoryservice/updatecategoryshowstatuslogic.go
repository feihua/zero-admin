package productcategoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCategoryShowStatusLogic 商品类别
/*
Author: LiuFeiHua
Date: 2024/5/15 11:25
*/
type UpdateCategoryShowStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCategoryShowStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCategoryShowStatusLogic {
	return &UpdateCategoryShowStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCategoryShowStatus 更新商品分类显示状态
func (l *UpdateCategoryShowStatusLogic) UpdateCategoryShowStatus(in *pmsclient.UpdateProductCategoryStatusReq) (*pmsclient.UpdateProductCategoryStatusResp, error) {
	q := query.PmsProductCategory
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.ShowStatus, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "更新商品分类显示状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新商品分类显示状态失败")
	}

	return &pmsclient.UpdateProductCategoryStatusResp{}, nil
}
