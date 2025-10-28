package noticeservicelogic

import (
	"context"
	"errors"

	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateNoticeStatusLogic 更新通知公告
/*
Author: 刘飞华
Date: 2025/10/27 15:49:44
*/
type UpdateNoticeStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNoticeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNoticeStatusLogic {
	return &UpdateNoticeStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateNoticeStatus 更新通知公告状态
func (l *UpdateNoticeStatusLogic) UpdateNoticeStatus(in *sysclient.UpdateNoticeStatusReq) (*sysclient.UpdateNoticeStatusResp, error) {
	q := query.SysNotice

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "更新通知公告状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新通知公告状态失败")
	}

	return &sysclient.UpdateNoticeStatusResp{}, nil
}
