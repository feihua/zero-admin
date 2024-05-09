package memberpriceservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberPriceAddLogic 会员价格
/*
Author: LiuFeiHua
Date: 2024/5/8 10:47
*/
type MemberPriceAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberPriceAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberPriceAddLogic {
	return &MemberPriceAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberPriceAdd 添加会员价格
func (l *MemberPriceAddLogic) MemberPriceAdd(in *pmsclient.MemberPriceAddReq) (*pmsclient.MemberPriceAddResp, error) {
	err := query.PmsMemberPrice.WithContext(l.ctx).Create(&model.PmsMemberPrice{
		ProductID:       in.ProductId,
		MemberLevelID:   in.MemberLevelId,
		MemberPrice:     float64(in.MemberPrice),
		MemberLevelName: in.MemberLevelName,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.MemberPriceAddResp{}, nil
}
