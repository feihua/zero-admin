package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

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
	// todo: add your logic here and delete this line
	err := l.svcCtx.UmsMemberProductCategoryRelationModel.Delete(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberProductCategoryRelationDeleteResp{}, nil
}
