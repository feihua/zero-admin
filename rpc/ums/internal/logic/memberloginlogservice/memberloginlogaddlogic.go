package memberloginlogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberLoginLogAddLogic 会员登录日志
/*
Author: LiuFeiHua
Date: 2024/5/7 10:09
*/
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

// MemberLoginLogAdd 添加会员登录日志
func (l *MemberLoginLogAddLogic) MemberLoginLogAdd(in *umsclient.MemberLoginLogAddReq) (*umsclient.MemberLoginLogAddResp, error) {
	err := query.UmsMemberLoginLog.WithContext(l.ctx).Create(&model.UmsMemberLoginLog{
		MemberID:   in.MemberId,
		CreateTime: time.Now(),
		MemberIP:   in.MemberIP,
		City:       in.City,
		LoginType:  in.LoginType,
		Province:   in.Province,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberLoginLogAddResp{}, nil
}
