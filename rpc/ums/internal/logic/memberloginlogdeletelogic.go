package logic

import (
	"context"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLoginLogDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLoginLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogDeleteLogic {
	return &MemberLoginLogDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLoginLogDeleteLogic) MemberLoginLogDelete(in *ums.MemberLoginLogDeleteReq) (*ums.MemberLoginLogDeleteResp, error) {
	err := l.svcCtx.UmsMemberLoginLogModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &ums.MemberLoginLogDeleteResp{}, nil
}
