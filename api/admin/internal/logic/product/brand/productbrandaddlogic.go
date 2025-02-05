package brand

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductBrandAddLogic 商品品牌
/*
Author: LiuFeiHua
Date: 2024/5/13 16:53
*/
type ProductBrandAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductBrandAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductBrandAddLogic {
	return ProductBrandAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductBrandAdd 添加商品品牌
func (l *ProductBrandAddLogic) ProductBrandAdd(req types.AddProductBrandReq) (*types.AddProductBrandResp, error) {
	_, err := l.svcCtx.BrandService.AddBrand(l.ctx, &pmsclient.AddBrandReq{
		Name:                req.Name,            // 品牌名称
		FirstLetter:         req.FirstLetter,     // 首字母
		Sort:                req.Sort,            // 排序
		FactoryStatus:       req.FactoryStatus,   // 是否为品牌制造商：0->不是；1->是
		ShowStatus:          req.ShowStatus,      // 品牌显示状态
		RecommendStatus:     req.RecommendStatus, // 推荐状态
		ProductCount:        0,                   // 产品数量
		ProductCommentCount: 0,                   // 产品评论数量
		Logo:                req.Logo,            // 品牌logo
		BigPic:              req.BigPic,          // 专区大图
		BrandStory:          req.BrandStory,      // 品牌故事
		CreateBy:            l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加商品品牌信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加商品品牌失败")
	}

	return &types.AddProductBrandResp{
		Code:    "000000",
		Message: "添加商品品牌成功",
	}, nil
}
