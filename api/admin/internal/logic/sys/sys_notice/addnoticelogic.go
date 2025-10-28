package sys_notice

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddNoticeLogic 添加通知公告表
/*
Author: 刘飞华
Date: 2025/10/27 15:51:14
*/
type AddNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddNoticeLogic {
	return &AddNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddNotice 添加通知公告表
func (l *AddNoticeLogic) AddNotice(req *types.AddNoticeReq) (resp *types.NoticeResp, err error) {
	userName, err := common.GetUserName(l.ctx)
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.NoticeService.AddNotice(l.ctx, &sysclient.AddNoticeReq{
		NoticeTitle:   req.NoticeTitle,   // 公告标题
		NoticeType:    req.NoticeType,    // 公告类型（1:通知,2:公告）
		NoticeContent: req.NoticeContent, // 公告内容
		Status:        req.Status,        // 公告状态（0:关闭,1:正常 ）
		Remark:        req.Remark,        // 备注
		CreateBy:      userName,          // 创建者

	})

	if err != nil {
		logc.Errorf(l.ctx, "添加通知公告表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.NoticeResp{
		Code:    "000000",
		Message: "添加通知公告表成功",
	}, nil
}
