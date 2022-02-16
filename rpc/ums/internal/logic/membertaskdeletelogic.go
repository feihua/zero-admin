package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberTaskDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTaskDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTaskDeleteLogic {
	return &MemberTaskDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTaskDeleteLogic) MemberTaskDelete(in *ums.MemberTaskDeleteReq) (*ums.MemberTaskDeleteResp, error) {
	err := l.svcCtx.UmsMemberTaskModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &ums.MemberTaskDeleteResp{}, nil
}
