package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLevelDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLevelDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLevelDeleteLogic {
	return &MemberLevelDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLevelDeleteLogic) MemberLevelDelete(in *ums.MemberLevelDeleteReq) (*ums.MemberLevelDeleteResp, error) {
	err := l.svcCtx.UmsMemberLevelModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &ums.MemberLevelDeleteResp{}, nil
}
