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
	_, err := l.svcCtx.Sms.HomeRecommendSubjectUpdate(l.ctx, &smsclient.HomeRecommendSubjectUpdateReq{
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
