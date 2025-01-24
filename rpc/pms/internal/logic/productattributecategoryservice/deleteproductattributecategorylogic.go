package productattributecategoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductAttributeCategoryLogic 删除产品属性分类
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:06
*/
type DeleteProductAttributeCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductAttributeCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductAttributeCategoryLogic {
	return &DeleteProductAttributeCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductAttributeCategory 删除产品属性分类
func (l *DeleteProductAttributeCategoryLogic) DeleteProductAttributeCategory(in *pmsclient.DeleteProductAttributeCategoryReq) (*pmsclient.DeleteProductAttributeCategoryResp, error) {
	q := query.PmsProductAttributeCategory

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除产品属性分类失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除产品属性分类失败")
	}

	logc.Infof(l.ctx, "删除产品属性分类成功,参数：%+v", in)
	return &pmsclient.DeleteProductAttributeCategoryResp{}, nil
}
