package memberpriceservicelogic

import (
	"context"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberPriceDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberPriceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberPriceDeleteLogic {
	return &MemberPriceDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberPriceDeleteLogic) MemberPriceDelete(in *pmsclient.MemberPriceDeleteReq) (*pmsclient.MemberPriceDeleteResp, error) {
	err := l.svcCtx.PmsMemberPriceModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &pmsclient.MemberPriceDeleteResp{}, nil
}
