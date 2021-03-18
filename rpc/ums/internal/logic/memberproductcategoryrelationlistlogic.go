package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberProductCategoryRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberProductCategoryRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberProductCategoryRelationListLogic {
	return &MemberProductCategoryRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberProductCategoryRelationListLogic) MemberProductCategoryRelationList(in *ums.MemberProductCategoryRelationListReq) (*ums.MemberProductCategoryRelationListResp, error) {
	all, _ := l.svcCtx.UmsMemberProductCategoryRelationModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsMemberProductCategoryRelationModel.Count()

	var list []*ums.MemberProductCategoryRelationListData
	for _, item := range *all {

		list = append(list, &ums.MemberProductCategoryRelationListData{
			Id:                item.Id,
			MemberId:          item.MemberId,
			ProductCategoryId: item.ProductCategoryId,
		})
	}

	fmt.Println(list)
	return &ums.MemberProductCategoryRelationListResp{
		Total: count,
		List:  list,
	}, nil

}
