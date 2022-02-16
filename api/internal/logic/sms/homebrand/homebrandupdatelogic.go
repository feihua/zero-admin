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

type HomeBrandUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeBrandUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeBrandUpdateLogic {
	return HomeBrandUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeBrandUpdateLogic) HomeBrandUpdate(req types.UpdateHomeBrandReq) (*types.UpdateHomeBrandResp, error) {
	_, err := l.svcCtx.Sms.HomeBrandUpdate(l.ctx, &smsclient.HomeBrandUpdateReq{
		Id:              req.Id,
		BrandId:         req.BrandId,
		BrandName:       req.BrandName,
		RecommendStatus: req.RecommendStatus,
		Sort:            req.Sort,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新首页品牌信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新首页品牌失败")
	}

	return &types.UpdateHomeBrandResp{
		Code:    "000000",
		Message: "更新首页品牌成功",
	}, nil
}
