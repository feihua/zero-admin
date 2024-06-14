package membertagservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberTagStatusLogic 更新用户标签表状态
type UpdateMemberTagStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberTagStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberTagStatusLogic {
	return &UpdateMemberTagStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberTagStatus 更新用户标签表状态
func (l *UpdateMemberTagStatusLogic) UpdateMemberTagStatus(in *umsclient.UpdateMemberTagStatusReq) (*umsclient.UpdateMemberTagStatusResp, error) {
	q := query.UmsMemberTag
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)
	if err != nil {
		return nil, err
	}

	return &umsclient.UpdateMemberTagStatusResp{}, nil
}
