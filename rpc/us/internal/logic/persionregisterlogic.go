package logic

import (
	"context"
	"database/sql"
	"go-zero-admin/rpc/model/usmodel"
	"time"

	"go-zero-admin/rpc/us/internal/svc"
	"go-zero-admin/rpc/us/us"

	"github.com/tal-tech/go-zero/core/logx"

	"github.com/gogf/gf/util/guid"
)

type PersionRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPersionRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PersionRegisterLogic {
	return &PersionRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PersionRegisterLogic) PersionRegister(in *us.PersionRegisterReq) (*us.PersionRegisterResp, error) {
	// todo: add your logic here and delete this line

	if _,err := l.svcCtx.UsUsersModel.FindOneByPhoneNumber(in.PhoneNumber);err == nil{
		return nil, errorDuplicateMobile
	}

	uniqueId := guid.S()
	_,err := l.svcCtx.UsUsersModel.Insert(usmodel.UsUsers{
		PhoneNumber: sql.NullString{
			String: in.PhoneNumber,
			Valid:  true,
		},
		UniqueId: sql.NullString{
			String: uniqueId,
			Valid:  true,
		},
		Number: sql.NullString{
			String: in.Number,
			Valid:  true,
		},
		Email: sql.NullString{
			String: in.Email,
			Valid:  true,
		},
		Name: sql.NullString{
			String: in.Name,
			Valid:  true,
		},
		Password: sql.NullString{
			String: in.Password,
			Valid:  true,
		},
		Sex: sql.NullString{
			String: in.Gender,
			Valid:  true,
		},
		RoleId: sql.NullInt64{
			Int64: in.RoleId,
			Valid: true,
		},
		State: sql.NullInt64{
			Int64: 1,
			Valid: true,
		},
		CreateTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdateTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil{
		logx.Error("register err:" + err.Error())
		return nil, errorUserRegisterFail
	}

	return &us.PersionRegisterResp{
		Result: true,
	}, nil
}
