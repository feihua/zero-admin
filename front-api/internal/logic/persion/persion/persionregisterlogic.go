package logic

import (
	"context"
	"go-zero-admin/rpc/us/usclient"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PersionRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPersionRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) PersionRegisterLogic {
	return PersionRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PersionRegisterLogic) PersionRegister(req types.PersionRegisterReq) (*types.PersionRegisterResp, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.Us.PersionRegister(l.ctx, &usclient.PersionRegisterReq{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
		Number:      req.Number,
		Email:       req.Email,
		Gender:      req.Gender,
		Name:        req.Name,
		RoleId:      req.RoleId,
		Class:       req.Class,
		Academy:     req.Academy,
		School:      req.School,
		Grade:       req.Grade,
	})
	if err != nil {
		return nil, err
	}

	return &types.PersionRegisterResp{
		Code:    0,
		Message: "success",
		Data: types.PersionRegisterRespData{
			Result: resp.Result,
		},
	}, nil
}
