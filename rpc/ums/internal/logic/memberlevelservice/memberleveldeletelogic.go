package memberlevelservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberLevelDeleteLogic 会员等级
/*
Author: LiuFeiHua
Date: 2024/5/7 9:35
*/
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

// MemberLevelDelete 删除会员等级
func (l *MemberLevelDeleteLogic) MemberLevelDelete(in *umsclient.MemberLevelDeleteReq) (*umsclient.MemberLevelDeleteResp, error) {
	q := query.UmsMemberLevel
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberLevelDeleteResp{}, nil
}
