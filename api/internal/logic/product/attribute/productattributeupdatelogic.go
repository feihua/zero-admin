package attribute

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

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
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新属性信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新属性失败")
	}

	return &types.UpdateProductAttributeResp{
		Code:    "000000",
		Message: "更新属性成功",
	}, nil
}
