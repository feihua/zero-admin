package subjectservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateSubjectStatusLogic 更新专题
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type UpdateSubjectStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSubjectStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubjectStatusLogic {
	return &UpdateSubjectStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateSubjectStatus 更新专题状态
func (l *UpdateSubjectStatusLogic) UpdateSubjectStatus(in *cmsclient.UpdateSubjectStatusReq) (*cmsclient.UpdateSubjectStatusResp, error) {
	q := query.CmsSubject

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.ShowStatus, in.ShowStatus)

	if err != nil {
		logc.Errorf(l.ctx, "更新专题状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新专题状态失败")
	}

	return &cmsclient.UpdateSubjectStatusResp{}, nil
}
