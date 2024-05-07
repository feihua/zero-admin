package membertaskservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberTaskDeleteLogic 会员任务
/*
Author: LiuFeiHua
Date: 2024/5/7 9:20
*/
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

// MemberTaskDelete 删除会员任务
func (l *MemberTaskDeleteLogic) MemberTaskDelete(in *umsclient.MemberTaskDeleteReq) (*umsclient.MemberTaskDeleteResp, error) {
	q := query.UmsMemberTask
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberTaskDeleteResp{}, nil
}
