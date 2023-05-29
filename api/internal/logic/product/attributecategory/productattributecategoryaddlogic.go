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

type ProductAttributecategoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAttributecategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributecategoryAddLogic {
	return &ProductAttributecategoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductAttributecategoryAddLogic) ProductAttributecategoryAdd(req *types.AddProductAttributecategoryReq) (resp *types.AddProductAttributecategoryResp, err error) {
	_, err = l.svcCtx.Pms.ProductAttributeCategoryAdd(l.ctx, &pmsclient.ProductAttributeCategoryAddReq{
		Name: req.Name,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加属性分类信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加属性分类失败")
	}

	return &types.AddProductAttributecategoryResp{
		Code:    "000000",
		Message: "添加属性分类成功",
	}, nil
}
