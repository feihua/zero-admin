package member_task

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberTaskLogic 删除会员任务
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
*/
type DeleteMemberTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMemberTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberTaskLogic {
	return &DeleteMemberTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteMemberTask 删除会员任务
func (l *DeleteMemberTaskLogic) DeleteMemberTask(req *types.DeleteMemberTaskReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.MemberTaskService.DeleteMemberTask(l.ctx, &umsclient.DeleteMemberTaskReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除会员任务失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "删除会员任务成功",
	}, nil
}
