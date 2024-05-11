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

type HomeNewProductUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeNewProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeNewProductUpdateLogic {
	return HomeNewProductUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeNewProductUpdateLogic) HomeNewProductUpdate(req types.UpdateHomeNewProductReq) (*types.UpdateHomeNewProductResp, error) {
	_, err := l.svcCtx.HomeNewProductService.HomeNewProductUpdate(l.ctx, &smsclient.HomeNewProductUpdateReq{
		Id:              req.Id,
		ProductId:       req.ProductId,
		ProductName:     req.ProductName,
		RecommendStatus: req.RecommendStatus,
		Sort:            req.Sort,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新新鲜好物表信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新新鲜好物表失败")
	}

	return &types.UpdateHomeNewProductResp{
		Code:    "000000",
		Message: "更新新鲜好物表成功",
	}, nil
}
