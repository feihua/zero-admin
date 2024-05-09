package memberpriceservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberPriceDeleteLogic 会员价格
/*
Author: LiuFeiHua
Date: 2024/5/8 9:52
*/
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

// MemberPriceDelete 删除会员价格
func (l *MemberPriceDeleteLogic) MemberPriceDelete(in *pmsclient.MemberPriceDeleteReq) (*pmsclient.MemberPriceDeleteResp, error) {
	q := query.PmsMemberPrice
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.MemberPriceDeleteResp{}, nil
}
