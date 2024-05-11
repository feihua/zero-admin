package memberpriceservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberPriceUpdateLogic 会员价格
/*
Author: LiuFeiHua
Date: 2024/5/8 10:48
*/
type MemberPriceUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberPriceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberPriceUpdateLogic {
	return &MemberPriceUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberPriceUpdate 更新会员价格
func (l *MemberPriceUpdateLogic) MemberPriceUpdate(in *pmsclient.MemberPriceUpdateReq) (*pmsclient.MemberPriceUpdateResp, error) {
	q := query.PmsMemberPrice
	_, err := q.WithContext(l.ctx).Updates(&model.PmsMemberPrice{
		ID:              in.Id,
		ProductID:       in.ProductId,
		MemberLevelID:   in.MemberLevelId,
		MemberPrice:     in.MemberPrice,
		MemberLevelName: in.MemberLevelName,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.MemberPriceUpdateResp{}, nil
}
