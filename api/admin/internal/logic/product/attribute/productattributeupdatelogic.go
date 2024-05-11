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

type ProductAttributeUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAttributeUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeUpdateLogic {
	return &ProductAttributeUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductAttributeUpdateLogic) ProductAttributeUpdate(req *types.UpdateProductAttributeReq) (resp *types.UpdateProductAttributeResp, err error) {
	_, err = l.svcCtx.ProductAttributeService.ProductAttributeUpdate(l.ctx, &pmsclient.ProductAttributeUpdateReq{
		ProductAttributeCategoryId: req.ProductAttributeCategoryId,
		Id:                         req.Id,
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
		logc.Errorf(l.ctx, "更新属性信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新属性失败")
	}

	return &types.UpdateProductAttributeResp{
		Code:    "000000",
		Message: "更新属性成功",
	}, nil
}
