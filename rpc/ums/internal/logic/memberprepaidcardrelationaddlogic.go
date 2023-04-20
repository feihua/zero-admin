package logic

import (
	"context"
	"time"

	"zero-admin/rpc/model/umsmodel"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberPrepaidCardRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberPrepaidCardRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberPrepaidCardRelationAddLogic {
	return &MemberPrepaidCardRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberPrepaidCardRelationAddLogic) MemberPrepaidCardRelationAdd(in *umsclient.MemberPrepaidCardRelationAddReq) (*umsclient.MemberPrepaidCardRelationAddResp, error) {
	currentTime := time.Now()
	_, err := l.svcCtx.UmsMemberPrepaidCardRelationModel.Insert(l.ctx, &umsmodel.UmsMemberPrepaidCardRelation{
		MemberId:      in.MemberId,
		PrepaidCardId: in.PrepaidCardId,
		PrepaidCardSn: in.PrepaidCardSn,
		Balance:       0.00,
		Status:        0,
		CreateTime:    currentTime,
		ExpiredTime:   currentTime.AddDate(3, 0, 0),
		Note:          "",
	})
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberPrepaidCardRelationAddResp{}, nil
}
