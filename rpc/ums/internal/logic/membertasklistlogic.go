package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberTaskListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberTaskListLogic {
	return &MemberTaskListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberTaskListLogic) MemberTaskList(in *ums.MemberTaskListReq) (*ums.MemberTaskListResp, error) {
	all, _ := l.svcCtx.UmsMemberTaskModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsMemberTaskModel.Count()

	var list []*ums.MemberTaskListData
	for _, item := range *all {

		list = append(list, &ums.MemberTaskListData{
			Id:           item.Id,
			Name:         item.Name,
			Growth:       item.Growth,
			Intergration: item.Intergration,
			Type:         item.Type,
		})
	}

	fmt.Println(list)
	return &ums.MemberTaskListResp{
		Total: count,
		List:  list,
	}, nil

}
