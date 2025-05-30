package home_recommend_product

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

// HomeRecommendProductAddLogic 人气推荐商品
/*
Author: LiuFeiHua
Date: 2024/5/14 9:40
*/
type HomeRecommendProductAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeRecommendProductAddLogic {
	return HomeRecommendProductAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// HomeRecommendProductAdd 添加人气推荐商品
// 1.根据productIds查询商品(pms-rpc)
// 2.添加首页人气推荐记录(sms-rpc)
// 3.修改商品的推荐状态(pms-rpc)
func (l *HomeRecommendProductAddLogic) HomeRecommendProductAdd(req *types.AddHomeRecommendProductReq) (*types.BaseResp, error) {
	// 1.根据productIds查询商品(pms-rpc)
	brandListResp, _ := l.svcCtx.ProductService.QueryProductListByIds(l.ctx, &pmsclient.QueryProductByIdsReq{Ids: req.ProductIds})

	var list []*smsclient.HomeRecommendProductAddData

	for _, item := range brandListResp.List {
		list = append(list, &smsclient.HomeRecommendProductAddData{
			ProductId:       item.Id,   // 商品id
			ProductName:     item.Name, // 商品名称
			RecommendStatus: 1,         // 推荐状态：0->不推荐;1->推荐
			Sort:            int32(item.Id),
		})
	}

	// 2.添加首页人气推荐记录(sms-rpc)
	_, err := l.svcCtx.HomeRecommendProductService.AddHomeRecommendProduct(l.ctx, &smsclient.AddHomeRecommendProductReq{
		RecommendProductAddData: list,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加人气推荐商品信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	// 3.修改商品的推荐状态(pms-rpc)
	_, err = l.svcCtx.ProductService.UpdateRecommendStatus(l.ctx, &pmsclient.UpdateProductStatusReq{
		Ids:    req.ProductIds,
		Status: 1,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Ids: %+v,修改商品的推荐状态异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}
	return res.Success()
}
