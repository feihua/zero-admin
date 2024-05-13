package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeNewProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeNewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeNewProductListLogic {
	return HomeNewProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeNewProductListLogic) HomeNewProductList(req types.ListHomeNewProductReq) (*types.ListHomeNewProductResp, error) {
	resp, err := l.svcCtx.HomeNewProductService.HomeNewProductList(l.ctx, &smsclient.HomeNewProductListReq{
		Current:         req.Current,
		PageSize:        req.PageSize,
		ProductName:     req.ProductName,
		RecommendStatus: req.RecommendStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询新鲜好物表列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询新鲜好物表失败")
	}

	var list []*types.ListHomeNewProductData

	for _, item := range resp.List {
		list = append(list, &types.ListHomeNewProductData{
			Id:              item.Id,
			ProductId:       item.ProductId,
			ProductName:     item.ProductName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
		})
	}

	return &types.ListHomeNewProductResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询新鲜好物表成功",
	}, nil
}
