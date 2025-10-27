package sys_notice_service

import (
	"context"
	"errors"

	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteNoticeLogic 删除通知公告
/*
Author: 刘飞华
Date: 2025/10/27 15:49:44
*/
type DeleteNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteNoticeLogic {
	return &DeleteNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteNotice 删除通知公告
func (l *DeleteNoticeLogic) DeleteNotice(in *sysclient.DeleteNoticeReq) (*sysclient.DeleteNoticeResp, error) {
	q := query.SysNotice

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除通知公告失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除通知公告失败")
	}

	logc.Infof(l.ctx, "删除通知公告成功,参数：%+v", in)
	return &sysclient.DeleteNoticeResp{}, nil
}
