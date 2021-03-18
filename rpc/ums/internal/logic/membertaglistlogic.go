package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTagListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTagListLogic {
	return &MemberTagListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTagListLogic) MemberTagList(in *ums.MemberTagListReq) (*ums.MemberTagListResp, error) {
	all, _ := l.svcCtx.UmsMemberTagModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsMemberTagModel.Count()

	var list []*ums.MemberTagListData
	for _, item := range *all {

		list = append(list, &ums.MemberTagListData{
			Id:                item.Id,
			Name:              item.Name,
			FinishOrderCount:  item.FinishOrderCount,
			FinishOrderAmount: int64(item.FinishOrderAmount),
		})
	}

	fmt.Println(list)
	return &ums.MemberTagListResp{
		Total: count,
		List:  list,
	}, nil

}
