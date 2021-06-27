package logic

import (
	"context"
	"database/sql"
	"go-zero-admin/rpc/model/usmodel"
	"time"

	"go-zero-admin/rpc/us/internal/svc"
	"go-zero-admin/rpc/us/us"

	"github.com/tal-tech/go-zero/core/logx"
)

type PersionUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPersionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PersionUpdateLogic {
	return &PersionUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PersionUpdateLogic) PersionUpdate(in *us.PersionUpdateReq) (*us.PersionUpdateResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.UsUsersModel.Update(usmodel.UsUsers{
		PhoneNumber: sql.NullString{
			String: in.Data.PhoneNumber,
			Valid:  true,
		},
		UniqueId: sql.NullString{
			String: in.Data.UniqueId,
			Valid:  true,
		},
		Number: sql.NullString{
			String: in.Data.Number,
			Valid:  true,
		},
		Email: sql.NullString{
			String: in.Data.Email,
			Valid:  true,
		},
		Name: sql.NullString{
			String: in.Data.Name,
			Valid:  true,
		},
		Password: sql.NullString{
			String: in.Data.Password,
			Valid:  true,
		},
		Sex: sql.NullString{
			String: in.Data.Gender,
			Valid:  true,
		},
		RoleId: sql.NullInt64{
			Int64: in.Data.RoleType,
			Valid: true,
		},
		State: sql.NullInt64{
			Int64: in.Data.State,
			Valid: true,
		},
		UpdateTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil{
		return &us.PersionUpdateResp{
			Result: false,
		}, err
	}

	return &us.PersionUpdateResp{
		Result: true,
	}, nil
}
