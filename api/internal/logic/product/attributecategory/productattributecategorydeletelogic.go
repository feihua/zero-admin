package attributecategory

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAttributecategoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAttributecategoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributecategoryDeleteLogic {
	return &ProductAttributecategoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductAttributecategoryDeleteLogic) ProductAttributecategoryDelete(req *types.DeleteProductAttributecategoryReq) (resp *types.DeleteProductAttributecategoryResp, err error) {
	_, err = l.svcCtx.ProductAttributeCategoryService.ProductAttributeCategoryDelete(l.ctx, &pmsclient.ProductAttributeCategoryDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除属性分类异常:%s", req.Ids, err.Error())
		return nil, errorx.NewDefaultError("删除属性分类失败")
	}

	return &types.DeleteProductAttributecategoryResp{
		Code:    "000000",
		Message: "",
	}, nil
}
