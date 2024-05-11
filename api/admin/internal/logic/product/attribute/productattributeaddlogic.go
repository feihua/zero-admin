package attribute

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAttributeAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAttributeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeAddLogic {
	return &ProductAttributeAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductAttributeAddLogic) ProductAttributeAdd(req *types.AddProductAttributeReq) (resp *types.AddProductAttributeResp, err error) {
	_, err = l.svcCtx.ProductAttributeService.ProductAttributeAdd(l.ctx, &pmsclient.ProductAttributeAddReq{
		ProductAttributeCategoryId: req.ProductAttributeCategoryId,
		Name:                       req.Name,
		SelectType:                 req.SelectType,
		InputType:                  req.InputType,
		InputList:                  req.InputList,
		Sort:                       req.Sort,
		FilterType:                 req.FilterType,
		SearchType:                 req.SearchType,
		RelatedStatus:              req.RelatedStatus,
		HandAddStatus:              req.HandAddStatus,
		Type:                       req.Type,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加属性信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加属性失败")
	}

	return &types.AddProductAttributeResp{
		Code:    "000000",
		Message: "添加属性成功",
	}, nil
}
