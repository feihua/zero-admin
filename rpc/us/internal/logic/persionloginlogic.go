package logic

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"go-zero-admin/rpc/us/internal/svc"
	"go-zero-admin/rpc/us/us"
)

type PersionLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPersionLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PersionLoginLogic {
	return &PersionLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PersionLoginLogic) PersionLogin(in *us.PersionLoginReq) (*us.PersionLoginResp, error) {
	// todo: add your logic here and delete this line
	userInfo, err := l.svcCtx.UsUsersModel.FindOneByPhoneNumber(in.PhoneNumber)


	switch err {
	case nil:
	case sqlc.ErrNotFound:
		return nil, errorUsernameUnRegister
	default:
		return nil, err
	}
	if userInfo.Password.String != in.Password {
		return nil, errorIncorrectPassword
	}

	usRole, err := l.svcCtx.UsRolesModel.FindOne(userInfo.RoleId.Int64)
	roleName := ""
	if err == nil{
		roleName = usRole.RoleName.String
	}

	return &us.PersionLoginResp{
		Id:          userInfo.Id,
		PhoneNumber: userInfo.PhoneNumber.String,
		UniqueId:    userInfo.UniqueId.String,
		Number:      userInfo.Number.String,
		Email:       userInfo.Email.String,
		Gender:      userInfo.Sex.String,
		Name:        userInfo.Name.String,
		State:       userInfo.State.Int64,
		CreateAt:    userInfo.CreateTime.Time.String(),
		RoleId:      userInfo.RoleId.Int64,
		RoleName:    roleName,
	}, nil

}

func (l *PersionLoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
