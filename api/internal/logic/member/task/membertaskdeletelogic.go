package logic

import (
	"context"
	"go-zero-admin/rpc/ums/umsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTaskDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTaskDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTaskDeleteLogic {
	return MemberTaskDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTaskDeleteLogic) MemberTaskDelete(req types.DeleteMemberTaskReq) (*types.DeleteMemberTaskResp, error) {
	_, _ = l.svcCtx.Ums.MemberTaskDelete(l.ctx, &umsclient.MemberTaskDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteMemberTaskResp{
		Code:    "000000",
		Message: "",
	}, nil
}
