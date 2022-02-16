package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberDeleteLogic {
	return &MemberDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberDeleteLogic) MemberDelete(in *ums.MemberDeleteReq) (*ums.MemberDeleteResp, error) {
	err := l.svcCtx.UmsMemberModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}
	return &ums.MemberDeleteResp{}, nil
}
