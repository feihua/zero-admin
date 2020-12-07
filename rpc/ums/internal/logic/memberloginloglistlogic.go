package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLoginLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogListLogic {
	return &MemberLoginLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLoginLogListLogic) MemberLoginLogList(in *ums.MemberLoginLogListReq) (*ums.MemberLoginLogListResp, error) {
	all, _ := l.svcCtx.UmsMemberLoginLogModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	var list []*ums.MemberLoginLogListData
	for _, item := range *all {

		list = append(list, &ums.MemberLoginLogListData{
			Id:         item.Id,
			MemberId:   item.MemberId,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
			Ip:         item.Ip,
			City:       item.City,
			LoginType:  item.LoginType,
			Province:   item.Province,
		})
	}

	fmt.Println(list)
	return &ums.MemberLoginLogListResp{
		Total: 10,
		List:  list,
	}, nil

}
