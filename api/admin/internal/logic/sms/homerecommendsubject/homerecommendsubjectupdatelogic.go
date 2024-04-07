package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeRecommendSubjectUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendSubjectUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeRecommendSubjectUpdateLogic {
	return HomeRecommendSubjectUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeRecommendSubjectUpdateLogic) HomeRecommendSubjectUpdate(req types.UpdateHomeRecommendSubjectReq) (*types.UpdateHomeRecommendSubjectResp, error) {
	_, err := l.svcCtx.HomeRecommendSubjectService.HomeRecommendSubjectUpdate(l.ctx, &smsclient.HomeRecommendSubjectUpdateReq{
		Id:              req.Id,
		SubjectId:       req.SubjectId,
		SubjectName:     req.SubjectName,
		RecommendStatus: req.RecommendStatus,
		Sort:            req.Sort,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新人气推荐专题信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新人气推荐专题失败")
	}

	return &types.UpdateHomeRecommendSubjectResp{
		Code:    "000000",
		Message: "更新人气推荐专题成功",
	}, nil
}
