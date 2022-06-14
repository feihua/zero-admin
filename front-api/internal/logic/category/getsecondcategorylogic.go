package category

import (
	"context"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSecondCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSecondCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetSecondCategoryLogic {
	return GetSecondCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSecondCategoryLogic) GetSecondCategory(req types.SecondCategoryReq) (resp *types.CategoryResp, err error) {
	result, _ := l.svcCtx.Pms.ProductCategorySecondList(l.ctx, &pmsclient.ProductCategorySecondListReq{
		ParentId: req.Id,
	})

	var list []types.CategoryData

	for _, item := range result.List {
		list = append(list, types.CategoryData{
			ID:        item.Id,
			Name:      item.Name,
			Keywords:  item.Keywords,
			Desc:      item.Description,
			PID:       item.ParentId,
			IconURL:   item.Icon,
			PicURL:    item.Icon,
			SortOrder: item.Sort,
		})
	}

	return &types.CategoryResp{
		Errno:  0,
		Data:   list,
		Errmsg: "",
	}, nil
}
