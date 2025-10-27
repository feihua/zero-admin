package sys_notice

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteNoticeLogic 删除通知公告表
/*
Author: 刘飞华
Date: 2025/10/27 15:51:14
*/
type DeleteNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteNoticeLogic {
	return &DeleteNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteNotice 删除通知公告表
func (l *DeleteNoticeLogic) DeleteNotice(req *types.DeleteNoticeReq) (resp *types.NoticeResp, err error) {
	_, err = l.svcCtx.NoticeService.DeleteNotice(l.ctx, &sysclient.DeleteNoticeReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除通知公告表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.NoticeResp{
		Code:    "000000",
		Message: "删除通知公告表成功",
	}, nil
}
