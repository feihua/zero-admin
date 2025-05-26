package product_brand

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddProductBrandLogic 添加商品品牌
/*
Author: LiuFeiHua
Date: 2025/05/26 10:33:54
*/
type AddProductBrandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProductBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductBrandLogic {
	return &AddProductBrandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddProductBrand 添加商品品牌
func (l *AddProductBrandLogic) AddProductBrand(req *types.AddProductBrandReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	result, err := l.svcCtx.ProductBrandService.AddProductBrand(l.ctx, &pmsclient.AddProductBrandReq{
		Name:            req.Name,            // 品牌名称
		Logo:            req.Logo,            // 品牌logo
		BigPic:          req.BigPic,          // 专区大图
		Description:     req.Description,     // 描述
		FirstLetter:     req.FirstLetter,     // 首字母
		Sort:            req.Sort,            // 排序
		RecommendStatus: req.RecommendStatus, // 推荐状态
		IsEnabled:       req.IsEnabled,       // 是否启用
		CreateBy:        userId,              // 创建人ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加商品品牌失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}
	if req.RecommendStatus == 1 {
		var list []*smsclient.HomeBrandAddData
		list = append(list, &smsclient.HomeBrandAddData{
			BrandId:         result.BrandId,      // 商品品牌id
			BrandName:       req.Name,            // 商品品牌名称
			RecommendStatus: req.RecommendStatus, // 推荐状态：0->不推荐;1->推荐
			Sort:            req.Sort,            // 排序
		})

		_, err = l.svcCtx.HomeBrandService.AddHomeBrand(l.ctx, &smsclient.AddHomeBrandReq{
			BrandAddData: list,
		})

		if err != nil {
			logc.Errorf(l.ctx, "添加首页品牌信息失败,参数：%+v,响应：%s", req, err.Error())
			s, _ := status.FromError(err)
			return nil, errorx.NewDefaultError(s.Message())
		}
	}

	return res.Success()
}
