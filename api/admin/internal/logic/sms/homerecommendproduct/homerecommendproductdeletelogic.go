package homerecommendproduct

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeRecommendProductDeleteLogic 人气推荐商品
/*
Author: LiuFeiHua
Date: 2024/5/14 9:39
*/
type HomeRecommendProductDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeRecommendProductDeleteLogic {
	return HomeRecommendProductDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// HomeRecommendProductDelete 删除人气推荐商品
func (l *HomeRecommendProductDeleteLogic) HomeRecommendProductDelete(req types.DeleteHomeRecommendProductReq) (*types.DeleteHomeRecommendProductResp, error) {
	_, err := l.svcCtx.HomeRecommendProductService.HomeRecommendProductDelete(l.ctx, &smsclient.HomeRecommendProductDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除人气推荐商品异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除人气推荐商品失败")
	}
	return &types.DeleteHomeRecommendProductResp{
		Code:    "000000",
		Message: "",
	}, nil
}
