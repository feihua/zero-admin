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

// DeleteProductCategoryLogic 删除产品分类
/*
Author: LiuFeiHua
Date: 2025/05/26 10:33:54
*/
type DeleteProductCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductCategoryLogic {
	return &DeleteProductCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductCategory 删除产品分类
func (l *DeleteProductCategoryLogic) DeleteProductCategory(in *pmsclient.DeleteProductCategoryReq) (*pmsclient.DeleteProductCategoryResp, error) {
	q := query.PmsProductCategory

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除产品分类失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除产品分类失败")
	}

	return &pmsclient.DeleteProductCategoryResp{}, nil
}
