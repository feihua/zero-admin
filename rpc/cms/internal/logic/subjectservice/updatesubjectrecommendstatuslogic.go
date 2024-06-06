package subjectservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateSubjectRecommendStatusLogic 批量更新专题状态
/*
Author: LiuFeiHua
Date: 2024/5/13 10:31
*/
type UpdateSubjectRecommendStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSubjectRecommendStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubjectRecommendStatusLogic {
	return &UpdateSubjectRecommendStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateSubjectRecommendStatus 批量更新专题状态
func (l *UpdateSubjectRecommendStatusLogic) UpdateSubjectRecommendStatus(in *cmsclient.UpdateSubjectRecommendStatusReq) (*cmsclient.UpdateSubjectRecommendStatusResp, error) {
	subject := query.CmsSubject

	_, err := subject.WithContext(l.ctx).Where(subject.ID.In(in.Ids...)).Update(subject.RecommendStatus, in.Status)
	if err != nil {
		logc.Errorf(l.ctx, "批量更新专题状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("批量更新专题状态失败")
	}

	return &cmsclient.UpdateSubjectRecommendStatusResp{}, nil
}
