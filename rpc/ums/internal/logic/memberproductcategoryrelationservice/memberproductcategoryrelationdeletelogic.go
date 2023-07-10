package memberproductcategoryrelationservicelogic

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberProductCategoryRelationDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberProductCategoryRelationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberProductCategoryRelationDeleteLogic {
	return &MemberProductCategoryRelationDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberProductCategoryRelationDeleteLogic) MemberProductCategoryRelationDelete(in *umsclient.MemberProductCategoryRelationDeleteReq) (*umsclient.MemberProductCategoryRelationDeleteResp, error) {
	err := l.svcCtx.UmsMemberProductCategoryRelationModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberProductCategoryRelationDeleteResp{}, nil
}
