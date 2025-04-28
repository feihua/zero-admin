package product_brand

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductBrandUpdateLogic 商品品牌
/*
Author: LiuFeiHua
Date: 2024/5/13 16:55
*/
type ProductBrandUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductBrandUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductBrandUpdateLogic {
	return ProductBrandUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductBrandUpdate 更新商品品牌
func (l *ProductBrandUpdateLogic) ProductBrandUpdate(req *types.UpdateProductBrandReq) (*types.BaseResp, error) {
	_, err := l.svcCtx.BrandService.UpdateBrand(l.ctx, &pmsclient.UpdateBrandReq{
		Id:              req.Id,
		Name:            req.Name,            // 品牌名称
		FirstLetter:     req.FirstLetter,     // 首字母
		Sort:            req.Sort,            // 排序
		FactoryStatus:   req.FactoryStatus,   // 是否为品牌制造商：0->不是；1->是
		ShowStatus:      req.ShowStatus,      // 品牌显示状态
		RecommendStatus: req.RecommendStatus, // 推荐状态
		Logo:            req.Logo,            // 品牌logo
		BigPic:          req.BigPic,          // 专区大图
		BrandStory:      req.BrandStory,      // 品牌故事
		UpdateBy:        l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新商品品牌信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	if req.RecommendStatus == 0 {
		_, err = l.svcCtx.HomeBrandService.DeleteHomeBrand(l.ctx, &smsclient.DeleteHomeBrandReq{
			Ids: []int64{req.Id},
		})

		if err != nil {
			logc.Errorf(l.ctx, "根据Id: %+v,删除首页品牌异常:%s", req, err.Error())
			s, _ := status.FromError(err)
			return nil, errorx.NewDefaultError(s.Message())
		}
	}

	if req.RecommendStatus == 1 {
		var list []*smsclient.HomeBrandAddData
		list = append(list, &smsclient.HomeBrandAddData{
			BrandId:         req.Id,              // 商品品牌id
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
