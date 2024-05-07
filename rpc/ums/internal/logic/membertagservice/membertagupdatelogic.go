package membertagservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/model"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberTagUpdateLogic 添加会员标签
/*
Author: LiuFeiHua
Date: 2024/5/7 9:53
*/
type MemberTagUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTagUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTagUpdateLogic {
	return &MemberTagUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberTagUpdate 更新添加会员标签
func (l *MemberTagUpdateLogic) MemberTagUpdate(in *umsclient.MemberTagUpdateReq) (*umsclient.MemberTagUpdateResp, error) {
	_, err := query.UmsMemberTag.WithContext(l.ctx).Updates(&model.UmsMemberTag{
		ID:                in.ID,
		Name:              in.Name,
		FinishOrderCount:  in.FinishOrderCount,
		FinishOrderAmount: float64(in.FinishOrderAmount),
	})

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberTagUpdateResp{}, nil
}
