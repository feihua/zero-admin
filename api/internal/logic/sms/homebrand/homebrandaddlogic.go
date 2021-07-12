package logic

import (
	"context"
	"encoding/json"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/sms/smsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeBrandAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeBrandAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeBrandAddLogic {
	return HomeBrandAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeBrandAddLogic) HomeBrandAdd(req types.AddHomeBrandReq) (*types.AddHomeBrandResp, error) {
	_, err := l.svcCtx.Sms.HomeBrandAdd(l.ctx, &smsclient.HomeBrandAddReq{
		BrandId:         req.BrandId,
		BrandName:       req.BrandName,
		RecommendStatus: req.RecommendStatus,
		Sort:            req.Sort,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.Errorf("添加首页品牌参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加首页品牌失败")
	}

	return &types.AddHomeBrandResp{
		Code:    "000000",
		Message: "添加首页品牌成功",
	}, nil
}
