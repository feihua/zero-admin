package logic

import (
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-admin/rpc/sys/internal/svc"
	"go-zero-admin/rpc/sys/sys"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListLogic) UserList(in *sys.UserListReq) (*sys.UserListResp, error) {

	all, _ := l.svcCtx.UserModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	var list []*sys.UserListData
	for _, user := range *all {

		list = append(list, &sys.UserListData{
			Id:             user.Id,
			Name:           user.Name,
			NickName:       user.NickName,
			Avatar:         user.Avatar,
			Password:       user.Password,
			Salt:           user.Salt,
			Email:          user.Email,
			Mobile:         user.Mobile,
			Status:         user.Status,
			CreateBy:       user.CreateBy,
			CreateTime:     user.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   user.LastUpdateBy,
			LastUpdateTime: user.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        user.DelFlag,
		})
	}

	fmt.Println(list)
	return &sys.UserListResp{
		Total: 10,
		List:  list,
	}, nil
}
