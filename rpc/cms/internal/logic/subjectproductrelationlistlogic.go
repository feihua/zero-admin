package logic

import (
	"context"

	"zero-admin/rpc/cms/cmsclient"
	"zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubjectProductRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubjectProductRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubjectProductRelationListLogic {
	return &SubjectProductRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubjectProductRelationListLogic) SubjectProductRelationList(in *cmsclient.SubjectProductRelationListReq) (*cmsclient.SubjectProductRelationListResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.SubjectProductRelationListResp{}, nil
}
