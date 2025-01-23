package memberreportservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberReportStatusLogic 更新用户举报
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type UpdateMemberReportStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberReportStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberReportStatusLogic {
	return &UpdateMemberReportStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberReportStatus 更新用户举报状态
func (l *UpdateMemberReportStatusLogic) UpdateMemberReportStatus(in *cmsclient.UpdateMemberReportStatusReq) (*cmsclient.UpdateMemberReportStatusResp, error) {
	q := query.CmsMemberReport

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.ReportStatus, in.ReportStatus)

	if err != nil {
		logc.Errorf(l.ctx, "更新用户举报状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新用户举报状态失败")
	}

	logc.Infof(l.ctx, "更新用户举报状态成功,参数：%+v", in)
	return &cmsclient.UpdateMemberReportStatusResp{}, nil
}
