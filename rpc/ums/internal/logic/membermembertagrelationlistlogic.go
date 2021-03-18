package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberMemberTagRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberMemberTagRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberMemberTagRelationListLogic {
	return &MemberMemberTagRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberMemberTagRelationListLogic) MemberMemberTagRelationList(in *ums.MemberMemberTagRelationListReq) (*ums.MemberMemberTagRelationListResp, error) {
	all, _ := l.svcCtx.UmsMemberMemberTagRelationModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsMemberMemberTagRelationModel.Count()

	var list []*ums.MemberMemberTagRelationListData
	for _, item := range *all {

		list = append(list, &ums.MemberMemberTagRelationListData{
			Id:       item.Id,
			MemberId: item.MemberId,
			TagId:    item.TagId,
		})
	}

	fmt.Println(list)
	return &ums.MemberMemberTagRelationListResp{
		Total: count,
		List:  list,
	}, nil

}
