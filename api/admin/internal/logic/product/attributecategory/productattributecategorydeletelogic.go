package attributecategory

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

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
		logc.Errorf(l.ctx, "根据Id: %+v,删除属性分类异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除属性分类失败")
	}

	return &types.DeleteProductAttributecategoryResp{
		Code:    "000000",
		Message: "",
	}, nil
}
