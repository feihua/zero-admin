package memberproductcategoryrelationservicelogic

import (
	"context"
	"zero-admin/rpc/model/umsmodel"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberProductCategoryRelationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberProductCategoryRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberProductCategoryRelationUpdateLogic {
	return &MemberProductCategoryRelationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberProductCategoryRelationUpdateLogic) MemberProductCategoryRelationUpdate(in *umsclient.MemberProductCategoryRelationUpdateReq) (*umsclient.MemberProductCategoryRelationUpdateResp, error) {
	err := l.svcCtx.UmsMemberProductCategoryRelationModel.Update(l.ctx, &umsmodel.UmsMemberProductCategoryRelation{
		Id:                in.Id,
		MemberId:          in.MemberId,
		ProductCategoryId: in.ProductCategoryId,
	})
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberProductCategoryRelationUpdateResp{}, nil
}
