package memberreadhistoryservicelogic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberReadHistoryDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/11/30 16:43
*/
type MemberReadHistoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberReadHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberReadHistoryDeleteLogic {
	return &MemberReadHistoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberReadHistoryDelete 删除浏览记录
func (l *MemberReadHistoryDeleteLogic) MemberReadHistoryDelete(in *umsclient.MemberReadHistoryDeleteReq) (*umsclient.MemberReadHistoryDeleteResp, error) {
	err := l.svcCtx.UmsMemberReadHistoryModel.DeleteByIdAndMemberId(l.ctx, in.Id, in.MemberId)

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberReadHistoryDeleteResp{}, nil
}
