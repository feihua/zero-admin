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

// UpdateNoticeStatusLogic 更新通知公告表状态状态
/*
Author: 刘飞华
Date: 2025/10/27 15:51:14
*/
type UpdateNoticeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateNoticeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNoticeStatusLogic {
	return &UpdateNoticeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateNoticeStatus 更新通知公告表状态
func (l *UpdateNoticeStatusLogic) UpdateNoticeStatus(req *types.UpdateNoticeStatusReq) (resp *types.NoticeResp, err error) {
	_, err = l.svcCtx.NoticeService.UpdateNoticeStatus(l.ctx, &sysclient.UpdateNoticeStatusReq{
		Ids:    req.Ids,    // 公告ID
		Status: req.Status, // 公告状态（0:关闭,1:正常 ）
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新通知公告表状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.NoticeResp{
		Code:    "000000",
		Message: "更新通知公告表状态成功",
	}, nil
}
