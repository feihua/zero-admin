package homenewproduct

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeNewProductAddLogic 首页新品
/*
Author: LiuFeiHua
Date: 2024/5/14 9:17
*/
type HomeNewProductAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeNewProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeNewProductAddLogic {
	return HomeNewProductAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// HomeNewProductAdd 添加首页新品
// 1.根据productIds查询商品(pms-rpc)
// 2.添加首页新品推荐记录(sms-rpc)
// 3.修改商品的推荐状态(pms-rpc)
func (l *HomeNewProductAddLogic) HomeNewProductAdd(req types.AddHomeNewProductReq) (*types.AddHomeNewProductResp, error) {
	// 1.根据productIds查询商品(pms-rpc)
	brandListResp, _ := l.svcCtx.ProductService.QueryProductListByIds(l.ctx, &pmsclient.QueryProductByIdsReq{Ids: req.ProductIds})

	var list []*smsclient.HomeNewProductAddData

	for _, item := range brandListResp.List {
		list = append(list, &smsclient.HomeNewProductAddData{
			ProductId:       item.Id,   // 商品id
			ProductName:     item.Name, // 商品名称
			RecommendStatus: 1,         // 推荐状态：0->不推荐;1->推荐
			Sort:            item.Sort, // 排序
		})
	}

	// 2.添加首页新品推荐记录(sms-rpc)
	_, err := l.svcCtx.HomeNewProductService.AddHomeNewProduct(l.ctx, &smsclient.AddHomeNewProductReq{
		NewProductAddData: list,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加新鲜好物信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	// 3.修改商品的推荐状态(pms-rpc)
	_, err = l.svcCtx.ProductService.UpdateNewStatus(l.ctx, &pmsclient.UpdateProductStatusReq{
		Ids:    req.ProductIds,
		Status: 1,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Ids: %+v,修改商品的推荐状态异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddHomeNewProductResp{
		Code:    "000000",
		Message: "添加新鲜好物表成功",
	}, nil
}
