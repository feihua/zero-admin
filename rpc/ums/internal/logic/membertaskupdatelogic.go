package logic

import (
	"context"
	"zero-admin/rpc/model/umsmodel"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberTaskUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTaskUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTaskUpdateLogic {
	return &MemberTaskUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTaskUpdateLogic) MemberTaskUpdate(in *ums.MemberTaskUpdateReq) (*ums.MemberTaskUpdateResp, error) {
	err := l.svcCtx.UmsMemberTaskModel.Update(umsmodel.UmsMemberTask{
		Id:           in.Id,
		Name:         in.Name,
		Growth:       in.Growth,
		Intergration: in.Intergration,
		Type:         in.Type,
	})
	if err != nil {
		return nil, err
	}

	return &ums.MemberTaskUpdateResp{}, nil
}
