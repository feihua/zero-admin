package home_recommend_subject

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeRecommendSubjectAddLogic 人气推荐专题
/*
Author: LiuFeiHua
Date: 2024/5/14 9:42
*/
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

// HomeRecommendSubjectAdd 添加人气推荐
func (l *HomeRecommendSubjectAddLogic) HomeRecommendSubjectAdd(req *types.AddHomeRecommendSubjectReq) (*types.BaseResp, error) {

	_, err := l.svcCtx.SubjectService.UpdateSubjectRecommendStatus(l.ctx, &cmsclient.UpdateSubjectRecommendStatusReq{
		Ids:    req.SubjectIds,
		Status: 1,
	})
	if err != nil {
		logc.Errorf(l.ctx, "根据Ids: %+v,更新人气推荐专题状态异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
