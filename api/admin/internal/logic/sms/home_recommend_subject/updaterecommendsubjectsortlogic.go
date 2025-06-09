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

// UpdateRecommendSubjectSortLogic 人气推荐专题
/*
Author: LiuFeiHua
Date: 2024/5/14 9:43
*/
type UpdateRecommendSubjectSortLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRecommendSubjectSortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRecommendSubjectSortLogic {
	return &UpdateRecommendSubjectSortLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateRecommendSubjectSort 修改推荐排序
func (l *UpdateRecommendSubjectSortLogic) UpdateRecommendSubjectSort(req *types.UpdateRecommendSubjectSortReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.SubjectService.UpdateSubjectSort(l.ctx, &cmsclient.UpdateSubjectSortReq{
		Id:   req.SubjectId,
		Sort: req.Sort,
	})

	if err != nil {
		logc.Errorf(l.ctx, "修改推荐排序失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}
	return res.Success()
}
