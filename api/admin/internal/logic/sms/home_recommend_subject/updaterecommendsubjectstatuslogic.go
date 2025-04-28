package home_recommend_subject

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateRecommendSubjectStatusLogic 人气推荐专题
/*
Author: LiuFeiHua
Date: 2024/5/14 9:43
*/
type UpdateRecommendSubjectStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRecommendSubjectStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRecommendSubjectStatusLogic {
	return &UpdateRecommendSubjectStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateRecommendSubjectStatus 批量修改推荐状态
// 1.修改sms_home_recommend_subject的记录推荐状态(sms-rpc)
// 2.修改cms_subject记录的推荐状态(cms-rpc)
func (l *UpdateRecommendSubjectStatusLogic) UpdateRecommendSubjectStatus(req *types.UpdateRecommendSubjectStatusReq) (resp *types.BaseResp, err error) {
	// 1.修改sms_home_recommend_subject的记录推荐状态(sms-rpc)
	_, err = l.svcCtx.HomeRecommendSubjectService.UpdateHomeRecommendSubjectStatus(l.ctx, &smsclient.UpdateHomeRecommendSubjectStatusReq{
		Ids:             req.Ids,
		RecommendStatus: req.RecommendStatus, // 推荐状态：0->不推荐;1->推荐
	})

	if err != nil {
		logc.Errorf(l.ctx, "批量修改推荐状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	// 2.修改cms_subject记录的推荐状态(cms-rpc)
	_, err = l.svcCtx.SubjectService.UpdateSubjectRecommendStatus(l.ctx, &cmsclient.UpdateSubjectRecommendStatusReq{
		Ids:    req.SubjectIds,
		Status: req.RecommendStatus,
	})
	if err != nil {
		logc.Errorf(l.ctx, "根据Ids: %+v,更新人气推荐专题状态异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
