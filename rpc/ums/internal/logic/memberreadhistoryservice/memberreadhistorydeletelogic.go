package memberreadhistoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberReadHistoryDeleteLogic 会员浏览记录
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

// MemberReadHistoryDelete 删除会员浏览记录
func (l *MemberReadHistoryDeleteLogic) MemberReadHistoryDelete(in *umsclient.MemberReadHistoryDeleteReq) (*umsclient.MemberReadHistoryDeleteResp, error) {
	q := query.UmsMemberReadHistory
	historyDo := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.MemberId))
	if len(in.Ids) > 0 {
		historyDo.Where(q.ID.In(in.Ids...))
	}
	_, err := historyDo.Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberReadHistoryDeleteResp{}, nil
}
