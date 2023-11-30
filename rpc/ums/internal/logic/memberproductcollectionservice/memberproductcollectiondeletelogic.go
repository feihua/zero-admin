package memberproductcollectionservicelogic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberProductCollectionDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberProductCollectionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberProductCollectionDeleteLogic {
	return &MemberProductCollectionDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberProductCollectionDeleteLogic) MemberProductCollectionDelete(in *umsclient.MemberProductCollectionDeleteReq) (*umsclient.MemberProductCollectionDeleteResp, error) {
	err := l.svcCtx.UmsMemberProductCollectionModel.DeleteByIdAndMemberId(l.ctx, in.Id, in.MemberId)

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberProductCollectionDeleteResp{}, nil
}
