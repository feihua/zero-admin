package logic

import (
	"context"
	"time"

	"zero-admin/rpc/model/umsmodel"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLoginLogUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLoginLogUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogUpdateLogic {
	return &MemberLoginLogUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLoginLogUpdateLogic) MemberLoginLogUpdate(in *umsclient.MemberLoginLogUpdateReq) (*umsclient.MemberLoginLogUpdateResp, error) {
	// todo: add your logic here and delete this line
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	err := l.svcCtx.UmsMemberLoginLogModel.Update(l.ctx, &umsmodel.UmsMemberLoginLog{
		Id:         in.Id,
		MemberId:   in.MemberId,
		CreateTime: CreateTime,
		Ip:         in.Ip,
		City:       in.City,
		LoginType:  in.LoginType,
		Province:   in.Province,
	})
	if err != nil {
		return nil, err
	}
	return &umsclient.MemberLoginLogUpdateResp{}, nil
}
