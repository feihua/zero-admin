package memberloginlogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/model/umsmodel"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLoginLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLoginLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogAddLogic {
	return &MemberLoginLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLoginLogAddLogic) MemberLoginLogAdd(in *umsclient.MemberLoginLogAddReq) (*umsclient.MemberLoginLogAddResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	_, err := l.svcCtx.UmsMemberLoginLogModel.Insert(l.ctx, &umsmodel.UmsMemberLoginLog{
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

	return &umsclient.MemberLoginLogAddResp{}, nil
}
