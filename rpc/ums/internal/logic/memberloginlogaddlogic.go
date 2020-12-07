package logic

import (
	"context"
	"go-zero-admin/rpc/model/umsmodel"
	"time"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *MemberLoginLogAddLogic) MemberLoginLogAdd(in *ums.MemberLoginLogAddReq) (*ums.MemberLoginLogAddResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	_, err := l.svcCtx.UmsMemberLoginLogModel.Insert(umsmodel.UmsMemberLoginLog{
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

	return &ums.MemberLoginLogAddResp{}, nil
}
