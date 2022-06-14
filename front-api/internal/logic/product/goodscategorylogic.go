package product

import (
	"context"
	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"
	"zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) GoodsCategoryLogic {
	return GoodsCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsCategoryLogic) GoodsCategory(req types.GoodsCategoryReq) (resp *types.GoodsCategoryResp, err error) {
	result, _ := l.svcCtx.Pms.ProductListByCategoryId(l.ctx, &pmsclient.ProductListByCategoryIdReq{
		ProductCategoryId: req.Id,
	})

	var list []types.GoodsCategoryList

	for _, item := range result.List {
		list = append(list, types.GoodsCategoryList{
			ID:             item.Id,
			GoodsID:        0,
			Specifications: nil,
			Price:          0,
			Number:         0,
			URL:            "",
			AddTime:        "",
			UpdateTime:     "",
			Deleted:        false,
		})
	}

	return &types.GoodsCategoryResp{
		Errno: 0,
		Data: types.GoodsCategoryData{
			Total: 0,
			Pages: 0,
			Limit: 0,
			Page:  0,
			List:  list,
		},
		Errmsg: "",
	}, nil
}
