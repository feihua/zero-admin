package homerecommendproduct

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

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
func (l *HomeRecommendProductAddLogic) HomeRecommendProductAdd(req types.AddHomeRecommendProductReq) (*types.AddHomeRecommendProductResp, error) {
	brandListResp, _ := l.svcCtx.ProductService.ProductListByIds(l.ctx, &pmsclient.ProductByIdsReq{Ids: req.ProductIds})

	var list []*smsclient.HomeRecommendProductAddData

	for _, item := range brandListResp.List {
		list = append(list, &smsclient.HomeRecommendProductAddData{
			ProductId:       item.Id,
			ProductName:     item.Name,
			RecommendStatus: 1,
			Sort:            item.Sort,
		})
	}

	_, err := l.svcCtx.HomeRecommendProductService.HomeRecommendProductAdd(l.ctx, &smsclient.HomeRecommendProductAddReq{
		RecommendProductAddData: list,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加人气推荐商品信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加人气推荐商品失败")
	}

	return &types.AddHomeRecommendProductResp{
		Code:    "000000",
		Message: "添加人气推荐商品成功",
	}, nil
}
