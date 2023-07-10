package membertaskservicelogic

import (
	"context"
	"zero-admin/rpc/model/umsmodel"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/rpc/ums/internal/svc"

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
