package category

import (
	"context"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductCateListByIdLogic 根据parentId查询分类
/*
Author: LiuFeiHua
Date: 2024/5/16 14:47
*/
type QueryProductCateListByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryProductCateListByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductCateListByIdLogic {
	return &QueryProductCateListByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryProductCateListById 根据parentId查询分类
func (l *QueryProductCateListByIdLogic) QueryProductCateListById(req *types.QueryProductCateListByIdReq) (resp *types.QueryProductCateListResp, err error) {
	categoryListResp, err := l.svcCtx.ProductCategoryService.QueryProductCategoryList(l.ctx, &pmsclient.QueryProductCategoryListReq{
		PageNum:    0,
		PageSize:   100,
		ParentId:   req.ParentId,
		ShowStatus: 1, // 显示状态：0->不显示；1->显示
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询商品分类失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	list := make([]types.ProductCateListData, 0)
	for _, item := range categoryListResp.List {
		list = append(list, types.ProductCateListData{
			Id:       item.Id,
			Name:     item.Name,
			ImageUrl: item.Icon,
		})
	}

	return &types.QueryProductCateListResp{
		Code:    0,
		Message: "操作成功",
		Data:    list,
	}, nil
}
