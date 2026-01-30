package brand

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

// QueryBrandListLogic 分页获取推荐品牌
/*
Author: LiuFeiHua
Date: 2025/6/19 10:22
*/
type QueryBrandListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryBrandListLogic {
	return &QueryBrandListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryBrandList 分页获取推荐品牌
func (l *QueryBrandListLogic) QueryBrandList(req *types.BrandListReq) (resp *types.BrandListResp, err error) {
	brandListResp, err := l.svcCtx.ProductBrandService.QueryProductBrandList(l.ctx, &pmsclient.QueryProductBrandListReq{
		PageNum:         req.Current,
		PageSize:        req.PageSize,
		RecommendStatus: 2, // 推荐状态
		IsEnabled:       1, // 是否启用
	})

	if err != nil {
		logc.Errorf(l.ctx, "分页获取推荐品牌失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []types.BrandData

	for _, detail := range brandListResp.List {
		list = append(list, types.BrandData{
			Id:                  detail.Id,                  //
			Name:                detail.Name,                // 品牌名称
			Logo:                detail.Logo,                // 品牌logo
			BigPic:              detail.BigPic,              // 专区大图
			Description:         detail.Description,         // 描述
			FirstLetter:         detail.FirstLetter,         // 首字母
			Sort:                detail.Sort,                // 排序
			RecommendStatus:     detail.RecommendStatus,     // 推荐状态
			ProductCount:        detail.ProductCount,        // 产品数量
			ProductCommentCount: detail.ProductCommentCount, // 产品评论数量
		})
	}

	return &types.BrandListResp{
		Code:    0,
		Message: "分页获取推荐品牌成功",
		Data:    list,
	}, nil
}
