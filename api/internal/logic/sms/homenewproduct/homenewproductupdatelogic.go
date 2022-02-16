package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sms/smsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

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
	_, err := l.svcCtx.Sms.HomeNewProductUpdate(l.ctx, &smsclient.HomeNewProductUpdateReq{
		Id:              req.Id,
		ProductId:       req.ProductId,
		ProductName:     req.ProductName,
		RecommendStatus: req.RecommendStatus,
		Sort:            req.Sort,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新新鲜好物表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新新鲜好物表失败")
	}

	return &types.UpdateHomeNewProductResp{
		Code:    "000000",
		Message: "更新新鲜好物表成功",
	}, nil
}
