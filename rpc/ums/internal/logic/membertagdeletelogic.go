package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberTagDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTagDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTagDeleteLogic {
	return &MemberTagDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTagDeleteLogic) MemberTagDelete(in *ums.MemberTagDeleteReq) (*ums.MemberTagDeleteResp, error) {
	err := l.svcCtx.UmsMemberTagModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &ums.MemberTagDeleteResp{}, nil
}
