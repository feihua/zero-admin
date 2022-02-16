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
	_, err := l.svcCtx.Sms.HomeNewProductAdd(l.ctx, &smsclient.HomeNewProductAddReq{
		ProductId:       req.ProductId,
		ProductName:     req.ProductName,
		RecommendStatus: req.RecommendStatus,
		Sort:            req.Sort,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加新鲜好物信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加新鲜好物表失败")
	}

	return &types.AddHomeNewProductResp{
		Code:    "000000",
		Message: "添加新鲜好物表成功",
	}, nil
}
