package logic

import (
	"context"

	"zero-admin/rpc/model/umsmodel"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberTaskAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTaskAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTaskAddLogic {
	return &MemberTaskAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTaskAddLogic) MemberTaskAdd(in *umsclient.MemberTaskAddReq) (*umsclient.MemberTaskAddResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UmsMemberTaskModel.Insert(l.ctx, &umsmodel.UmsMemberTask{
		Name:         in.Name,
		Growth:       in.Growth,
		Intergration: in.Intergration,
		Type:         in.Type,
	})
	if err != nil {
		return nil, err
	}
	return &umsclient.MemberTaskAddResp{}, nil
}
