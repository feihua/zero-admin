package membertagservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberTagAddLogic 添加会员标签
/*
Author: LiuFeiHua
Date: 2024/5/7 9:51
*/
type MemberTagAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTagAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTagAddLogic {
	return &MemberTagAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberTagAdd 添加会员标签
func (l *MemberTagAddLogic) MemberTagAdd(in *umsclient.MemberTagAddReq) (*umsclient.MemberTagAddResp, error) {
	err := query.UmsMemberTag.WithContext(l.ctx).Create(&model.UmsMemberTag{
		TagName:           in.Name,
		FinishOrderCount:  in.FinishOrderCount,
		FinishOrderAmount: in.FinishOrderAmount,
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberTagAddResp{}, nil
}
