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

type HomeRecommendSubjectAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendSubjectAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeRecommendSubjectAddLogic {
	return HomeRecommendSubjectAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeRecommendSubjectAddLogic) HomeRecommendSubjectAdd(req types.AddHomeRecommendSubjectReq) (*types.AddHomeRecommendSubjectResp, error) {
	_, err := l.svcCtx.Sms.HomeRecommendSubjectAdd(l.ctx, &smsclient.HomeRecommendSubjectAddReq{
		SubjectId:       req.SubjectId,
		SubjectName:     req.SubjectName,
		RecommendStatus: req.RecommendStatus,
		Sort:            req.Sort,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.Errorf("添加人气推荐专题参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加人气推荐专题失败")
	}

	return &types.AddHomeRecommendSubjectResp{
		Code:    "000000",
		Message: "添加人气推荐专题成功",
	}, nil
}
