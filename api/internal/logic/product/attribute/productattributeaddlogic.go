package attribute

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/api/internal/svc"
	"github.com/feihua/zero-admin/api/internal/types"

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
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加属性信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加属性失败")
	}

	return &types.AddProductAttributeResp{
		Code:    "000000",
		Message: "添加属性成功",
	}, nil
}
