package home_brand

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryHomeBrandListLogic 首页品牌信息
/*
Author: LiuFeiHua
Date: 2024/5/13 15:53
*/
type QueryHomeBrandListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryHomeBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryHomeBrandListLogic {
	return QueryHomeBrandListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryHomeBrandList 查询首页品牌信息
func (l *QueryHomeBrandListLogic) QueryHomeBrandList(req *types.QueryHomeBrandListReq) (*types.QueryHomeBrandListResp, error) {
	result, err := l.svcCtx.ProductBrandService.QueryProductBrandList(l.ctx, &pmsclient.QueryProductBrandListReq{
		PageNum:         int32(req.Current),
		PageSize:        int32(req.PageSize),
		Name:            req.BrandName,       // 品牌名称
		RecommendStatus: req.RecommendStatus, // 推荐状态
		IsEnabled:       1,                   // 是否启用
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字商品品牌列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryHomeBrandListData

	for _, detail := range result.List {
		list = append(list, &types.QueryHomeBrandListData{
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
			IsEnabled:           detail.IsEnabled,           // 是否启用
			CreateBy:            detail.CreateBy,            // 创建人ID
			CreateTime:          detail.CreateTime,          // 创建时间
			UpdateBy:            detail.UpdateBy,            // 更新人ID
			UpdateTime:          detail.UpdateTime,          // 更新时间

		})
	}

	return &types.QueryHomeBrandListResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询首页品牌成功",
	}, nil
}
