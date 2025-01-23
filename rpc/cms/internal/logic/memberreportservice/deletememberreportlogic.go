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

// DeleteMemberReportLogic 删除用户举报
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type DeleteMemberReportLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberReportLogic {
	return &DeleteMemberReportLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteMemberReport 删除用户举报
func (l *DeleteMemberReportLogic) DeleteMemberReport(in *cmsclient.DeleteMemberReportReq) (*cmsclient.DeleteMemberReportResp, error) {
	q := query.CmsMemberReport

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除用户举报失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除用户举报失败")
	}

	logc.Infof(l.ctx, "删除用户举报成功,参数：%+v", in)
	return &cmsclient.DeleteMemberReportResp{}, nil
}
