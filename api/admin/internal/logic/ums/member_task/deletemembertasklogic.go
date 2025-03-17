package member_task

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberTaskLogic 删除会员任务
/*
Author: LiuFeiHua
Date: 2024/5/23 9:12
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
func (l *DeleteMemberTaskLogic) DeleteMemberTask(req *types.DeleteMemberTaskReq) (resp *types.DeleteMemberTaskResp, err error) {
	_, err = l.svcCtx.MemberTaskService.DeleteMemberTask(l.ctx, &umsclient.DeleteMemberTaskReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Ids: %+v,删除会员任务异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除会员任务失败")
	}

	return &types.DeleteMemberTaskResp{
		Code:    "000000",
		Message: "删除会员任务成功",
	}, nil
}
