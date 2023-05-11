package attributecategory

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAttributecategoryUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAttributecategoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributecategoryUpdateLogic {
	return &ProductAttributecategoryUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductAttributecategoryUpdateLogic) ProductAttributecategoryUpdate(req *types.UpdateProductAttributecategoryReq) (resp *types.UpdateProductAttributecategoryResp, err error) {
	_, err = l.svcCtx.Pms.ProductAttributeCategoryUpdate(l.ctx, &pmsclient.ProductAttributeCategoryUpdateReq{
		Id:             req.Id,
		Name:           req.Name,
		AttributeCount: req.AttributecategoryCount,
		ParamCount:     req.ParamCount,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新属性分类信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新属性分类失败")
	}

	return &types.UpdateProductAttributecategoryResp{
		Code:    "000000",
		Message: "更新属性分类成功",
	}, nil
}
