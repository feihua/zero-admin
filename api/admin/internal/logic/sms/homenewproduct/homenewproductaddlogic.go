package logic

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

func (l *HomeNewProductAddLogic) HomeNewProductAdd(req types.AddHomeNewProductReq) (*types.AddHomeNewProductResp, error) {
	brandListResp, _ := l.svcCtx.ProductService.ProductListByIds(l.ctx, &pmsclient.ProductByIdsReq{Ids: req.ProductIds})

	var list []*smsclient.HomeNewProductAddData

	for _, item := range brandListResp.List {
		list = append(list, &smsclient.HomeNewProductAddData{
			ProductId:       item.Id,
			ProductName:     item.Name,
			RecommendStatus: 1,
			Sort:            item.Sort,
		})
	}

	_, err := l.svcCtx.HomeNewProductService.HomeNewProductAdd(l.ctx, &smsclient.HomeNewProductAddReq{
		NewProductAddData: list,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加新鲜好物信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加新鲜好物表失败")
	}

	return &types.AddHomeNewProductResp{
		Code:    "000000",
		Message: "添加新鲜好物表成功",
	}, nil
}
