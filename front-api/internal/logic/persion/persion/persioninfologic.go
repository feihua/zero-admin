package logic

import (
	"context"
	"go-zero-admin/rpc/us/usclient"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type PersionInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPersionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) PersionInfoLogic {
	return PersionInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PersionInfoLogic) PersionInfo(req types.PersionInfoReq) (*types.PersionInfoResp, error) {
	// todo: add your logic here and delete this line

	resp, err := l.svcCtx.Us.PersionInfo(l.ctx, &usclient.PersionInfoReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	return &types.PersionInfoResp{
		Code:    0,
		Message: "success",
		Data: types.PersionInfo{
			Id:          resp.Info.Id,
			Name:        resp.Info.Name,
			PhoneNumber: resp.Info.PhoneNumber,
			Email:       resp.Info.Email,
			RoleName:    resp.Info.RoleName,
			RoleId:      resp.Info.RoleId,
			Gender:      resp.Info.Gender,
		},
	}, nil

}
