package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/stores/sqlc"

	"go-zero-admin/rpc/us/internal/svc"
	"go-zero-admin/rpc/us/us"

	"github.com/tal-tech/go-zero/core/logx"
)

type PersionInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPersionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PersionInfoLogic {
	return &PersionInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PersionInfoLogic) PersionInfo(in *us.PersionInfoReq) (*us.PersionInfoResp, error) {
	// todo: add your logic here and delete this line

	usUser, err := l.svcCtx.UsUsersModel.FindOne(in.Id)
	switch err {
	case nil:
	case sqlc.ErrNotFound:
		return nil, errorUserNotFound
	default:
		return nil, err
	}

	usRole, err := l.svcCtx.UsRolesModel.FindOne(usUser.RoleId.Int64)
	roleName := ""
	if err == nil {
		roleName = usRole.RoleName.String
	}

	roleExtendMap, _ := GetRoleExtendInfoByRoleName(l.svcCtx, usRole.RoleName.String, usUser.Id)

	return &us.PersionInfoResp{
		Info: &us.PersionInfo{
			Id:          usUser.Id,
			PhoneNumber: usUser.PhoneNumber.String,
			UniqueId:    usUser.UniqueId.String,
			Number:      usUser.Number.String,
			Email:       usUser.Email.String,
			Gender:      usUser.Sex.String,
			Name:        usUser.Name.String,
			State:       usUser.State.Int64,
			CreateTime:  usUser.CreateTime.Time.String(),
			RoleId:      usUser.RoleId.Int64,
			RoleName:    roleName,
			Class:       roleExtendMap["class"],
			Academy:     roleExtendMap["academy"],
			School:      roleExtendMap["school"],
			Grade:       roleExtendMap["grade"],
		},
	}, nil

}
