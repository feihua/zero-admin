package membertagservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/model/umsmodel"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *MemberTagUpdateLogic) MemberTagUpdate(in *umsclient.MemberTagUpdateReq) (*umsclient.MemberTagUpdateResp, error) {
	err := l.svcCtx.UmsMemberTagModel.Update(l.ctx, &umsmodel.UmsMemberTag{
		Id:                in.Id,
		Name:              in.Name,
		FinishOrderCount:  in.FinishOrderCount,
		FinishOrderAmount: float64(in.FinishOrderAmount),
	})
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberTagUpdateResp{}, nil
}
