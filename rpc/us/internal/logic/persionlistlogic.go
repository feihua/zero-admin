package logic

import (
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-admin/rpc/us/internal/svc"
	"go-zero-admin/rpc/us/us"
)

type PersionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPersionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PersionListLogic {
	return &PersionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PersionListLogic) PersionList(in *us.PersionListReq) (*us.PersionListResp, error) {
	// todo: add your logic here and delete this line
	all, _ := l.svcCtx.UsUsersModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UsUsersModel.Count()

	//var list []*ums.GrowthChangeHistoryListData
	var list []*us.PersionData
	for _, item := range *all {

		list = append(list, &us.PersionData{
			Id:          item.Id,
			PhoneNumber: item.PhoneNumber.String,
			Password:    item.Password.String,
			UniqueId:    item.UniqueId.String,
			Number:      item.Number.String,
			Email:       item.Email.String,
			Gender:      item.Sex.String,
			Name:        item.Name.String,
			RoleType:    item.RoleId.Int64,
			State:       item.State.Int64,
			CreateTime:    item.CreateTime.Time.String(),
			UpdateTime:    item.UpdateTime.Time.String(),
		})
	}

	fmt.Println(list)

	return &us.PersionListResp{
		Total: count,
		List:  list,
	}, nil
	return &us.PersionListResp{}, nil
}
