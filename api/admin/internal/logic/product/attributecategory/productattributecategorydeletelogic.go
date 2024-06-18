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

type ProductAttributeCategoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAttributeCategoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeCategoryDeleteLogic {
	return &ProductAttributeCategoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductAttributeCategoryDeleteLogic) ProductAttributeCategoryDelete(req *types.DeleteProductAttributeCategoryReq) (resp *types.DeleteProductAttributeCategoryResp, err error) {
	_, err = l.svcCtx.ProductAttributeCategoryService.DeleteProductAttributeCategory(l.ctx, &pmsclient.DeleteProductAttributeCategoryReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除属性分类异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除属性分类失败")
	}

	return &types.DeleteProductAttributeCategoryResp{
		Code:    "000000",
		Message: "",
	}, nil
}
